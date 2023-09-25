package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	clickhouseRepo "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/clickhouse"
	"gorm.io/driver/clickhouse"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/consumers"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	myerrors "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	gm "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/gomail"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/workers"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
	"gopkg.in/gomail.v2"
	"moul.io/zapgorm2"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/kafka"

	"github.com/Shopify/sarama"

	myHTTP "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/http"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/controllers"
	postgresRepo "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/postgres"
	openapi "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	ClickhouseDB = "clickhouse"
	PostgresDB   = "postgres"
)

type App struct {
	Server  myHTTP.Server
	handler http.Handler

	postgresAdminDB  *gorm.DB
	postgresTargetDB *gorm.DB
	postgresClientDB *gorm.DB
	analyticDB       *gorm.DB

	worker   *workers.AdWorker
	consumer *consumers.AdConsumer
	client   sarama.ConsumerGroup

	cfg    Config
	logger *zap.SugaredLogger
}

func New() *App {
	return &App{}
}

func initEventProducer(brokerList []string, logger *zap.SugaredLogger) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 1 * time.Second

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		logger.Errorw("cannot start sarama async producer", "error", err)
		return nil, fmt.Errorf("failed to start sarama async producer: %w", err)
	}

	go func() {
		for err := range producer.Errors() {
			logger.Errorw("cannot write access log entry", "error", err)
		}
	}()

	return producer, nil
}

func initAdProducer(brokerList []string, logger *zap.SugaredLogger) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		logger.Errorw("cannot start sarama sync producer", "error", err)
		return nil, fmt.Errorf("start sarama sync producer: %w", err)
	}

	return producer, nil
}

func (a *App) readConfig(cfgFile string) error {
	viper.SetConfigName(filepath.Base(cfgFile))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(cfgFile))

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read in config: %w", err)
	}

	err = viper.Unmarshal(&a.cfg)
	if err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	return nil
}

func (a *App) initLogger() error {
	lvl, err := zap.ParseAtomicLevel(a.cfg.Logger.Level)
	if err != nil {
		return fmt.Errorf("parse level: %w", err)
	}

	logConfig := zap.Config{
		Level:    lvl,
		Encoding: a.cfg.Logger.Encoding,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "caller",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeTime:   zapcore.RFC3339TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{a.cfg.Logger.File},
	}

	logger, err := logConfig.Build()
	if err != nil {
		return fmt.Errorf("build logger: %w", err)
	}

	a.logger = logger.Sugar()

	return nil
}

func (a *App) initConsumer(logger *zap.SugaredLogger) (sarama.ConsumerGroup, error) {
	version, err := sarama.ParseKafkaVersion(a.cfg.Kafka.Version)
	if err != nil {
		logger.Errorw("cannot parse Kafka version", "error", err)
		return nil, fmt.Errorf("parsing Kafka version: %w", err)
	}

	config := sarama.NewConfig()
	config.Version = version

	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategySticky}

	client, err := sarama.NewConsumerGroup(a.cfg.Kafka.Brokers, "blurby-ad-consumers", config)
	if err != nil {
		logger.Errorw("cannot create consumer group client", "error", err)
		return nil, fmt.Errorf("creating consumer group client: %w", err)
	}

	return client, nil
}

func (a *App) initAdmin(u *controllers.Profile, logger *zap.SugaredLogger) error {
	_, err := u.GetByLogin(context.Background(), "admin")
	if !errors.Is(err, myerrors.ErrUserNotFound) {
		if err != nil {
			logger.Errorw("cannot get admin", "error", err)
			return fmt.Errorf("get: %w", err)
		}
		return nil
	}

	admin := &models.User{
		Login:   a.cfg.AdminLogin,
		IsAdmin: true,
	}

	_, err = u.Register(context.Background(), admin, a.cfg.AdminPassword)
	if err != nil {
		logger.Errorw("cannot register admin", "error", err)
		return fmt.Errorf("create: %w", err)
	}

	return nil
}

func (a *App) Init(cfg string) error {
	err := a.readConfig(cfg)
	if err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	err = a.initLogger()
	if err != nil {
		return fmt.Errorf("init logger: %w", err)
	}

	a.postgresAdminDB, err = gorm.Open(postgres.Open(a.cfg.PGA.toDSN()), &gorm.Config{
		Logger: zapgorm2.New(a.logger.Desugar()),
	})
	if err != nil {
		a.logger.Fatalw("cannot open gorm pg admin connection", "error", err)
		return fmt.Errorf("gorm open: %w", err)
	}

	a.postgresTargetDB, err = gorm.Open(postgres.Open(a.cfg.PGT.toDSN()), &gorm.Config{
		Logger: zapgorm2.New(a.logger.Desugar()),
	})
	if err != nil {
		a.logger.Fatalw("cannot open gorm pg targetologist connection", "error", err)
		return fmt.Errorf("gorm open: %w", err)
	}

	a.postgresClientDB, err = gorm.Open(postgres.Open(a.cfg.PGC.toDSN()), &gorm.Config{
		Logger: zapgorm2.New(a.logger.Desugar()),
	})
	if err != nil {
		a.logger.Fatalw("cannot open gorm pg client connection", "error", err)
		return fmt.Errorf("gorm open: %w", err)
	}

	a.analyticDB, err = gorm.Open(clickhouse.Open(a.cfg.CH.toDSN()), &gorm.Config{
		Logger: zapgorm2.New(a.logger.Desugar()),
	})

	if err != nil {
		a.logger.Fatalw("cannot open gorm ch connection", "error", err)
		return fmt.Errorf("gorm open: %w", err)
	}

	asdr := clickhouseRepo.NewASDR(a.analyticDB)

	ecProducer, err := initEventProducer(a.cfg.Kafka.Brokers, a.logger)
	if err != nil {
		a.logger.Fatalw("cannot init event producer", "error", err)
		return fmt.Errorf("init event producer: %w", err)
	}

	asProducer, err := initAdProducer(a.cfg.Kafka.Brokers, a.logger)
	if err != nil {
		a.logger.Fatalw("cannot init ad producer", "error", err)
		return fmt.Errorf("init ad producer: %w", err)
	}

	var ecr interfaces.EventCreationRepository
	var efr interfaces.EventFilteringRepository

	if a.cfg.DB == ClickhouseDB {
		efr = clickhouseRepo.NewCHEFR(a.analyticDB)
		ecr = kafka.NewECR(ecProducer)
	} else if a.cfg.DB == PostgresDB {
		repo := postgresRepo.NewPGEFR(a.postgresTargetDB, a.postgresClientDB)
		efr = repo
		ecr = repo
	}

	cr := postgresRepo.NewCR(a.postgresAdminDB, a.postgresTargetDB, a.postgresClientDB)
	etr := postgresRepo.NewETR(a.postgresTargetDB)
	ur := postgresRepo.NewUR(a.postgresAdminDB, a.postgresTargetDB)
	ar := postgresRepo.NewAR(a.postgresTargetDB)
	sr := postgresRepo.NewSR(a.postgresTargetDB)

	asr := kafka.NewASR(asProducer)
	astr := kafka.NewASTR(ecProducer)

	d := gomail.NewDialer(a.cfg.SMTP.Host, a.cfg.SMTP.Port, a.cfg.SMTP.Login, a.cfg.SMTP.Password)
	sdr := gm.NewS(d, a.cfg.Email)

	el := controllers.NewEL(cr, ecr, etr)
	pl := controllers.NewPL(ur, a.cfg.SecretKey, a.cfg.TokenExp)
	fl := controllers.NewFL(cr, efr)
	al := controllers.NewAL(ar, sr)
	ap := controllers.NewAP(al, asr, sr)
	sl := controllers.NewSL(fl, cr, sdr, astr)
	cl := controllers.NewCL(cr)
	stl := controllers.NewStatL(cr, ar, asdr, a.cfg.Lim)
	etl := controllers.NewETL(etr)

	a.worker = workers.New(ap, a.cfg.SpanAd, a.logger)
	a.consumer = consumers.NewAC(sl, a.logger)

	a.client, err = a.initConsumer(a.logger)
	if err != nil {
		a.logger.Fatalw("cannot init consumer", "error", err)
		return fmt.Errorf("gorm open: %w", err)
	}

	err = a.initAdmin(pl, a.logger)
	if err != nil {
		a.logger.Fatalw("cannot init admin", "error", err)
		return fmt.Errorf("init admin: %w", err)
	}

	service := myHTTP.New(pl, el, fl, al, cl, stl, etl)

	a.handler = myHTTP.Middleware(pl, a.logger, openapi.NewRouter(openapi.NewDefaultApiController(service)))

	return nil
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Infow("started running application")
	ctx = mycontext.LoggerToContext(ctx, a.logger)

	eg := &errgroup.Group{}

	eg.Go(func() error {
		err := a.worker.Run(ctx)
		if err != nil {
			a.logger.Fatalw("run worker", "error", err)
			return fmt.Errorf("worker run: %w", err)
		}

		return nil
	})

	eg.Go(func() error {
		for {
			if err := a.client.Consume(ctx, []string{"blurby-ads"}, a.consumer); err != nil {
				a.logger.Errorw("run client", "error", err)
				return fmt.Errorf("consume: %w", err)
			}

			a.logger.Debugw("toggled reconsuming")

			if ctx.Err() != nil {
				return nil
			}
		}
	})

	eg.Go(func() error {
		port := fmt.Sprintf(":%d", a.cfg.HTTPPort)
		err := http.ListenAndServe(port, cors.AllowAll().Handler(a.handler))
		if err != nil {
			a.logger.Errorw("listen and serve", "error", err)
			return fmt.Errorf("listen and serve: %w", err)
		}

		return nil
	})

	err := eg.Wait()
	if err != nil {
		a.logger.Fatalw("run app", "error", err)
		return fmt.Errorf("run app: %w", err)
	}

	return nil
}

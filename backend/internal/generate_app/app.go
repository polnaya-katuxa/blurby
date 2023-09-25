package app

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"gorm.io/driver/clickhouse"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"moul.io/zapgorm2"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/kafka"

	"github.com/Shopify/sarama"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/controllers"
	postgresRepo "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	ClickhouseDB = "clickhouse"
	PostgresDB   = "postgres"
)

type GenApp struct {
	postgresTargetDB *gorm.DB
	postgresClientDB *gorm.DB
	analyticDB       *gorm.DB

	g        interfaces.GenerateLogic
	producer sarama.AsyncProducer

	cfg    Config
	logger *zap.SugaredLogger
}

func NewGen() *GenApp {
	return &GenApp{}
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

func (a *GenApp) readConfig(cfgFile string) error {
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

func (a *GenApp) initLogger() error {
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

func (a *GenApp) Init(cfg string) error {
	err := a.readConfig(cfg)
	if err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	err = a.initLogger()
	if err != nil {
		return fmt.Errorf("init logger: %w", err)
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

	var ecr interfaces.EventCreationRepository

	if a.cfg.DB == ClickhouseDB {
		a.producer, err = initEventProducer(a.cfg.Kafka.Brokers, a.logger)
		if err != nil {
			a.logger.Fatalw("cannot init event producer", "error", err)
			return fmt.Errorf("init event producer: %w", err)
		}

		ecr = kafka.NewECR(a.producer)
	} else if a.cfg.DB == PostgresDB {
		ecr = postgresRepo.NewPGEFR(a.postgresTargetDB, a.postgresClientDB)
	}

	cr := postgresRepo.NewCR(a.postgresTargetDB, a.postgresTargetDB, a.postgresClientDB)
	etr := postgresRepo.NewETR(a.postgresTargetDB)

	a.g = controllers.NewGL(cr, etr, ecr)

	return nil
}

func (a *GenApp) Run(ctx context.Context) error {
	a.logger.Infow("started running application")
	ctx = mycontext.LoggerToContext(ctx, a.logger)

	err := a.g.Generate(ctx)
	if err != nil {
		a.logger.Errorw("generate", "error", err)
		return fmt.Errorf("generate: %w", err)
	}

	if a.producer != nil {
		if err := a.producer.Close(); err != nil {
			a.logger.Errorw("close producer", "error", err)
			return fmt.Errorf("close producer: %w", err)
		}
	}

	return nil
}

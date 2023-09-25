package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"

	"moul.io/zapgorm2"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	clickhouseRepo "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/clickhouse"
	postgresRepo "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/postgres"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	NumTests   = 1
	NumFilters = 3
)

func toPGDSN(host string, port string, user string, password string, dbName string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
}

func toCHDSN(host string, port string, user string, password string, dbName string, timeout string) string {
	return fmt.Sprintf("tcp://%s:%s/%s?username=%s&password=%s&dial_timeout=%sms",
		host, port, dbName, user, password, timeout)
}

func makeFilters(etr interfaces.EventTypeRepository) ([]*models.EventFilter, error) {
	f := make([]*models.EventFilter, NumFilters)

	ets, err := etr.GetAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("file open: %w", err)
	}

	for i := range f {
		f[i] = &models.EventFilter{
			Alias: ets[i].Alias,
			Rate:  5,
		}
	}

	return f, nil
}

func RunBenchmark(ctx context.Context, f []*models.EventFilter, efr interfaces.EventFilteringRepository) int {
	start := time.Now()
	for i := 0; i < NumTests; i++ {
		_, err := efr.Filter(ctx, f)
		if err != nil {
			log.Println(err)
		}
	}
	t := int(time.Since(start).Milliseconds()) / NumTests

	return t
}

func main() {
	pgHost := os.Getenv("PG_HOST")
	chHost := os.Getenv("CH_HOST")
	pgPort := os.Getenv("PG_PORT")
	chPort := os.Getenv("CH_PORT")

	pgUserT := os.Getenv("PG_USER_T")
	pgUserC := os.Getenv("PG_USER_C")
	chUser := os.Getenv("CH_USER")

	pgPasswordT := os.Getenv("PG_PASSWORD_T")
	pgPasswordC := os.Getenv("PG_PASSWORD_C")
	chPassword := os.Getenv("CH_PASSWORD")

	chTimeout := os.Getenv("CH_TIMEOUT")

	dbName := os.Getenv("DB_NAME")

	pgTarget, err := gorm.Open(postgres.Open(toPGDSN(pgHost, pgPort, pgUserT, pgPasswordT, dbName)), &gorm.Config{
		Logger: zapgorm2.New(zap.NewNop()),
	})
	if err != nil {
		log.Fatalf("gorm open: %s", err)
	}

	pgClient, err := gorm.Open(postgres.Open(toPGDSN(pgHost, pgPort, pgUserC, pgPasswordC, dbName)), &gorm.Config{
		Logger: zapgorm2.New(zap.NewNop()),
	})
	if err != nil {
		log.Fatalf("gorm open: %s", err)
	}

	ch, err := gorm.Open(clickhouse.Open(toCHDSN(chHost, chPort, chUser, chPassword, dbName, chTimeout)), &gorm.Config{
		Logger: zapgorm2.New(zap.NewNop()),
	})
	if err != nil {
		log.Fatalf("gorm open: %s", err)
	}

	efrCH := clickhouseRepo.NewCHEFR(ch)
	efrPG := postgresRepo.NewPGEFR(pgTarget, pgClient)

	etr := postgresRepo.NewETR(pgTarget)

	timeFilePG, err := os.Create("pg.txt")
	if err != nil {
		log.Fatalf("file open: %s", err)
	}
	defer func() { _ = timeFilePG.Close() }()

	timeFileCH, err := os.Create("ch.txt")
	if err != nil {
		log.Fatalf("file open: %s", err)
	}
	defer func() { _ = timeFileCH.Close() }()

	filters, err := makeFilters(etr)
	if err != nil {
		log.Fatalf("make filters: %s", err)
	}

	for i := range filters {
		log.Printf("Starting benchmark for %d filters", i)

		res := RunBenchmark(context.Background(), filters[:i], efrPG)
		fmt.Fprintf(timeFilePG, "%d %d\n", i, res)

		res = RunBenchmark(context.Background(), filters[:i], efrCH)
		fmt.Fprintf(timeFileCH, "%d %d\n", i, res)

		log.Printf("Ended benchmark for %d filters", i)
	}
}

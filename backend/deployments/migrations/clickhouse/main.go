package main

import (
	"flag"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/pressly/goose/v3"
)

var (
	dir = flag.String("dir", ".", "directory with migration files")
	b   = flag.String("b", "localhost:9092", "kafka broker address")
	dsn = flag.String("dsn", "tcp://localhost:9000/$(SERVICE_NAME)?username=$(SERVICE_NAME)&password=password&dial_timeout=500ms", "db clickhouse connection string")
	cmd = flag.String("cmd", "up", "goose command")
)

func main() {
	flag.Parse()

	Brokers = *b

	goose.SetTableName("goose_db_coursework_version")

	db, err := goose.OpenDBWithDriver("clickhouse", *dsn)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.Run(*cmd, db, *dir); err != nil {
		log.Fatalf("goose %v: %v", *cmd, err)
	}
}

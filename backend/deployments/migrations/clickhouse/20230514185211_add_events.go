package main

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

var Brokers = "localhost:9092"

func init() {
	goose.AddMigration(upAddEvents, downAddEvents)
}

func upAddEvents(tx *sql.Tx) error {
	_, err := tx.Exec(fmt.Sprintf(`
		CREATE TABLE coursework.events_queue (
		id UUID,
		client_id UUID,
		alias String,
		time DateTime('UTC')
		) ENGINE = Kafka('%s', 'blurby-events', 'clickhouse-blurby-events', 'JSONEachRow');
	`, Brokers))
	if err != nil {
		return fmt.Errorf("create event table: %w", err)
	}

	_, err = tx.Exec(`
		CREATE TABLE coursework.events
		(
			id UUID,
			client_id UUID,
			alias String,
			time DateTime('UTC')
		) ENGINE = MergeTree ORDER BY (time);
	`)
	if err != nil {
		return fmt.Errorf("create event tree: %w", err)
	}

	_, err = tx.Exec(`
		CREATE MATERIALIZED VIEW coursework.events_mv TO coursework.events AS
		SELECT *
		FROM coursework.events_queue;
	`)
	if err != nil {
		return fmt.Errorf("create event view: %w", err)
	}

	return nil
}

func downAddEvents(tx *sql.Tx) error {
	_, err := tx.Exec("DROP VIEW IF EXISTS coursework.events_mv;")
	if err != nil {
		return fmt.Errorf("drop event view: %w", err)
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS coursework.events;")
	if err != nil {
		return fmt.Errorf("drop event tree: %w", err)
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS coursework.events_queue;")
	if err != nil {
		return fmt.Errorf("drop event table: %w", err)
	}

	return nil
}

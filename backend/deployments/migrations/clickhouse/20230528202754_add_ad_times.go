package main

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddAdTimes, downAddAdTimes)
}

func upAddAdTimes(tx *sql.Tx) error {
	_, err := tx.Exec(fmt.Sprintf(`
		CREATE TABLE coursework.ad_send_times_queue (
		id UUID,
		time DateTime('UTC')
		) ENGINE = Kafka('%s', 'blurby-ad-send-times', 'clickhouse-blurby-ad-send-times', 'JSONEachRow');
	`, Brokers))
	if err != nil {
		return fmt.Errorf("create ad table: %w", err)
	}

	_, err = tx.Exec(`
		CREATE TABLE coursework.ad_send_times
		(
			id UUID,
			time DateTime('UTC'),
			date alias toDate(time)
		) ENGINE = MergeTree ORDER BY (time);
	`)
	if err != nil {
		return fmt.Errorf("create ad tree: %w", err)
	}

	_, err = tx.Exec(`
		CREATE MATERIALIZED VIEW coursework.ad_send_times_mv TO coursework.ad_send_times AS
		SELECT *
		FROM coursework.ad_send_times_queue;
	`)
	if err != nil {
		return fmt.Errorf("create ad view: %w", err)
	}

	return nil
}

func downAddAdTimes(tx *sql.Tx) error {
	_, err := tx.Exec("DROP VIEW IF EXISTS coursework.ad_send_times_mv;")
	if err != nil {
		return fmt.Errorf("drop ad view: %w", err)
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS coursework.ad_send_times;")
	if err != nil {
		return fmt.Errorf("drop ad tree: %w", err)
	}

	_, err = tx.Exec("DROP TABLE IF EXISTS coursework.ad_send_times_queue;")
	if err != nil {
		return fmt.Errorf("drop ad table: %w", err)
	}

	return nil
}

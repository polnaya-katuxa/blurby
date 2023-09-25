package app

import (
	"fmt"
)

type DB struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
}

type Clickhouse struct {
	DB      `mapstructure:",squash"`
	Timeout int
}

type Postgres struct {
	DB `mapstructure:",squash"`
}

type Kafka struct {
	Brokers []string
	Version string
}

type Logger struct {
	Encoding string
	Level    string
	File     string
}

type Config struct {
	PGT   Postgres
	PGC   Postgres
	CH    Clickhouse
	Kafka Kafka

	DB string

	Logger Logger
}

func (d *Postgres) toDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", d.Host,
		d.Port, d.User, d.Password, d.DBName)
}

func (d *Clickhouse) toDSN() string {
	return fmt.Sprintf("tcp://%s:%d/%s?username=%s&password=%s&dial_timeout=%dms",
		d.Host, d.Port, d.DBName, d.User, d.Password, d.Timeout)
}

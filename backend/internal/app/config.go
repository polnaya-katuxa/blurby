package app

import (
	"fmt"
	"time"
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

type SMTP struct {
	Host     string
	Port     int
	Login    string
	Password string
}

type Logger struct {
	Encoding string
	Level    string
	File     string
}

type Config struct {
	PGA   Postgres
	PGT   Postgres
	PGC   Postgres
	CH    Clickhouse
	Kafka Kafka
	SMTP  SMTP

	HTTPPort      int
	AdminLogin    string
	AdminPassword string
	AdminEmail    string

	DB string

	Email     string
	TokenExp  time.Duration
	SecretKey string
	SpanAd    time.Duration
	Lim       int

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

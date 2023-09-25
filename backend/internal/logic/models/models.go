package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	Gender     string
	FilterType string
	FieldType  string
	CmpType    string
)

const (
	Female Gender = "female"
	Male   Gender = "male"

	ByEvent FilterType = "by event"
	ByField FilterType = "by field"

	ByAge    FieldType = "age"
	ByBD     FieldType = "birth date"
	ByName   FieldType = "name"
	ByGender FieldType = "gender"

	Greater CmpType = ">"
	Less    CmpType = "<"
	Between CmpType = "between"
	Equal   CmpType = "="
	NoneCmp CmpType = ""
)

type Client struct {
	ID               uuid.UUID
	Name             string
	Surname          string
	Patronymic       string
	Gender           Gender
	BirthDate        time.Time
	RegistrationDate time.Time
	Email            string
	Data             map[string]string
}

type User struct {
	ID       uuid.UUID
	Login    string
	Password string
	IsAdmin  bool
}

type Event struct {
	ID        uuid.UUID
	ClientID  uuid.UUID
	Alias     string
	EventTime time.Time
}

type EventType struct {
	ID    uuid.UUID
	Name  string
	Alias string
}

type EventFilter struct {
	Alias string
	Span  time.Duration
	Rate  int
}

type FieldFilter struct {
	Field  FieldType
	Cmp    CmpType
	Value1 string
	Value2 string
}

type Filter struct {
	Type        FilterType
	EventFilter *EventFilter
	FieldFilter *FieldFilter
}

type Ad struct {
	ID         uuid.UUID
	Content    string
	Filters    []Filter
	UserID     uuid.UUID
	Schedule   *Schedule
	CreateTime time.Time
}

type AdSendTime struct {
	ID   uuid.UUID
	Time time.Time
}

type AdSendStat struct {
	Num  int
	Date time.Time
}

type Schedule struct {
	ID       uuid.UUID
	Finished bool
	Periodic bool
	NextTime time.Time
	Span     time.Duration
}

type ClientStat struct {
	Num    int
	AvgAge int
}

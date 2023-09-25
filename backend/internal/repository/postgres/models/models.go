package models

import (
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `gorm:"primaryKey,autoIncrement"`
	Login    string
	Password string
	IsAdmin  bool
}

type Client struct {
	UUID             uuid.UUID `gorm:"primaryKey,autoIncrement"`
	Name             string
	Surname          string
	Patronymic       string
	Gender           models.Gender
	BirthDate        time.Time
	RegistrationDate time.Time
	Email            string
	Data             string
}

type EventType struct {
	UUID  uuid.UUID `gorm:"primaryKey,autoIncrement"`
	Name  string
	Alias string
}

type Event struct {
	UUID     uuid.UUID `gorm:"primaryKey,autoIncrement"`
	ClientID uuid.UUID
	Alias    string
	Time     time.Time
}

type Ad struct {
	UUID         uuid.UUID `gorm:"primaryKey,autoIncrement"`
	Content      string
	Filters      string
	UserID       uuid.UUID
	ScheduleID   uuid.UUID
	CreationTime time.Time
}

type AdWithSchedule struct {
	UUID         uuid.UUID `gorm:"primaryKey,autoIncrement"`
	Content      string
	Filters      string
	UserID       uuid.UUID
	NextTime     time.Time
	ScheduleID   uuid.UUID
	Finished     bool
	Periodic     bool
	Span         string
	CreationTime time.Time
}

type Schedule struct {
	UUID     uuid.UUID `gorm:"primaryKey,autoIncrement"`
	Finished bool
	Periodic bool
	NextTime time.Time
	Span     string
}

type ClientStat struct {
	Num    int
	AvgAge int
}

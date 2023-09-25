package interfaces

import (
	"context"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/google/uuid"
)

type EventCreationRepository interface {
	Create(context.Context, *models.Event) error
}

type AdSendTimeRepository interface {
	Create(context.Context, *models.AdSendTime) error
}

type AdSendDatesRepository interface {
	GetAdSendDates(ctx context.Context, lim int) ([]*models.AdSendStat, error)
}

type EventFilteringRepository interface {
	Filter(context.Context, []*models.EventFilter) ([]uuid.UUID, error)
}

type AdQueueSendingRepository interface {
	SendToQueue(_ context.Context, ads []*models.Ad) error
}

type ClientRepository interface {
	Create(context.Context, *models.Client) error
	Delete(context.Context, uuid.UUID) error
	GetAll(context.Context) ([]*models.Client, error)
	GetByID(context.Context, uuid.UUID) (*models.Client, error)
	GetByIDs(context.Context, []uuid.UUID) ([]*models.Client, error)
	Filter(context.Context, []*models.FieldFilter) ([]uuid.UUID, error)
	GetStats(context.Context) (*models.ClientStat, error)
}

type EventTypeRepository interface {
	GetByAlias(context.Context, string) (*models.EventType, error)
	Create(context.Context, *models.EventType) error
	GetAll(ctx context.Context) ([]*models.EventType, error)
}

type AdRepository interface {
	Create(context.Context, *models.Ad) error
	GetAll(ctx context.Context) ([]*models.Ad, error)
	GetBySpan(context.Context, time.Time) ([]*models.Ad, error)
}

type ScheduleRepository interface {
	Create(context.Context, *models.Schedule) (*models.Schedule, error)
	Update(ctx context.Context, sch *models.Schedule) error
}

type UserRepository interface {
	Create(context.Context, *models.User) (*models.User, error)
	Delete(context.Context, string) error
	Update(ctx context.Context, user *models.User) error
	GetAll(context.Context) ([]*models.User, error)
	GetByID(context.Context, uuid.UUID) (*models.User, error)
	GetByLogin(context.Context, string) (*models.User, error)
}

type SenderRepository interface {
	Send(context.Context, []*models.Client, string) error
}

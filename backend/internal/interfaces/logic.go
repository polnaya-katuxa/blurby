package interfaces

import (
	"context"
	"time"

	"github.com/google/uuid"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
)

// TODO
type ProfileLogic interface {
	AuthByToken(ctx context.Context, token string) (*models.User, error)
	Register(ctx context.Context, user *models.User, password string) (string, error)
	Login(ctx context.Context, login string, password string) (string, error)
	GetByLogin(ctx context.Context, login string) (*models.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	Delete(ctx context.Context, login string) error
	GetAll(ctx context.Context) ([]*models.User, error)
	GrantAdmin(ctx context.Context, login string) error
}

// TODO
type AdPlanner interface {
	Plan(ctx context.Context, t time.Time, span time.Duration) error
}

// TODO
type AdLogic interface {
	Create(ctx context.Context, add *models.Ad) error
	GetAll(ctx context.Context) ([]*models.Ad, error)
	GetBySpan(ctx context.Context, to time.Time) ([]*models.Ad, error)
}

type ClientLogic interface {
	Create(ctx context.Context, client *models.Client) error
	Get(ctx context.Context, id uuid.UUID) (*models.Client, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*models.Client, error)
}

// TODO
type EventLogic interface {
	Create(ctx context.Context, event *models.Event) error
}

type StatsLogic interface {
	GetClientStats(ctx context.Context) (*models.ClientStat, error)
	GetAdStats(ctx context.Context) ([]*models.AdSendStat, error)
}

type EventTypeLogic interface {
	Create(ctx context.Context, et *models.EventType) error
	GetAll(ctx context.Context) ([]*models.EventType, error)
}

// TODO
type FilterLogic interface {
	Filter(ctx context.Context, filters []models.Filter) ([]uuid.UUID, error)
}

// TODO
type SenderLogic interface {
	Send(ctx context.Context, ad *models.Ad) error
}

type GenerateLogic interface {
	Generate(ctx context.Context) error
}

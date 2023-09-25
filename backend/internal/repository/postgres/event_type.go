package postgres

import (
	"context"
	"errors"
	"fmt"

	my_errors "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventTypeRepository struct {
	dbT *gorm.DB
}

func NewETR(db *gorm.DB) *EventTypeRepository {
	return &EventTypeRepository{dbT: db}
}

func (et *EventTypeRepository) GetByAlias(ctx context.Context, alias string) (*models.EventType, error) {
	eventType := repoModels.EventType{}

	res := et.dbT.WithContext(ctx).Table("event_types").Where("alias = ?", alias).Take(&eventType)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("select: %w", my_errors.ErrNotFound)
	}
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	resET := models.EventType{
		ID:    eventType.UUID,
		Name:  eventType.Name,
		Alias: eventType.Alias,
	}

	return &resET, nil
}

func (et *EventTypeRepository) GetAll(ctx context.Context) ([]*models.EventType, error) {
	var types []*repoModels.EventType

	res := et.dbT.WithContext(ctx).Table("event_types").Find(&types)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	typesLogic := make([]*models.EventType, 0, len(types))
	for _, t := range types {
		typesLogic = append(typesLogic, &models.EventType{
			ID:    t.UUID,
			Name:  t.Name,
			Alias: t.Alias,
		})
	}

	return typesLogic, nil
}

func (et *EventTypeRepository) Create(ctx context.Context, eventType *models.EventType) error {
	eventTypeDB := &repoModels.EventType{
		UUID:  uuid.New(),
		Name:  eventType.Name,
		Alias: eventType.Alias,
	}

	res := et.dbT.WithContext(ctx).Table("event_types").Create(eventTypeDB)
	if res.Error != nil {
		return fmt.Errorf("insert: %w", res.Error)
	}

	return nil
}

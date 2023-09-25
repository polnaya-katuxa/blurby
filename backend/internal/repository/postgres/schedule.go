package postgres

import (
	"context"
	"fmt"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	dbT *gorm.DB
}

func NewSR(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{dbT: db}
}

func (s *ScheduleRepository) Create(ctx context.Context, sch *models.Schedule) (*models.Schedule, error) {
	schDB := &repoModels.Schedule{
		UUID:     uuid.New(),
		Finished: sch.Finished,
		Periodic: sch.Periodic,
		NextTime: sch.NextTime,
		Span:     sch.Span.String(),
	}

	res := s.dbT.WithContext(ctx).Table("schedules").Create(schDB)
	if res.Error != nil {
		return nil, fmt.Errorf("insert: %w", res.Error)
	}

	resSch := &models.Schedule{
		ID:       schDB.UUID,
		Finished: schDB.Finished,
		Periodic: sch.Periodic,
		NextTime: schDB.NextTime,
		Span:     sch.Span,
	}

	return resSch, nil
}

func (s *ScheduleRepository) Update(ctx context.Context, sch *models.Schedule) error {
	schDB := &repoModels.Schedule{
		UUID:     sch.ID,
		Finished: sch.Finished,
		Periodic: sch.Periodic,
		NextTime: sch.NextTime,
		Span:     sch.Span.String(),
	}

	res := s.dbT.WithContext(ctx).Table("schedules").Where("uuid = ?", sch.ID).Save(schDB)
	if res.Error != nil {
		return fmt.Errorf("update: %w", res.Error)
	}

	return nil
}

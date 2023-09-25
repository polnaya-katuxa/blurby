package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdRepository struct {
	dbT *gorm.DB
}

func NewAR(db *gorm.DB) *AdRepository {
	return &AdRepository{dbT: db}
}

func (a *AdRepository) Create(ctx context.Context, ad *models.Ad) error {
	filters, err := json.Marshal(ad.Filters)
	if err != nil {
		return fmt.Errorf("json filters marshall: %w", err)
	}

	adDB := &repoModels.Ad{
		UUID:         uuid.New(),
		Content:      ad.Content,
		Filters:      string(filters),
		UserID:       ad.UserID,
		ScheduleID:   ad.Schedule.ID,
		CreationTime: ad.CreateTime,
	}

	res := a.dbT.WithContext(ctx).Table("ads").Create(adDB)
	if res.Error != nil {
		return fmt.Errorf("insert: %w", res.Error)
	}

	return nil
}

func (a *AdRepository) GetBySpan(ctx context.Context, to time.Time) ([]*models.Ad, error) {
	var adsDB []*repoModels.AdWithSchedule

	res := a.dbT.WithContext(ctx).Raw("select * from get_ads_by_span(?)", to).Scan(&adsDB)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	adsLogic := make([]*models.Ad, 0, len(adsDB))
	for _, a := range adsDB {
		var f []models.Filter
		if a.Filters != "" {
			err := json.Unmarshal([]byte(a.Filters), &f)
			if err != nil {
				return nil, fmt.Errorf("json unmarshal filters: %w", err)
			}
		}

		span, err := time.ParseDuration(a.Span)
		if err != nil {
			return nil, fmt.Errorf("parse duration: %w", err)
		}

		adsLogic = append(adsLogic, &models.Ad{
			ID:      a.UUID,
			Content: a.Content,
			Filters: f,
			UserID:  a.UserID,
			Schedule: &models.Schedule{
				ID:       a.ScheduleID,
				Finished: a.Finished,
				Periodic: a.Periodic,
				NextTime: a.NextTime,
				Span:     span,
			},
			CreateTime: a.CreationTime,
		})
	}

	return adsLogic, nil
}

func (a *AdRepository) GetAll(ctx context.Context) ([]*models.Ad, error) {
	var adsDB []*repoModels.AdWithSchedule

	res := a.dbT.WithContext(ctx).Table("ads").Select("ads.uuid, ads.content, ads.filters, ads.user_id," +
		"ads.schedule_id, ads.creation_time, schedules.finished, schedules.periodic, schedules.next_time," +
		"schedules.span").Joins("left join schedules on ads.schedule_id =" +
		"schedules.uuid").Where("schedules.finished = false").Find(&adsDB)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	adsLogic := make([]*models.Ad, 0, len(adsDB))
	for _, a := range adsDB {
		var f []models.Filter
		if a.Filters != "" {
			err := json.Unmarshal([]byte(a.Filters), &f)
			if err != nil {
				return nil, fmt.Errorf("json unmarshal filters: %w", err)
			}
		}

		span, err := time.ParseDuration(a.Span)
		if err != nil {
			return nil, fmt.Errorf("parse duration: %w", err)
		}

		adsLogic = append(adsLogic, &models.Ad{
			ID:      a.UUID,
			Content: a.Content,
			Filters: f,
			UserID:  a.UserID,
			Schedule: &models.Schedule{
				ID:       a.ScheduleID,
				Finished: a.Finished,
				Periodic: a.Periodic,
				NextTime: a.NextTime,
				Span:     span,
			},
			CreateTime: a.CreationTime,
		})
	}

	return adsLogic, nil
}

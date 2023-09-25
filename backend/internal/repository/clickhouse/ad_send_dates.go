package clickhouse

import (
	"context"
	"fmt"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"gorm.io/gorm"
)

type AdSendDatesRepository struct {
	db *gorm.DB
}

func NewASDR(db *gorm.DB) *AdSendDatesRepository {
	return &AdSendDatesRepository{db: db}
}

func (ef *AdSendDatesRepository) GetAdSendDates(ctx context.Context, lim int) ([]*models.AdSendStat, error) {
	var stats []*models.AdSendStat
	res := ef.db.WithContext(ctx).Raw("select count(*) as num, date "+
		"from coursework.ad_send_times group by date order by date limit ?", lim).Find(&stats)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	return stats, nil
}

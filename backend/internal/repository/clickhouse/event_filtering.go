package clickhouse

import (
	"context"
	"fmt"
	"strings"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventFilteringRepository struct {
	db *gorm.DB
}

func NewCHEFR(db *gorm.DB) *EventFilteringRepository {
	return &EventFilteringRepository{db: db}
}

func (ef *EventFilteringRepository) Filter(ctx context.Context, f []*models.EventFilter) ([]uuid.UUID, error) {
	if len(f) == 0 {
		var uuidsStr []string
		res := ef.db.WithContext(ctx).Table("events").Distinct("client_id").Find(&uuidsStr)
		if res.Error != nil {
			return nil, fmt.Errorf("select: %w", res.Error)
		}

		uuids := make([]uuid.UUID, 0, len(uuidsStr))
		for _, u := range uuidsStr {
			uid, err := uuid.Parse(u)
			if err != nil {
				return nil, fmt.Errorf("parse uuid: %w", err)
			}

			uuids = append(uuids, uid)
		}

		return uuids, nil
	}

	query := make([]interface{}, len(f))

	for i, v := range f {
		if v.Rate < 1 {
			v.Rate = 1
		}

		if v.Span != time.Duration(0) {
			query[i] = ef.db.Raw(`
				select distinct client_id
				from (
					 select
						 client_id,
						 alias,
						 count(*) as num
					 from coursework.events
					 where time between now() - interval ? hour and now()
					 group by client_id, alias
					 order by num desc
					 limit ? by client_id
				) where alias = ?
			`, int(v.Span.Hours()), v.Rate, v.Alias)
		} else {
			query[i] = ef.db.Raw(`
				select distinct client_id
				from (
					 select
						 client_id,
						 alias,
						 count(*) as num
					 from coursework.events
					 group by client_id, alias
					 order by num desc
					 limit ? by client_id
				) where alias = ?
			`, v.Rate, v.Alias)
		}
	}

	questions := make([]string, len(f))
	for i := range questions {
		questions[i] = "?"
	}
	intersectionQuery := strings.Join(questions, " intersect ")

	var uuidsStr []string
	res := ef.db.Debug().WithContext(ctx).Raw(intersectionQuery, query...).Find(&uuidsStr)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	uuids := make([]uuid.UUID, 0, len(uuidsStr))
	for _, u := range uuidsStr {
		uid, err := uuid.Parse(u)
		if err != nil {
			return nil, fmt.Errorf("parse uuid: %w", err)
		}

		uuids = append(uuids, uid)
	}

	return uuids, nil
}

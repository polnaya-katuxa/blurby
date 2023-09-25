package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	repoModels "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/postgres/models"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventFilteringRepository struct {
	dbT *gorm.DB
	dbC *gorm.DB
}

func NewPGEFR(dbT *gorm.DB, dbC *gorm.DB) *EventFilteringRepository {
	return &EventFilteringRepository{
		dbT: dbT,
		dbC: dbC,
	}
}

func (ef *EventFilteringRepository) Create(ctx context.Context, event *models.Event) error {
	eventDB := &repoModels.Event{
		UUID:     uuid.New(),
		ClientID: event.ClientID,
		Alias:    event.Alias,
		Time:     event.EventTime,
	}

	res := ef.dbC.WithContext(ctx).Table("events").Create(eventDB)
	if res.Error != nil {
		return fmt.Errorf("insert: %w", res.Error)
	}

	return nil
}

func (ef *EventFilteringRepository) Filter(ctx context.Context, f []*models.EventFilter) ([]uuid.UUID, error) {
	if len(f) == 0 {
		var uuidsStr []string
		res := ef.dbT.WithContext(ctx).Table("events").Distinct("client_id").Find(&uuidsStr)
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
			query[i] = ef.dbT.Raw(`
		select distinct cast sub1.client_id
			from (
			select distinct sub.client_id, sub.alias, sub.c,
			       row_number() over (partition by sub.client_id order by sub.c desc) as r
				from (
				select client_id, alias, count(*) as c from events
				    where time between current_timestamp - '`+fmt.Sprint(int(v.Span.Hours()))+` hours'::interval and current_timestamp
					group by events.client_id, alias) as sub
				) as sub1
					where alias = ? and r <= ?`, v.Alias, v.Rate)
		} else {
			query[i] = ef.dbT.Raw(`
		select distinct sub1.client_id
			from (
			select sub.client_id, sub.alias, sub.c,
			       row_number() over (partition by sub.client_id order by sub.c desc) as r
				from (
				select client_id, alias, count(*) as c from events
					group by events.client_id, alias) as sub
				) as sub1
					where alias = ? and r <= ?`, v.Alias, v.Rate)
		}
	}

	questions := make([]string, len(f))
	for i := range questions {
		questions[i] = "?"
	}
	intersectionQuery := strings.Join(questions, " intersect ")

	var uuidsStr []string
	res := ef.dbT.WithContext(ctx).Raw(intersectionQuery, query...).Find(&uuidsStr)
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

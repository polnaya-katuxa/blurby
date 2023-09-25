package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"

	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/google/uuid"
)

type FilterLogic struct {
	cr interfaces.ClientRepository
	er interfaces.EventFilteringRepository
}

func NewFL(cr interfaces.ClientRepository, er interfaces.EventFilteringRepository) *FilterLogic {
	return &FilterLogic{
		cr: cr,
		er: er,
	}
}

func (fl *FilterLogic) Filter(ctx context.Context, filters []models.Filter) ([]uuid.UUID, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started filtering")

	fieldFilters := make([]*models.FieldFilter, 0, len(filters))
	eventFilters := make([]*models.EventFilter, 0, len(filters))

	for _, f := range filters {
		if f.Type == models.ByField {
			fieldFilters = append(fieldFilters, f.FieldFilter)
		} else if f.Type == models.ByEvent {
			eventFilters = append(eventFilters, f.EventFilter)
		}
	}

	fieldID, err := fl.cr.Filter(ctx, fieldFilters)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot filter by field", "error", err)
		return nil, fmt.Errorf("field filter: %w", err)
	}

	eventID, err := fl.er.Filter(ctx, eventFilters)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot filter by event", "error", err)
		return nil, fmt.Errorf("event filter: %w", err)
	}

	if len(fieldFilters)+len(eventFilters) == 0 {
		unionSet := sets.New[uuid.UUID](fieldID...)
		unionSet = unionSet.Insert(eventID...)

		return unionSet.UnsortedList(), nil
	}

	if len(fieldFilters) == 0 {
		mycontext.LoggerFromContext(ctx).Debugw("filtered", "ids", eventID)
		return eventID, nil
	}

	if len(eventFilters) == 0 {
		mycontext.LoggerFromContext(ctx).Debugw("filtered", "ids", fieldID)
		return fieldID, nil
	}

	fieldSet := sets.New[uuid.UUID](fieldID...)
	eventSet := sets.New[uuid.UUID](eventID...)
	ids := fieldSet.Intersection(eventSet).UnsortedList()

	mycontext.LoggerFromContext(ctx).Debugw("filtered", "ids", ids)

	return ids, nil
}

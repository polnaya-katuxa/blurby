package controllers

import (
	"context"
	"fmt"
	"time"

	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
)

type Event struct {
	eventCreationRepo interfaces.EventCreationRepository
	clientRepo        interfaces.ClientRepository
	eventTypeRepo     interfaces.EventTypeRepository
}

func NewEL(cr interfaces.ClientRepository, ecr interfaces.EventCreationRepository,
	etr interfaces.EventTypeRepository,
) *Event {
	return &Event{
		eventCreationRepo: ecr,
		clientRepo:        cr,
		eventTypeRepo:     etr,
	}
}

func (e *Event) Create(ctx context.Context, event *models.Event) error {
	mycontext.LoggerFromContext(ctx).Debugw("started creating event")

	client, err := e.clientRepo.GetByID(ctx, event.ClientID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get client by id", "error", err)
		return fmt.Errorf("get client by id: %w", err)
	}

	if client == nil {
		mycontext.LoggerFromContext(ctx).Warnw("non existing client", "error", errors.ErrExist)
		return fmt.Errorf("get client by id: %w", errors.ErrExist)
	}

	t, err := e.eventTypeRepo.GetByAlias(ctx, event.Alias)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get event type by alias", "error", err)
		return fmt.Errorf("get event type by id: %w", err)
	}

	if t == nil {
		mycontext.LoggerFromContext(ctx).Warnw("non existing event type", "error", errors.ErrExist)
		return fmt.Errorf("get event type by id: %w", errors.ErrExist)
	}

	event.EventTime = time.Now().UTC()

	if err = e.eventCreationRepo.Create(ctx, event); err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot create event", "error", err)
		return fmt.Errorf("create: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("created event")

	return nil
}

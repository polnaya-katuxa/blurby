package controllers

import (
	"context"
	"fmt"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	myerrors "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
)

type EventType struct {
	eventTypeRepo interfaces.EventTypeRepository
}

func NewETL(etr interfaces.EventTypeRepository) *EventType {
	return &EventType{
		eventTypeRepo: etr,
	}
}

func (e *EventType) Create(ctx context.Context, et *models.EventType) error {
	mycontext.LoggerFromContext(ctx).Debugw("started creating event type")

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to create event type", "error", err)
		return fmt.Errorf("user from context: %w", err)
	}

	if user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("not a targetologist tried to create event type", "error",
			myerrors.ErrAdmin, "user", user.Login)
		return fmt.Errorf("user is: %w", myerrors.ErrAdmin)
	}

	err = e.eventTypeRepo.Create(ctx, et)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot create event type", "error", err)
		return fmt.Errorf("create: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("created event type")

	return nil
}

func (e *EventType) GetAll(ctx context.Context) ([]*models.EventType, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started getting all event types")

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get all event types", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	if user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("not a targetologist tried to get all event types", "error", myerrors.ErrAdmin, "user", user.Login)
		return nil, fmt.Errorf("user is: %w", myerrors.ErrAdmin)
	}

	ets, err := e.eventTypeRepo.GetAll(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get all event types", "error", err, "user", user.Login)
		return nil, fmt.Errorf("get all: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("got all event types")

	return ets, nil
}

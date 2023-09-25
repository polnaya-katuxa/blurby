package controllers

import (
	"context"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	myerrors "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
)

type Ad struct {
	adRepository       interfaces.AdRepository
	scheduleRepository interfaces.ScheduleRepository
}

func NewAL(ar interfaces.AdRepository, sr interfaces.ScheduleRepository) *Ad {
	return &Ad{
		adRepository:       ar,
		scheduleRepository: sr,
	}
}

func (a *Ad) Create(ctx context.Context, ad *models.Ad) error {
	mycontext.LoggerFromContext(ctx).Debugw("started creating ad")

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to create ad", "error", err)
		return fmt.Errorf("user from context: %w", err)
	}

	if user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("not a targetologist tried to create ad", "error",
			myerrors.ErrAdmin, "user", user.Login)
		return fmt.Errorf("user is: %w", myerrors.ErrAdmin)
	}

	ad.UserID = user.ID
	ad.CreateTime = time.Now().UTC()

	ad.Schedule.NextTime = ad.CreateTime.Add(ad.Schedule.Span).UTC()

	ad.Schedule, err = a.scheduleRepository.Create(ctx, ad.Schedule)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot create schedule", "error", err)
		return fmt.Errorf("create schedule: %w", err)
	}

	err = a.adRepository.Create(ctx, ad)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot create ad", "error", err)
		return fmt.Errorf("create ad: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("created ad")

	return nil
}

func (a *Ad) GetAll(ctx context.Context) ([]*models.Ad, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started getting all ads")

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get all ads", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	if user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("not a targetologist tried to get all ads", "error",
			myerrors.ErrAdmin, "user", user.Login)
		return nil, fmt.Errorf("user is: %w", myerrors.ErrAdmin)
	}

	ads, err := a.adRepository.GetAll(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get all ads", "error", err)
		return nil, fmt.Errorf("get all ads: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("got all ads")

	return ads, nil
}

func (a *Ad) GetBySpan(ctx context.Context, to time.Time) ([]*models.Ad, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started getting ads by span")

	ads, err := a.adRepository.GetBySpan(ctx, to)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get ads by span", "error", err)
		return nil, fmt.Errorf("get ads by span: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("got ads by span")

	return ads, nil
}

package controllers

import (
	"context"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
)

type AdPlanner struct {
	adLogic  interfaces.AdLogic
	adSender interfaces.AdQueueSendingRepository
	schRepo  interfaces.ScheduleRepository
}

func NewAP(al interfaces.AdLogic, as interfaces.AdQueueSendingRepository, sr interfaces.ScheduleRepository) *AdPlanner {
	return &AdPlanner{
		adLogic:  al,
		adSender: as,
		schRepo:  sr,
	}
}

func (ap *AdPlanner) Plan(ctx context.Context, t time.Time, span time.Duration) error {
	mycontext.LoggerFromContext(ctx).Debugw("started planning ads")

	ads, err := ap.adLogic.GetBySpan(ctx, t.UTC().Add(span))
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get ads by span", "error", err)
		return fmt.Errorf("get ads by span: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("got ads", "ads", ads)

	err = ap.adSender.SendToQueue(ctx, ads)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot send ads to queue", "error", err)
		return fmt.Errorf("send to queue: %w", err)
	}

	for _, a := range ads {
		if !a.Schedule.Periodic {
			a.Schedule.Finished = true

			err = ap.schRepo.Update(ctx, a.Schedule)
			if err != nil {
				mycontext.LoggerFromContext(ctx).Warnw("cannot update schedule", "error", err)
				return fmt.Errorf("update schedule: %w", err)
			}
		} else {
			a.Schedule.NextTime = a.Schedule.NextTime.Add(a.Schedule.Span)

			err = ap.schRepo.Update(ctx, a.Schedule)
			if err != nil {
				mycontext.LoggerFromContext(ctx).Warnw("cannot update schedule", "error", err)
				return fmt.Errorf("update schedule: %w", err)
			}
		}
	}

	mycontext.LoggerFromContext(ctx).Debugw("planned ads")

	return nil
}

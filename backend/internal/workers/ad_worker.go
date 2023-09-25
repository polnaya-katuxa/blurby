package workers

import (
	"context"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	"go.uber.org/zap"
)

type AdWorker struct {
	AdPlanner interfaces.AdPlanner
	logger    *zap.SugaredLogger

	span time.Duration
}

func New(ap interfaces.AdPlanner, span time.Duration, logger *zap.SugaredLogger) *AdWorker {
	return &AdWorker{
		AdPlanner: ap,
		logger:    logger,
		span:      span,
	}
}

func (aw *AdWorker) Run(ctx context.Context) error {
	aw.logger.Debugw("started working on ad planning")

	timer := time.NewTimer(aw.span)

	for t := range timer.C {
		timer.Reset(aw.span)

		err := aw.AdPlanner.Plan(ctx, t, aw.span)
		if err != nil {
			aw.logger.Warnw("cannot plan an ad", "error", err)
		}

		aw.logger.Debugw("planned an ad")
	}

	aw.logger.Debugw("stopped working on ad planning")

	return nil
}

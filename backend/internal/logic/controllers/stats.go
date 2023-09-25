package controllers

import (
	"context"
	"fmt"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
)

type StatsLogic struct {
	cr   interfaces.ClientRepository
	ar   interfaces.AdRepository
	asdr interfaces.AdSendDatesRepository

	lim int
}

func NewStatL(cr interfaces.ClientRepository, ar interfaces.AdRepository, asdr interfaces.AdSendDatesRepository, lim int) *StatsLogic {
	return &StatsLogic{
		cr:   cr,
		ar:   ar,
		asdr: asdr,
		lim:  lim,
	}
}

func (s *StatsLogic) GetClientStats(ctx context.Context) (*models.ClientStat, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started getting client stats")

	_, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get client stats", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	stats, err := s.cr.GetStats(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get client stats", "error", err)
		return nil, fmt.Errorf("get: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("got client stats")

	return stats, nil
}

func (s *StatsLogic) GetAdStats(ctx context.Context) ([]*models.AdSendStat, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started getting ad stats")

	_, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get ad stats", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	stats, err := s.asdr.GetAdSendDates(ctx, s.lim)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get ad stats", "error", err)
		return nil, fmt.Errorf("get: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("got ad stats")

	return stats, nil
}

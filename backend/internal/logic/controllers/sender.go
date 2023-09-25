package controllers

import (
	"context"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
)

type SenderLogic struct {
	fl   interfaces.FilterLogic
	cr   interfaces.ClientRepository
	mr   interfaces.SenderRepository
	astr interfaces.AdSendTimeRepository
}

func NewSL(fl interfaces.FilterLogic, cr interfaces.ClientRepository, mr interfaces.SenderRepository,
	astr interfaces.AdSendTimeRepository,
) *SenderLogic {
	return &SenderLogic{
		fl:   fl,
		cr:   cr,
		mr:   mr,
		astr: astr,
	}
}

func (s *SenderLogic) Send(ctx context.Context, ad *models.Ad) error {
	mycontext.LoggerFromContext(ctx).Debugw("started sending mail", "filters", ad.Filters)

	ids, err := s.fl.Filter(ctx, ad.Filters)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot filter", "filters", ad.Filters)
		return fmt.Errorf("filter: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("filtered ids", "ids", ids)

	clients, err := s.cr.GetByIDs(ctx, ids)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get by ids", "ids", ids)
		return fmt.Errorf("get by ids: %w", err)
	}

	err = s.mr.Send(ctx, clients, ad.Content)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot send mail", "ids", ids)
		return fmt.Errorf("send: %w", err)
	}

	err = s.astr.Create(ctx, &models.AdSendTime{
		Time: time.Now().UTC(),
	})
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot create ad send time", "ids", ids)
		return fmt.Errorf("create: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("sent mail", "ids", ids)

	return nil
}

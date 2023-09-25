package gomail

import (
	"context"
	"fmt"

	context2 "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"gopkg.in/gomail.v2"
)

type Sender struct {
	dialer *gomail.Dialer
	email  string
}

func NewS(dialer *gomail.Dialer, email string) *Sender {
	return &Sender{dialer: dialer, email: email}
}

func (s *Sender) Send(ctx context.Context, clients []*models.Client, content string) error {
	for _, c := range clients {
		msg := gomail.NewMessage()
		msg.SetHeader("From", s.email)
		msg.SetHeader("To", c.Email)
		msg.SetHeader("Subject", "Advert")
		msg.SetBody("text/html", content)

		context2.LoggerFromContext(ctx).Debugw("sending email", "to", c.Email)

		if err := s.dialer.DialAndSend(msg); err != nil {
			context2.LoggerFromContext(ctx).Errorw("cant send email", "to", c.Email, "err", err)
			return fmt.Errorf("dial and send: %w", err)
		}
	}

	return nil
}

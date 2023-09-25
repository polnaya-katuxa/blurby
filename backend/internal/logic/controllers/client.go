package controllers

import (
	"context"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	myerrors "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/google/uuid"
)

type Client struct {
	clientRepository interfaces.ClientRepository
}

func NewCL(cr interfaces.ClientRepository) *Client {
	return &Client{clientRepository: cr}
}

func (c *Client) Create(ctx context.Context, client *models.Client) error {
	mycontext.LoggerFromContext(ctx).Debugw("started creating client", "client_name",
		client.Name+client.Surname, "client_mail", client.Email)

	client.RegistrationDate = time.Now()

	err := c.clientRepository.Create(ctx, client)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot create client", "error", err)
		return fmt.Errorf("create: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("created client", "client_name",
		client.Name+client.Surname, "client_mail", client.Email)

	return nil
}

func (c *Client) Get(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started getting client", "client_id", id)

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get client", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	if !user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("not an admin tried to get client", "error", myerrors.ErrNotAdmin, "user", user.Login)
		return nil, fmt.Errorf("user is: %w", myerrors.ErrNotAdmin)
	}

	userRes, err := c.clientRepository.GetByID(ctx, id)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get client by id", "error", err, "client_id", id)
		return nil, fmt.Errorf("get: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("got client", "client_id", id)

	return userRes, nil
}

func (c *Client) Delete(ctx context.Context, id uuid.UUID) error {
	mycontext.LoggerFromContext(ctx).Debugw("started deleting client", "client_id", id)

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to delete client", "error", err,
			"client_id", id)
		return fmt.Errorf("user from context: %w", err)
	}

	if !user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("not an admin tried to delete client", "error", myerrors.ErrNotAdmin, "user", user.Login)
		return fmt.Errorf("user is: %w", myerrors.ErrNotAdmin)
	}

	if err := c.clientRepository.Delete(ctx, id); err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot delete client", "error", err, "client_id", id)
		return fmt.Errorf("delete: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("deleted client", "client_id", id)

	return nil
}

func (c *Client) GetAll(ctx context.Context) ([]*models.Client, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started getting all clients")

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get all clients", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	if !user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("not an admin tried to get all clients", "error", myerrors.ErrNotAdmin, "user", user.Login)
		return nil, fmt.Errorf("user is: %w", myerrors.ErrNotAdmin)
	}

	users, err := c.clientRepository.GetAll(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get all clients", "error", err, "user", user.Login)
		return nil, fmt.Errorf("get all: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("got all clients")

	return users, nil
}

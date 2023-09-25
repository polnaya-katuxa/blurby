package context

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	my_errors "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
)

type ctxKey string

const (
	key    ctxKey = "user"
	keyLog ctxKey = "logger"
)

func UserToContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, key, user)
}

func UserFromContext(ctx context.Context) (*models.User, error) {
	user := ctx.Value(key)
	if user == nil {
		return nil, fmt.Errorf("get user: %w", my_errors.ErrGet)
	}

	return user.(*models.User), nil
}

func LoggerToContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, keyLog, logger)
}

func LoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	v := ctx.Value(keyLog)
	if v == nil {
		return zap.NewNop().Sugar()
	}
	return v.(*zap.SugaredLogger)
}

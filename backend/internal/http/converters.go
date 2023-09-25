package http

import (
	"errors"
	"net/http"

	myerrors "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"
	openapi "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/server"
)

func toErrorResponse(err error, defaultMessage string) (openapi.ImplResponse, error) {
	switch {
	case errors.Is(err, myerrors.ErrGet):
		return openapi.ImplResponse{
			Code: http.StatusUnauthorized,
			Body: openapi.ErrorResponse{
				Message:       "Authorization failed.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrBadLogin):
		return openapi.ImplResponse{
			Code: http.StatusPreconditionFailed,
			Body: openapi.ErrorResponse{
				Message:       "Incorrect login length. Try again.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrBadPassword):
		return openapi.ImplResponse{
			Code: http.StatusPreconditionFailed,
			Body: openapi.ErrorResponse{
				Message:       "Incorrect password length. Try again.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrExistingLogin):
		return openapi.ImplResponse{
			Code: http.StatusConflict,
			Body: openapi.ErrorResponse{
				Message:       "User already exists.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrNotAdmin):
		return openapi.ImplResponse{
			Code: http.StatusForbidden,
			Body: openapi.ErrorResponse{
				Message:       "Operation can be proceeded only by an admin.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrNotFound),
		errors.Is(err, myerrors.ErrUserNotFound),
		errors.Is(err, myerrors.ErrExist):
		return openapi.ImplResponse{
			Code: http.StatusNotFound,
			Body: openapi.ErrorResponse{
				Message:       "Not found.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrAdmin):
		return openapi.ImplResponse{
			Code: http.StatusForbidden,
			Body: openapi.ErrorResponse{
				Message:       "Operation can be proceeded only by a targetologist.",
				SystemMessage: err.Error(),
			},
		}, nil
	case errors.Is(err, myerrors.ErrFilter):
		return openapi.ImplResponse{
			Code: http.StatusTeapot,
			Body: openapi.ErrorResponse{
				Message:       "Invalid filtering parameters.",
				SystemMessage: err.Error(),
			},
		}, nil
	default:
		return openapi.ImplResponse{
			Code: http.StatusInternalServerError,
			Body: openapi.ErrorResponse{
				Message:       defaultMessage,
				SystemMessage: err.Error(),
			},
		}, nil
	}
}

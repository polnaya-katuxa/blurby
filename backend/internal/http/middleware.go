package http

import (
	"net/http"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"

	"go.uber.org/zap"

	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
)

func Middleware(prl interfaces.ProfileLogic, logger *zap.SugaredLogger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user, err := prl.AuthByToken(ctx, r.Header.Get("User-Token"))
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx = mycontext.UserToContext(ctx, user)
		ctx = mycontext.LoggerToContext(ctx, logger)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

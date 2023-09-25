package controllers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	myerrors "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Profile struct {
	userRepository interfaces.UserRepository

	secretKey       string
	tokenExpiration time.Duration
}

func NewPL(ur interfaces.UserRepository, s string, te time.Duration) *Profile {
	return &Profile{
		userRepository:  ur,
		secretKey:       s,
		tokenExpiration: te,
	}
}

func (p *Profile) AuthByToken(ctx context.Context, token string) (*models.User, error) {
	mycontext.LoggerFromContext(ctx).Debugw("authentification started", "token", token)

	var claims jwt.RegisteredClaims
	_, err := jwt.ParseWithClaims(token, &claims, func(_ *jwt.Token) (interface{}, error) {
		return []byte(p.secretKey), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			mycontext.LoggerFromContext(ctx).Warnw("authentification failed: expired", "token", token, "error", err)
			return nil, fmt.Errorf("authentification failed: expired %w", err)
		} else {
			mycontext.LoggerFromContext(ctx).Errorw("authentification failed", "token", token, "error", err)
			return nil, fmt.Errorf("authentification failed: %w", err)
		}
	}

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Errorw("cannot parse token claims", "token", token, "error", err)
		return nil, fmt.Errorf("parse id: %w", err)
	}

	user, err := p.userRepository.GetByID(ctx, id)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Errorw("cannot get user by id", "token", token, "id", id, "error", err)
		return nil, fmt.Errorf("get: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("authentification succeeded", "user", user.Login)

	return user, nil
}

func (p *Profile) Register(ctx context.Context, user *models.User, password string) (string, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started registration", "user", user.Login)

	userOld, err := p.userRepository.GetByLogin(ctx, user.Login)

	if err != nil && !errors.Is(err, myerrors.ErrUserNotFound) {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user by login", "user", user.Login, "error", err)
		return "", fmt.Errorf("get by login: %w", err)
	}

	if userOld != nil {
		mycontext.LoggerFromContext(ctx).Warnw("user already exists", "user", user.Login)
		return "", fmt.Errorf("already exists: %w", myerrors.ErrExistingLogin)
	}

	if len(user.Login) < 4 {
		mycontext.LoggerFromContext(ctx).Errorw("incorrect login", "user", user.Login)
		return "", fmt.Errorf("login: %w", myerrors.ErrBadLogin)
	}

	if len(password) < 8 {
		mycontext.LoggerFromContext(ctx).Errorw("incorrect password", "user", user.Login)
		return "", fmt.Errorf("password: %w", myerrors.ErrBadPassword)
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hash)

	user, err = p.userRepository.Create(ctx, user)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Errorw("cannot create user", "user", user.Login, "error", err)
		return "", fmt.Errorf("register: %w", err)
	}

	ctx = mycontext.UserToContext(ctx, user)

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(p.tokenExpiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ID:        user.ID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(p.secretKey))

	mycontext.LoggerFromContext(ctx).Debugw("registration succeeded", "user", user.Login)

	return ss, nil
}

func (p *Profile) Login(ctx context.Context, login string, password string) (string, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started login", "user", login)

	user, err := p.userRepository.GetByLogin(ctx, login)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user by login", "user", login, "error", err)
		return "", fmt.Errorf("get: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		mycontext.LoggerFromContext(ctx).Errorw("login with wrong password", "user", login, "error", err)
		return "", fmt.Errorf("cmp hash & password: %w", err)
	}

	ctx = mycontext.UserToContext(ctx, user)

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(p.tokenExpiration)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ID:        user.ID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(p.secretKey))

	mycontext.LoggerFromContext(ctx).Debugw("successful login", "user", login)

	return ss, nil
}

func (p *Profile) GetByLogin(ctx context.Context, login string) (*models.User, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started getting user by login", "user", login)

	userRes, err := p.userRepository.GetByLogin(ctx, login)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user by login", "user", login, "error", err)
		return nil, fmt.Errorf("get: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("successfully got user by login", "user", login)

	return userRes, nil
}

func (p *Profile) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started getting user by id", "user", id)

	userRes, err := p.userRepository.GetByID(ctx, id)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user by id", "user", id, "error", err)
		return nil, fmt.Errorf("get: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("successfully got user by id", "user", id)

	return userRes, nil
}

func (p *Profile) Delete(ctx context.Context, login string) error {
	mycontext.LoggerFromContext(ctx).Debugw("started deleting user", "user to delete", login)

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to delete another", "user to delete", login, "error", err)
		return fmt.Errorf("user from context: %w", err)
	}

	if !user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("trying to delete user without admin perms", "user", user.Login, "deleted", login)
		return fmt.Errorf("user is: %w", myerrors.ErrNotAdmin)
	}

	if login == user.Login {
		mycontext.LoggerFromContext(ctx).Warnw("cannot delete yourself", "user", user.Login, "user to delete", login)
		return fmt.Errorf("delete: %w", myerrors.ErrAutoDelete)
	}

	if err := p.userRepository.Delete(ctx, login); err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot delete user", "user", user.Login, "user to delete", login, "error", err)
		return fmt.Errorf("delete: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("admin deleted user", "user", user.Login, "deleted", login)

	return nil
}

func (p *Profile) GetAll(ctx context.Context) ([]*models.User, error) {
	mycontext.LoggerFromContext(ctx).Debugw("started getting all users")

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to get all users", "error", err)
		return nil, fmt.Errorf("user from context: %w", err)
	}

	if !user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("trying to get all users without admin perms", "user", user.Login)
		return nil, fmt.Errorf("user is: %w", myerrors.ErrNotAdmin)
	}

	users, err := p.userRepository.GetAll(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get all users", "user", user.Login, "error", err)
		return nil, fmt.Errorf("get all: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("admin got all users", "user", user.Login)

	return users, nil
}

func (p *Profile) GrantAdmin(ctx context.Context, login string) error {
	mycontext.LoggerFromContext(ctx).Debugw("started granting admin to user", "to", login)

	user, err := mycontext.UserFromContext(ctx)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user to grant admin", "error", err, "to", login)
		return fmt.Errorf("user from context: %w", err)
	}

	if !user.IsAdmin {
		mycontext.LoggerFromContext(ctx).Errorw("trying to grant admin without admin perms", "user", user.Login, "to", login)
		return fmt.Errorf("user is: %w", myerrors.ErrNotAdmin)
	}

	newAdmin, err := p.userRepository.GetByLogin(ctx, login)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot get user by login", "user", login, "error", err)
		return fmt.Errorf("grant: %w", err)
	}

	newAdmin.IsAdmin = true

	err = p.userRepository.Update(ctx, newAdmin)
	if err != nil {
		mycontext.LoggerFromContext(ctx).Warnw("cannot grant admin", "user", user.Login, "to", login, "error", err)
		return fmt.Errorf("grant: %w", err)
	}

	mycontext.LoggerFromContext(ctx).Debugw("admin granted admin to user", "user", user.Login, "to", login)

	return nil
}

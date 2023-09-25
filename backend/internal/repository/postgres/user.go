package postgres

import (
	"context"
	"errors"
	"fmt"

	my_errors "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"

	logicModels "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	repoModels "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/postgres/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	dbA *gorm.DB
	dbT *gorm.DB
}

func NewUR(dbA *gorm.DB, dbT *gorm.DB) *UserRepository {
	return &UserRepository{
		dbA: dbA,
		dbT: dbT,
	}
}

func (u *UserRepository) GetByLogin(ctx context.Context, login string) (*logicModels.User, error) {
	user := repoModels.User{}

	res := u.dbT.WithContext(ctx).Table("users").Where("login = ?", login).Take(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("select: %w", my_errors.ErrUserNotFound)
	}
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	resUser := logicModels.User{
		ID:       user.UUID,
		Login:    user.Login,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}

	return &resUser, nil
}

func (u *UserRepository) GetByID(ctx context.Context, uuid uuid.UUID) (*logicModels.User, error) {
	user := repoModels.User{}

	res := u.dbT.WithContext(ctx).Table("users").Where("uuid = ?", uuid).Take(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("select: %w", my_errors.ErrUserNotFound)
	}
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	resUser := logicModels.User{
		ID:       user.UUID,
		Login:    user.Login,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}

	return &resUser, nil
}

func (u *UserRepository) Delete(ctx context.Context, login string) error {
	res := u.dbA.WithContext(ctx).Table("users").Where("login = ?", login).Delete(&repoModels.User{})
	if res.Error != nil {
		return fmt.Errorf("delete: %w", res.Error)
	}

	return nil
}

func (u *UserRepository) Update(ctx context.Context, user *logicModels.User) error {
	userDB := &repoModels.User{
		UUID:     user.ID,
		Login:    user.Login,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}

	res := u.dbA.WithContext(ctx).Table("users").Where("uuid = ?", userDB.UUID).Save(userDB)
	if res.Error != nil {
		return fmt.Errorf("update: %w", res.Error)
	}

	return nil
}

func (u *UserRepository) GetAll(ctx context.Context) ([]*logicModels.User, error) {
	var usersDB []*repoModels.User

	res := u.dbT.WithContext(ctx).Table("users").Find(&usersDB)
	if res.Error != nil {
		return nil, fmt.Errorf("select: %w", res.Error)
	}

	usersLogic := make([]*logicModels.User, 0, len(usersDB))
	for _, u := range usersDB {
		usersLogic = append(usersLogic, &logicModels.User{
			ID:       u.UUID,
			Login:    u.Login,
			Password: u.Password,
			IsAdmin:  u.IsAdmin,
		})
	}

	return usersLogic, nil
}

func (u *UserRepository) Create(ctx context.Context, user *logicModels.User) (*logicModels.User, error) {
	id := uuid.New()
	userDB := &repoModels.User{
		UUID:     id,
		Login:    user.Login,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}

	res := u.dbT.WithContext(ctx).Table("users").Create(userDB)
	if res.Error != nil {
		return nil, fmt.Errorf("insert: %w", res.Error)
	}

	resUser := logicModels.User{
		ID:       id,
		Login:    userDB.Login,
		Password: userDB.Password,
		IsAdmin:  userDB.IsAdmin,
	}

	return &resUser, nil
}

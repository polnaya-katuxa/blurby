package controllers

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	my_errors "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/errors"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/mocks"
	"github.com/gojuno/minimock/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func TestProfile_AuthByToken(t *testing.T) {
	mc := minimock.NewController(t)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ID:        uuid.New().String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte("secretkey"))

	claims1 := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ID:        "kejhrfbe",
	}

	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims1)
	ss1, _ := token1.SignedString([]byte("secretkey"))

	claims2 := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Microsecond)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ID:        uuid.New().String(),
	}

	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)
	ss2, _ := token2.SignedString([]byte("secretkey"))

	user := &models.User{
		ID:       uuid.New(),
		Login:    "uehfkjs",
		Password: "r5tg55g",
		IsAdmin:  false,
	}
	type fields struct {
		userRepo        interfaces.UserRepository
		tokenExpiration time.Duration
		secretKey       string
	}
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "successful auth",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByIDMock.Return(user, nil),
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx:   context.Background(),
				token: ss,
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "unsuccessful auth: parse",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByIDMock.Return(user, nil),
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx:   context.Background(),
				token: "ejrfnlekrfl",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful auth: id",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByIDMock.Return(user, nil),
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx:   context.Background(),
				token: ss1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful auth: get by id",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByIDMock.Return(user, errors.New("error")),
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx:   context.Background(),
				token: ss,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful auth: expired",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByIDMock.Return(user, nil),
				tokenExpiration: time.Microsecond,
				secretKey:       "secretkey",
			},
			args: args{
				ctx:   context.Background(),
				token: ss2,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				userRepository:  tt.fields.userRepo,
				tokenExpiration: tt.fields.tokenExpiration,
				secretKey:       tt.fields.secretKey,
			}
			got, err := p.AuthByToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthByToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfile_Register(t *testing.T) {
	mc := minimock.NewController(t)
	user := &models.User{
		ID:       uuid.New(),
		Login:    "uehfkjsyg",
		Password: "r3",
		IsAdmin:  false,
	}
	user2 := &models.User{
		ID:       uuid.New(),
		Login:    "g",
		Password: "3rgok34tglok",
		IsAdmin:  false,
	}
	type fields struct {
		userRepo        interfaces.UserRepository
		dailyBonus      int
		tokenExpiration time.Duration
		secretKey       string
	}
	type args struct {
		ctx      context.Context
		user     *models.User
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "successful register",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByLoginMock.Return(nil, my_errors.ErrUserNotFound).CreateMock.Return(user, nil),
				dailyBonus:      0,
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx: mycontext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "uehfkjs",
					Password: "ftioypue",
				}),
				user:     user,
				password: "ftioypue",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "unsuccessful register: get by login",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByLoginMock.Return(nil, errors.New("error")).CreateMock.Return(user, nil),
				dailyBonus:      0,
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx: mycontext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "uehfkjs",
					Password: "ftioypue",
				}),
				user:     user,
				password: "ftioypue",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful register: already exists",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByLoginMock.Return(user, nil).CreateMock.Return(user, nil),
				dailyBonus:      0,
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx: mycontext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "uehfkjs",
					Password: "ftioypue",
				}),
				user:     user,
				password: "ftioypue",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful register: login",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByLoginMock.Return(nil, my_errors.ErrUserNotFound).CreateMock.Return(user, nil),
				dailyBonus:      0,
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx: mycontext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "uehfkjs",
					Password: "ftioypue",
				}),
				user:     user2,
				password: "ftioypue",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful register: password",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByLoginMock.Return(nil, my_errors.ErrUserNotFound).CreateMock.Return(user, nil),
				dailyBonus:      0,
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx: mycontext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "uehfkjjs",
					Password: "ft",
				}),
				user:     user,
				password: "ft",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful register: create",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByLoginMock.Return(nil, my_errors.ErrUserNotFound).CreateMock.Return(user, errors.New("error")),
				dailyBonus:      0,
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx: mycontext.UserToContext(context.Background(), &models.User{
					ID:       uuid.New(),
					Login:    "uehfkjs",
					Password: "fttuytuty",
				}),
				user:     user,
				password: "fttuytuty",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				userRepository:  tt.fields.userRepo,
				tokenExpiration: tt.fields.tokenExpiration,
				secretKey:       tt.fields.secretKey,
			}
			_, err := p.Register(tt.args.ctx, tt.args.user, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProfile_Enter(t *testing.T) {
	mc := minimock.NewController(t)
	password := "password"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{
		ID:       uuid.New(),
		Login:    "uehfkjsyg",
		Password: string(hash),
		IsAdmin:  false,
	}
	type fields struct {
		userRepo        interfaces.UserRepository
		dailyBonus      int
		tokenExpiration time.Duration
		secretKey       string
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "successful enter",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByLoginMock.Return(user, nil),
				dailyBonus:      5,
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx:      context.Background(),
				login:    "uehfkjsyg",
				password: password,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "unsuccessful enter: wrong password",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByLoginMock.Return(user, nil),
				dailyBonus:      5,
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx:      context.Background(),
				login:    "uehfkjsyg",
				password: "r;glazyft",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "unsuccessful enter: get by login",
			fields: fields{
				userRepo:        mocks.NewUserRepositoryMock(mc).GetByLoginMock.Return(user, errors.New("error")),
				dailyBonus:      5,
				tokenExpiration: time.Hour,
				secretKey:       "secretkey",
			},
			args: args{
				ctx:      context.Background(),
				login:    "uehfkjsyg",
				password: password,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				userRepository:  tt.fields.userRepo,
				tokenExpiration: tt.fields.tokenExpiration,
				secretKey:       tt.fields.secretKey,
			}
			_, err := p.Login(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

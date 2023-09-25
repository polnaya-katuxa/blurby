package controllers

import (
	"context"
	"reflect"
	"testing"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/containers"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/repository/postgres"
	"github.com/google/uuid"
)

func TestClient_Get(t *testing.T) {
	dbContainer, db, err := containers.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = dbContainer.Terminate(context.Background()) }()

	cr := postgres.NewCR(db, db, db)

	user := &models.User{
		ID:       uuid.New(),
		Login:    "uehfkjsyg",
		Password: "r3",
		IsAdmin:  true,
	}

	userNA := &models.User{
		ID:       uuid.New(),
		Login:    "uehfkjsyg",
		Password: "r3",
		IsAdmin:  false,
	}

	id, err := uuid.Parse("77dcd288-79b2-4655-9584-cc9b5329665d")
	if err != nil {
		t.Errorf("uuid parse: %v", err)
		return
	}

	idNE, err := uuid.Parse("77dcd288-79b2-4655-9584-cc9b5329665a")
	if err != nil {
		t.Errorf("uuid parse: %v", err)
		return
	}

	bd, err := time.Parse(time.RFC3339, "2008-01-02T15:04:05Z")
	if err != nil {
		t.Errorf("bd parse: %v", err)
		return
	}

	rd, err := time.Parse(time.RFC3339, "2009-01-02T15:04:05Z")
	if err != nil {
		t.Errorf("bd parse: %v", err)
		return
	}

	client := &models.Client{
		ID:               id,
		Name:             "3",
		Surname:          "3",
		Patronymic:       "3",
		Gender:           "female",
		BirthDate:        bd,
		RegistrationDate: rd,
		Email:            "3@mail.ru",
		Data:             nil,
	}

	type fields struct {
		clientRepository interfaces.ClientRepository
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Client
		wantErr bool
	}{
		{
			name:   "successful get test",
			fields: fields{cr},
			args: args{
				ctx: mycontext.UserToContext(context.Background(), user),
				id:  id,
			},
			want:    client,
			wantErr: false,
		},
		{
			name:   "unsuccessful get test: empty context",
			fields: fields{cr},
			args: args{
				ctx: context.Background(),
				id:  id,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "unsuccessful get test: user is not an admin",
			fields: fields{cr},
			args: args{
				ctx: mycontext.UserToContext(context.Background(), userNA),
				id:  id,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "unsuccessful get test: non existing id",
			fields: fields{cr},
			args: args{
				ctx: mycontext.UserToContext(context.Background(), user),
				id:  idNE,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				clientRepository: tt.fields.clientRepository,
			}
			got, err := c.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

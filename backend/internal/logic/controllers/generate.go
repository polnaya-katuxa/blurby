package controllers

import (
	"context"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	mycontext "git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/util/rand"
)

type Generate struct {
	cr  interfaces.ClientRepository
	er  interfaces.EventTypeRepository
	ecr interfaces.EventCreationRepository
}

func NewGL(cr interfaces.ClientRepository, er interfaces.EventTypeRepository, ecr interfaces.EventCreationRepository) *Generate {
	return &Generate{
		cr:  cr,
		er:  er,
		ecr: ecr,
	}
}

var (
	ids     []uuid.UUID
	aliases []string
)

func genClient() (*models.Client, error) {
	id := uuid.New()
	ids = append(ids, id)

	bd, err := time.Parse(time.RFC3339, "2006-02-01T14:00:00Z")
	if err != nil {
		return nil, fmt.Errorf("parse bd: %w", err)
	}

	rd, err := time.Parse(time.RFC3339, "2008-02-01T14:00:00Z")
	if err != nil {
		return nil, fmt.Errorf("parse rd: %w", err)
	}

	return &models.Client{
		ID:               id,
		Name:             "a",
		Surname:          "a",
		Patronymic:       "a",
		Gender:           "male",
		BirthDate:        bd,
		RegistrationDate: rd,
		Email:            "a@a.ru",
		Data:             nil,
	}, nil
}

func genEventType() *models.EventType {
	alias := rand.String(15)
	aliases = append(aliases, alias)
	return &models.EventType{
		Name:  "aaa",
		Alias: alias,
	}
}

func genEvent() *models.Event {
	return &models.Event{
		ClientID:  ids[rand.Int()%1000],
		Alias:     aliases[rand.Int()%50],
		EventTime: time.Now(),
	}
}

func (g *Generate) Generate(ctx context.Context) error {
	for i := 0; i < 1000; i++ {
		c, err := genClient()
		if err != nil {
			return fmt.Errorf("generate: %w", err)
		}

		err = g.cr.Create(ctx, c)
		if err != nil {
			return fmt.Errorf("create client: %w", err)
		}

		fmt.Printf("\rGenerated %d/%d clients.", i+1, 1000)
	}

	fmt.Println()

	user := &models.User{
		ID:       uuid.New(),
		Login:    rand.String(10),
		Password: rand.String(10),
		IsAdmin:  false,
	}

	ctx = mycontext.UserToContext(ctx, user)

	for i := 0; i < 50; i++ {
		et := genEventType()

		err := g.er.Create(ctx, et)
		if err != nil {
			return fmt.Errorf("create event type: %w", err)
		}

		fmt.Printf("\rGenerated %d/%d event types.", i+1, 50)
	}

	fmt.Println()

	for i := 0; i < 1000000; i++ {
		e := genEvent()

		err := g.ecr.Create(ctx, e)
		if err != nil {
			return fmt.Errorf("create event: %w", err)
		}

		fmt.Printf("\rGenerated %d/%d events.", i+1, 1000000)
	}

	fmt.Println()

	return nil
}

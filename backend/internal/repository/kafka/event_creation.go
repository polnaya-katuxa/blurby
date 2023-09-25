package kafka

import (
	"context"
	"encoding/json"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
)

const timeFmt = "2006-01-02 15:04:05"

type EventCreationRepository struct {
	producer sarama.AsyncProducer
}

func NewECR(producer sarama.AsyncProducer) *EventCreationRepository {
	return &EventCreationRepository{producer: producer}
}

type event struct {
	ID       string `json:"id"`
	ClientID string `json:"client_id"`
	Alias    string `json:"alias"`
	Time     string `json:"time"`

	encoded []byte
	err     error
}

func (e *event) ensureEncoded() {
	if e.encoded == nil && e.err == nil {
		e.encoded, e.err = json.Marshal(e)
	}
}

func (e *event) Length() int {
	e.ensureEncoded()
	return len(e.encoded)
}

func (e *event) Encode() ([]byte, error) {
	e.ensureEncoded()
	return e.encoded, e.err
}

func (ec *EventCreationRepository) Create(_ context.Context, e *models.Event) error {
	v := &event{
		ID:       uuid.New().String(),
		ClientID: e.ClientID.String(),
		Alias:    e.Alias,
		Time:     e.EventTime.Format(timeFmt),
	}

	ec.producer.Input() <- &sarama.ProducerMessage{
		Topic: "blurby-events",
		Value: v,
	}

	return nil
}

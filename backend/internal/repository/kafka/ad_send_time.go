package kafka

import (
	"context"
	"encoding/json"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
)

type AdSendTimeRepository struct {
	producer sarama.AsyncProducer
}

func NewASTR(producer sarama.AsyncProducer) *AdSendTimeRepository {
	return &AdSendTimeRepository{producer: producer}
}

type adSendTime struct {
	ID   string `json:"id"`
	Time string `json:"time"`

	encoded []byte
	err     error
}

func (e *adSendTime) ensureEncoded() {
	if e.encoded == nil && e.err == nil {
		e.encoded, e.err = json.Marshal(e)
	}
}

func (e *adSendTime) Length() int {
	e.ensureEncoded()
	return len(e.encoded)
}

func (e *adSendTime) Encode() ([]byte, error) {
	e.ensureEncoded()
	return e.encoded, e.err
}

func (astr *AdSendTimeRepository) Create(_ context.Context, a *models.AdSendTime) error {
	v := &adSendTime{
		ID:   uuid.New().String(),
		Time: a.Time.Format(timeFmt),
	}

	astr.producer.Input() <- &sarama.ProducerMessage{
		Topic: "blurby-ad-send-times",
		Value: v,
	}

	return nil
}

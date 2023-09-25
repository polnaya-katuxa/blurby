package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/Shopify/sarama"
)

type AdQueueSendingRepository struct {
	producer sarama.SyncProducer
}

func NewASR(producer sarama.SyncProducer) *AdQueueSendingRepository {
	return &AdQueueSendingRepository{producer: producer}
}

type ad struct {
	ID         string          `json:"id"`
	Content    string          `json:"content"`
	Filters    []models.Filter `json:"filters"`
	UserID     string          `json:"user_id"`
	NextTime   string          `json:"next_time"`
	CreateTime string          `json:"create_time"`

	encoded []byte
	err     error
}

func (a *ad) ensureEncoded() {
	if a.encoded == nil && a.err == nil {
		a.encoded, a.err = json.Marshal(a)
	}
}

func (a *ad) Length() int {
	a.ensureEncoded()
	return len(a.encoded)
}

func (a *ad) Encode() ([]byte, error) {
	a.ensureEncoded()
	return a.encoded, a.err
}

func (as *AdQueueSendingRepository) SendToQueue(_ context.Context, ads []*models.Ad) error {
	msgs := make([]*sarama.ProducerMessage, 0, len(ads))

	for _, a := range ads {
		v := &ad{
			ID:         a.ID.String(),
			Content:    a.Content,
			Filters:    a.Filters,
			UserID:     a.UserID.String(),
			NextTime:   a.Schedule.NextTime.Format(time.RFC3339),
			CreateTime: a.CreateTime.Format(time.RFC3339),
		}

		msgs = append(msgs, &sarama.ProducerMessage{
			Topic: "blurby-ads",
			Value: v,
		})
	}

	err := as.producer.SendMessages(msgs)
	if err != nil {
		return fmt.Errorf("send ad to queue: %w", err)
	}

	return nil
}

package consumers

import (
	"encoding/json"
	"time"

	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/interfaces"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/context"
	"git.iu7.bmstu.ru/iu7-kostritsky/iu7-db-course-project-2023-karpovaekaterina-backend/internal/logic/models"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AdConsumer struct {
	senderLogic interfaces.SenderLogic
	logger      *zap.SugaredLogger
}

func NewAC(sl interfaces.SenderLogic, logger *zap.SugaredLogger) *AdConsumer {
	return &AdConsumer{senderLogic: sl, logger: logger}
}

func (c *AdConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *AdConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

type ad struct {
	ID         string          `json:"id"`
	Content    string          `json:"content"`
	Filters    []models.Filter `json:"filters"`
	UserID     string          `json:"user_id"`
	NextTime   string          `json:"next_time"`
	CreateTime string          `json:"create_time"`
}

func (c *AdConsumer) logError(session sarama.ConsumerGroupSession, message *sarama.ConsumerMessage, text string, err error) {
	c.logger.Errorw(
		text,
		"message", message.Value,
		"error", err,
		"key", message.Key,
		"topic", message.Topic,
		"partition", message.Partition,
		"offset", message.Offset,
	)

	session.MarkMessage(message, "")
}

func (c *AdConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			c.logger.Debugw(
				"claimed message",
				"message", message.Value,
				"key", message.Key,
				"topic", message.Topic,
				"partition", message.Partition,
				"offset", message.Offset,
			)

			var v ad
			err := json.Unmarshal(message.Value, &v)
			if err != nil {
				c.logError(session, message, "cannot unmarshal message", err)
				continue
			}

			t, err := time.Parse(time.RFC3339, v.NextTime)
			if err != nil {
				c.logError(session, message, "cannot parse time", err)
				continue
			}

			t = t.UTC()

			for time.Now().UTC().Before(t) {
			}

			id, err := uuid.Parse(v.ID)
			if err != nil {
				c.logError(session, message, "cannot parse uuid", err)
				continue
			}

			userID, err := uuid.Parse(v.UserID)
			if err != nil {
				c.logError(session, message, "cannot parse user uuid", err)
				continue
			}

			ad := &models.Ad{
				ID:      id,
				Content: v.Content,
				Filters: v.Filters,
				UserID:  userID,
			}

			err = c.senderLogic.Send(context.LoggerToContext(session.Context(), c.logger), ad)
			if err != nil {
				c.logError(session, message, "cannot send message", err)
				continue
			}

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

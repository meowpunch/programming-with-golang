package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"net/http"
)

type PaymentData struct {
	UserID        string  `json:"user_id"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"payment_method"`
}

func main() {
	app := fx.New(
		fx.Provide(
			NewConsumer,
			NewLogger,
		),
		fx.Invoke(RegisterConsumer),
	)
	app.Run()
}

func NewConsumer() sarama.Consumer {
	consumer, _ := sarama.NewConsumer([]string{"kafka:9093"}, nil)
	return consumer
}

func NewLogger() *logrus.Logger {
	return logrus.New()
}

func RegisterConsumer(lc fx.Lifecycle, consumer sarama.Consumer, logger *logrus.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			partitionConsumer, _ := consumer.ConsumePartition("payments", 0, sarama.OffsetNewest)
			go func() {
				for message := range partitionConsumer.Messages() {
					var payment PaymentData
					if err := json.Unmarshal(message.Value, &payment); err != nil {
						logger.Error("Failed to unmarshal message", err)
						continue
					}
					body, _ := json.Marshal(payment)
					_, err := http.Post("http://purchase-api:8081/purchase", "application/json", bytes.NewBuffer(body))
					if err != nil {
						logger.Error("Failed to send HTTP Post", err)
					}
				}
			}()
			return nil
		},
	})
}

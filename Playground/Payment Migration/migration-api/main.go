package main

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/go-chi/chi"
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
			NewRouter,
			NewProducer,
			NewLogger,
		),
		fx.Invoke(RegisterHandlers),
	)
	app.Run()
}

func NewRouter() *chi.Mux {
	return chi.NewRouter()
}

func NewProducer(logger *logrus.Logger) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"kafka:9093"}, config)
	if err != nil {
		logger.Error("Failed to create producer: ", err)
		return nil, err
	}
	return producer, nil
}

func NewLogger() *logrus.Logger {
	return logrus.New()
}

func RegisterHandlers(lc fx.Lifecycle, router *chi.Mux, producer sarama.SyncProducer, logger *logrus.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			router.Post("/migrate", func(w http.ResponseWriter, r *http.Request) {
				var payment PaymentData
				if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
					http.Error(w, "Invalid request body", http.StatusBadRequest)
					return
				}
				message, _ := json.Marshal(payment)
				if producer == nil {
					logger.Error("Producer is not initialized")
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				_, _, err := producer.SendMessage(&sarama.ProducerMessage{Topic: "payments", Value: sarama.ByteEncoder(message)})
				if err != nil {
					logger.Error("Failed to publish message", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
			})
			go http.ListenAndServe(":8080", router)
			return nil
		},
	})
}

package main

import (
	"context"
	"encoding/json"
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
			NewLogger,
		),
		fx.Invoke(RegisterHandlers),
	)
	app.Run()
}

func NewRouter() *chi.Mux {
	return chi.NewRouter()
}

func NewLogger() *logrus.Logger {
	return logrus.New()
}

func RegisterHandlers(lc fx.Lifecycle, router *chi.Mux, logger *logrus.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			router.Post("/purchase", func(w http.ResponseWriter, r *http.Request) {
				var payment PaymentData
				if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
					http.Error(w, "Invalid request body", http.StatusBadRequest)
					return
				}
				// Here, you would typically process the payment data and create a new payment in the global platform.
				logger.Infof("Received payment: %+v", payment)
			})
			go http.ListenAndServe(":8081", router)
			return nil
		},
	})
}

package transport

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
	"github.com/sail3/zemoga_test/internal/logger"
	"github.com/sail3/zemoga_test/pkg/health"
	"github.com/sail3/zemoga_test/pkg/portfolio"
)

const correlationIDHeader = "X-Correlation-ID"

// NewHTTPRouter initializes the router using the services as dependencies to build the handlers.
func NewHTTPRouter(healthSvc health.Service, portfolioSvc portfolio.Handler, log logger.Logger) http.Handler {
	hh := health.Handler{
		Service: healthSvc,
	}

	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(correlationIDMiddleware)

	// Health.
	r.Get("/health", hh.Health)

	r.Get("/profile/{id}", portfolioSvc.GetProfileHandler)
	r.Get("/profile", portfolioSvc.ListProfileHandler)
	r.Get("/profile/{id}/tweet", portfolioSvc.ListTweetsHandler)
	r.Patch("/profile/{id}", portfolioSvc.UpdateProfileHandler)
	r.Handle("/swagger/*", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger"))))

	return r
}

// correlationIDMiddleware injects the correlation ID in the context if it's missing or
// uses the received one if present on the request
func correlationIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := r.Header.Get(correlationIDHeader)
		if id == "" {
			// generate new version 4 uuid
			id = uuid.New().String()
		}
		// set the id to the request context
		ctx = context.WithValue(ctx, logger.CorrelationIDKey, id)
		r = r.WithContext(ctx)

		// set the response header
		w.Header().Set(correlationIDHeader, id)
		next.ServeHTTP(w, r)
	})
}

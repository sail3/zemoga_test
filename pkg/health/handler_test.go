package health

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errNotFound = errors.New("not found")

type serviceHealthCheckSuccess struct{}

func (s serviceHealthCheckSuccess) HealthCheck(_ context.Context) (service bool, err error) {
	return true, nil
}

type serviceHealthCheckFailure struct{}

func (s serviceHealthCheckFailure) HealthCheck(_ context.Context) (service bool, err error) {
	return false, errNotFound
}

func TestHandler_Health(t *testing.T) {
	tests := []struct {
		name       string
		service    Service
		statusCode int
	}{
		{
			name:       "Success",
			service:    serviceHealthCheckSuccess{},
			statusCode: http.StatusOK,
		}, {
			name:       "Failure",
			service:    serviceHealthCheckFailure{},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := Handler{Service: test.service}

			mux := chi.NewRouter()

			mux.Get("/health", h.Health)

			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "/health", nil)
			if err != nil {
				require.NoError(t, err)
			}

			mux.ServeHTTP(w, r)

			statusCode := w.Result().StatusCode
			assert.Equal(t, test.statusCode, statusCode)
		})
	}

}

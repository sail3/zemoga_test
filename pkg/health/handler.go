package health

import (
	"net/http"

	"github.com/sail3/zemoga_test/internal/response"
)

// Handler contains the methods to handle the endpoints of health.
type Handler struct {
	Service Service
}

// Health is the handler for the health endpoint.
func (c *Handler) Health(w http.ResponseWriter, r *http.Request) {
	service, err := c.Service.HealthCheck(r.Context())
	if err != nil {
		_ = response.RespondWithError(w, response.StandardInternalServerError)
		return
	}

	hr := HealthResponse{
		Services: []HealthService{
			{
				Name:  "service",
				Alive: service,
			},
		},
	}

	_ = response.RespondWithData(w, http.StatusOK, hr)
}

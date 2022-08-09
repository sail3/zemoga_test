package analytics

import (
	"context"
	"net/http"

	"github.com/sail3/zemoga_test/internal/logger"
	"github.com/sail3/zemoga_test/internal/response"
)

type Handler struct {
	service Service
	log     logger.Logger
}

func NewHandler(svc Service, log logger.Logger) Handler {
	return Handler{
		service: svc,
		log:     log,
	}
}

func (h *Handler) AnalyticsResumeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.log.WithCorrelation(ctx)
	res, err := h.service.GetAnalyticsResume(ctx)
	if err != nil {
		log.Error(err)
		_ = response.RespondWithError(w, response.StandardInternalServerError)
		return
	}

	_ = response.RespondWithData(w, http.StatusOK, res)
}

func (h *Handler) InitWatcher(ch chan map[string]string, l logger.Logger) {
	go func() {
		for url := range ch {
			h.service.RegisterCall(context.Background(), RegisterCallRequest{
				URL:    url["url"],
				Method: url["method"],
			})
		}
	}()
}

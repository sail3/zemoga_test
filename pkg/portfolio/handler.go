package portfolio

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

func (h Handler) GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.log.WithCorrelation(ctx)
	profileID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(profileID)
	if err != nil {
		log.Error(err)
		_ = response.RespondWithError(w, response.StandardBadBodyRequest)
		return
	}

	profile, err := h.service.GetProfile(ctx, GetProfileRequest{
		ID: id,
	})
	if err != nil {
		log.Error(err)
		_ = response.RespondWithError(w, response.StandardInternalServerError)
		return
	}

	_ = response.RespondWithHTML(w, "index", profile)
}

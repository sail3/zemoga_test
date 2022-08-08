package portfolio

import (
	"encoding/json"
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

	profile, err := h.service.GetProfile(ctx, GetProfileRequest{
		ID: profileID,
	})
	if err != nil {
		log.Error(err)
		_ = response.RespondWithError(w, response.StandardInternalServerError)
		return
	}

	_ = response.RespondWithHTML(w, "index", profile)
}

func (h Handler) ListProfileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.log.WithCorrelation(ctx)

	res, err := h.service.ListProfile(ctx)
	if err != nil {
		log.Error(err)
		_ = response.RespondWithError(w, response.StandardInternalServerError)
		return
	}

	_ = response.RespondWithData(w, http.StatusOK, res)
}

func (h Handler) ListTweetsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.log.WithCorrelation(ctx)
	u := chi.URLParam(r, "id")

	q := r.URL.Query().Get("quantity")
	quantity, err := strconv.Atoi(q)
	if err != nil {
		log.Error(err)
		quantity = 5
	}

	res, err := h.service.GetTweetList(ctx, u, quantity)
	if err != nil {
		log.Error(err)
		_ = response.RespondWithError(w, response.StandardInternalServerError)
		return
	}

	_ = response.RespondWithData(w, http.StatusOK, res)
}

func (h Handler) UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.log.WithCorrelation(ctx)
	id := chi.URLParam(r, "id")

	var pr ProfileRequest
	err := json.NewDecoder(r.Body).Decode(&pr)
	if err != nil {
		log.Error(err)
		_ = response.RespondWithError(w, response.StandardBadBodyRequest)
		return
	}
	defer r.Body.Close()

	res, err := h.service.UpdateProfile(ctx, id, pr)
	if err != nil {
		log.Error(err)
		_ = response.RespondWithError(w, response.StandardInternalServerError)
		return
	}

	_ = response.RespondWithData(w, http.StatusOK, res)

}

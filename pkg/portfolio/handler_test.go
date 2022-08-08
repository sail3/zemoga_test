package portfolio

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/sail3/zemoga_test/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type listProfileServiceMock struct {
	baseServiceMock
	res []ProfileResponse
	err error
}

func (m listProfileServiceMock) ListProfile(ctx context.Context) ([]ProfileResponse, error) {
	return m.res, m.err
}
func TestHandler_ListProfile(t *testing.T) {
	genericError := errors.New("generic error")
	tests := []struct {
		name           string
		statusCode     int
		req            []byte
		serviceListRes []ProfileResponse
		serviceListErr error
	}{
		{
			name:           "success",
			statusCode:     http.StatusOK,
			req:            []byte(`*`),
			serviceListRes: []ProfileResponse{},
			serviceListErr: nil,
		},
		{
			name:           "fail",
			statusCode:     http.StatusInternalServerError,
			req:            []byte(`*`),
			serviceListRes: []ProfileResponse{},
			serviceListErr: genericError,
		},
	}
	l := logger.Mock()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sm := listProfileServiceMock{
				res: test.serviceListRes,
				err: test.serviceListErr,
			}
			w := httptest.NewRecorder()
			b := bytes.NewBuffer(test.req)
			r, err := http.NewRequest(http.MethodGet, "/profile", b)
			if err != nil {
				require.NoError(t, err)
			}

			h := NewHandler(sm, l)
			mux := chi.NewMux()
			mux.Get("/profile", h.ListProfileHandler)
			mux.ServeHTTP(w, r)

			statusCode := w.Result().StatusCode
			assert.Equal(t, test.statusCode, statusCode)
		})
	}

}

type baseServiceMock struct {
}

func (_ baseServiceMock) GetProfile(ctx context.Context, p GetProfileRequest) (GetProfileResponse, error) {
	panic("implement me")
}
func (_ baseServiceMock) ListProfile(ctx context.Context) ([]ProfileResponse, error) {
	panic("implement me")
}
func (_ baseServiceMock) GetTweetList(ctx context.Context, userID string, quantity int) ([]TweetResponse, error) {
	panic("implement me")
}
func (_ baseServiceMock) UpdateProfile(ctx context.Context, userID string, pr ProfileRequest) (ProfileResponse, error) {
	panic("implement me")
}

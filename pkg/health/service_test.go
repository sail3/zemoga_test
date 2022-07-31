package health

import (
	"context"
	"testing"

	"github.com/sail3/zemoga_test/internal/logger"
)

func TestHealthCheck(t *testing.T) {
	service := NewService(
		logger.Mock(),
	)

	s, err := service.HealthCheck(context.Background())
	if s != true || err != nil {
		t.Errorf("Unexpected values from method: service %t, error %s", s, err)
	}
}

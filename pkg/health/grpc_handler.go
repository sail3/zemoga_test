package health

import (
	"context"

	"github.com/sail3/zemoga_test/internal/logger"
)

type server struct {
	UnimplementedHealthServiceServer
	log logger.Logger
}

// GetTopPlayersInfo returns a list of top player information.
func (s server) GetHealthStatus(_ context.Context, _ *GetHealthStatusRequest) (*GetHealthStatusResponse, error) {
	return &GetHealthStatusResponse{
		Name:  "zemoga_test",
		Alive: true,
	}, nil
}

// NewGRPCServer returns the gRPC server to the zemoga_test service.
func NewGRPCServer(log logger.Logger) HealthServiceServer {
	return &server{
		log: log,
	}
}

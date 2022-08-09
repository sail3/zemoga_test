package analytics

import (
	"context"

	"github.com/sail3/zemoga_test/internal/logger"
)

type Service interface {
	GetAnalyticsResume(ctx context.Context) ([]ResumeResponse, error)
	RegisterCall(ctx context.Context, req RegisterCallRequest)
}

func NewService(repo Repository, log logger.Logger) Service {
	return &service{
		repository: repo,
		log:        log,
	}
}

type service struct {
	repository Repository
	log        logger.Logger
}

func (s service) GetAnalyticsResume(ctx context.Context) ([]ResumeResponse, error) {
	log := s.log.WithCorrelation(ctx)
	res, err := s.repository.GetCallResume(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	list := make([]ResumeResponse, 0)
	for _, v := range res {
		vp := ResumeResponse{
			Url:      v.URL,
			Method:   v.Method,
			Quantity: v.Quantity,
		}
		list = append(list, vp)
	}

	return list, nil
}

func (s service) RegisterCall(ctx context.Context, req RegisterCallRequest) {
	log := s.log.WithCorrelation(ctx)
	err := s.repository.IncrementRequestPath(ctx, Call{
		Method:   req.Method,
		URL:      req.URL,
		Quantity: 1,
	})
	if err != nil {
		log.Error(err)
	}
}

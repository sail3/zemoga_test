package portfolio

import (
	"context"

	"github.com/sail3/zemoga_test/internal/logger"
	"github.com/sail3/zemoga_test/pkg/twitter"
)

type Service interface {
	GetProfile(ctx context.Context, p GetProfileRequest) (GetProfileResponse, error)
}

func NewService(repo Repository, tw twitter.Service, log logger.Logger) Service {
	return &service{
		repository: repo,
		twService:  tw,
		log:        log,
	}
}

type service struct {
	repository Repository
	twService  twitter.Service
	log        logger.Logger
}

func (s service) GetProfile(ctx context.Context, u GetProfileRequest) (GetProfileResponse, error) {
	log := s.log.WithCorrelation(ctx)
	res, err := s.repository.GetProfileByID(ctx, u.ID)
	if err != nil {
		log.Error(err)
		return GetProfileResponse{}, err
	}
	twTimeline, err := s.twService.GetUserTimeLine(res.TwitterID, 5)
	if err != nil {
		log.Info("could not be obtained twitts", err)
	}
	return GetProfileResponse{
		ID:              res.ID,
		Title:           res.Title,
		Name:            res.Name,
		Description:     res.Description,
		Image:           res.Image,
		TwitterUser:     res.TwitterUser,
		TwitterID:       res.TwitterID,
		TwitterTimeLine: twTimeline,
	}, nil
}

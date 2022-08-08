package portfolio

import (
	"context"

	"github.com/sail3/zemoga_test/internal/logger"
	"github.com/sail3/zemoga_test/pkg/twitter"
)

type Service interface {
	GetProfile(ctx context.Context, p GetProfileRequest) (GetProfileResponse, error)
	ListProfile(ctx context.Context) ([]ProfileResponse, error)
	GetTweetList(ctx context.Context, userID string, quantity int) ([]TweetResponse, error)
	UpdateProfile(ctx context.Context, userID string, pr ProfileRequest) (ProfileResponse, error)
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

func (s service) ListProfile(ctx context.Context) ([]ProfileResponse, error) {
	log := s.log.WithCorrelation(ctx)
	res, err := s.repository.GetAllProfile(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	list := make([]ProfileResponse, 0)
	for _, v := range res {
		vp := ProfileResponse{
			ID:          v.ID,
			Title:       v.Title,
			Name:        v.Name,
			Description: v.Description,
			Image:       v.Image,
			TwitterUser: v.TwitterUser,
			TwitterID:   v.TwitterID,
		}
		twTimeline, err := s.twService.GetUserTimeLine(v.TwitterID, 5)
		if err != nil {
			log.Info("could not be obtained twitts", err)
		}
		vp.TwitterTimeLine = twTimeline
		list = append(list, vp)
	}

	return list, nil
}

func (s service) GetTweetList(ctx context.Context, userID string, quantity int) ([]TweetResponse, error) {
	log := s.log.WithCorrelation(ctx)
	user, err := s.repository.GetProfileByID(ctx, userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	tl, err := s.twService.GetUserTimeLine(user.TwitterID, quantity)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	res := make([]TweetResponse, 0)
	for _, v := range tl {
		res = append(res, TweetResponse(v))
	}
	return res, nil
}

func (s service) UpdateProfile(ctx context.Context, userID string, pr ProfileRequest) (ProfileResponse, error) {
	log := s.log.WithCorrelation(ctx)
	op, err := s.repository.GetProfileByID(ctx, userID)
	if err != nil {
		log.Error(err)
		return ProfileResponse{}, err
	}
	if pr.Description != "" {
		op.Description = pr.Description
	}
	if pr.Image != "" {
		op.Image = pr.Image
	}
	if pr.Name != "" {
		op.Name = pr.Name
	}
	if pr.Title != "" {
		op.Title = pr.Title
	}
	if pr.TwitterUser != "" {
		op.TwitterUser = pr.TwitterUser
		tu, err := s.twService.GetUserProfile(op.TwitterUser)
		if err != nil {
			log.Error(err)
			return ProfileResponse{}, err
		}
		op.TwitterID = tu.ID
	}
	err = s.repository.UpdateProfile(ctx, userID, op)
	if err != nil {
		log.Error(err)
		return ProfileResponse{}, err
	}
	return ProfileResponse{
		ID:          op.ID,
		Title:       op.Title,
		Name:        op.Name,
		Description: op.Description,
		Image:       op.Image,
		TwitterUser: op.TwitterUser,
		TwitterID:   op.TwitterID,
	}, nil
}

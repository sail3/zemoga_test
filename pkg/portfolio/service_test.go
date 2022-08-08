package portfolio

import (
	"context"
	"errors"
	"testing"

	"github.com/sail3/zemoga_test/internal/logger"
	"github.com/sail3/zemoga_test/pkg/twitter"
	"github.com/stretchr/testify/assert"
)

type getProfileMock struct {
	baseRepoMock
	resp Profile
	err  error
}

func (m getProfileMock) GetProfileByID(ctx context.Context, ID string) (Profile, error) {
	return m.resp, m.err
}

type getUserTimeLineMock struct {
	baseTwitterMock
	resp []twitter.Tweet
	err  error
}

func (m getUserTimeLineMock) GetUserTimeLine(userID int64, size int) ([]twitter.Tweet, error) {
	return m.resp, m.err
}

func TestService_GetProfile(t *testing.T) {
	req := GetProfileRequest{
		ID: "62f17a504756615790132def",
	}
	twTimeline := make([]twitter.Tweet, 0)
	resp := GetProfileResponse{}
	resp.TwitterTimeLine = twTimeline
	errNotFound := errors.New("user Not found")
	tests := []struct {
		name           string
		req            GetProfileRequest
		resp           GetProfileResponse
		err            error
		getProfileResp Profile
		getProfileErr  error
		twitterResp    []twitter.Tweet
		twitterErr     error
	}{
		{
			name:           "Success",
			req:            req,
			resp:           resp,
			err:            nil,
			getProfileResp: Profile{},
			getProfileErr:  nil,
			twitterResp:    twTimeline,
			twitterErr:     nil,
		},
		{
			name:           "Repo fail",
			req:            req,
			resp:           GetProfileResponse{},
			err:            errNotFound,
			getProfileResp: Profile{},
			getProfileErr:  errNotFound,
			twitterResp:    twTimeline,
			twitterErr:     nil,
		},
		{
			name:           "Time line fail",
			req:            req,
			resp:           resp,
			err:            nil,
			getProfileResp: Profile{},
			getProfileErr:  nil,
			twitterResp:    twTimeline,
			twitterErr:     errNotFound,
		},
	}
	l := logger.Mock()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rm := getProfileMock{
				resp: test.getProfileResp,
				err:  test.getProfileErr,
			}
			tm := getUserTimeLineMock{
				resp: test.twitterResp,
				err:  test.twitterErr,
			}
			s := NewService(rm, tm, l)
			ctx := context.Background()
			resp, err := s.GetProfile(ctx, test.req)
			assert.Equal(t, test.resp, resp)
			assert.Equal(t, test.err, err)
		})
	}
}

type listProfileMock struct {
	baseRepoMock
	resp []Profile
	err  error
}

func (m listProfileMock) GetAllProfile(ctx context.Context) ([]Profile, error) {
	return m.resp, m.err
}

func TestService_ListProfile(t *testing.T) {
	genericError := errors.New("generic error")
	twTimeline := make([]twitter.Tweet, 0)
	allProfilesResp := []Profile{
		{
			ID:          "12",
			Title:       "some title",
			Name:        "some name",
			Description: "some description",
			Image:       "some image",
			TwitterUser: "some twitter user",
			TwitterID:   123123,
		},
	}
	resp := []ProfileResponse{
		{
			ID:              "12",
			Title:           "some title",
			Name:            "some name",
			Description:     "some description",
			Image:           "some image",
			TwitterUser:     "some twitter user",
			TwitterID:       123123,
			TwitterTimeLine: twTimeline,
		},
	}
	tests := []struct {
		name              string
		resp              []ProfileResponse
		err               error
		getAllProfileResp []Profile
		getAllProfileErr  error
		twitterResp       []twitter.Tweet
		twitterErr        error
	}{
		{
			name:              "Success",
			resp:              resp,
			err:               nil,
			getAllProfileResp: allProfilesResp,
			getAllProfileErr:  nil,
			twitterResp:       twTimeline,
			twitterErr:        nil,
		},
		{
			name:              "Repo all fail",
			resp:              nil,
			err:               genericError,
			getAllProfileResp: []Profile{},
			getAllProfileErr:  genericError,
			twitterResp:       twTimeline,
			twitterErr:        nil,
		},
		{
			name:              "Time Line Fail",
			resp:              resp,
			err:               nil,
			getAllProfileResp: allProfilesResp,
			getAllProfileErr:  nil,
			twitterResp:       twTimeline,
			twitterErr:        genericError,
		},
	}
	l := logger.Mock()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rm := listProfileMock{
				resp: test.getAllProfileResp,
				err:  test.getAllProfileErr,
			}
			tm := getUserTimeLineMock{
				resp: test.twitterResp,
				err:  test.twitterErr,
			}
			s := NewService(rm, tm, l)
			ctx := context.Background()
			resp, err := s.ListProfile(ctx)
			assert.Equal(t, test.resp, resp)
			assert.Equal(t, test.err, err)
		})
	}
}

type getTweetListMock struct {
	baseRepoMock
	resp Profile
	err  error
}

func (m getTweetListMock) GetProfileByID(ctx context.Context, ID string) (Profile, error) {
	return m.resp, m.err
}

func TestService_GetTweetList(t *testing.T) {
	genericError := errors.New("generic error")
	twTimeline := []twitter.Tweet{
		{
			ID:    123,
			Image: "some image",
			Text:  "some text",
		},
	}
	profileResp := Profile{
		ID:          "12",
		Title:       "some title",
		Name:        "some name",
		Description: "some description",
		Image:       "some image",
		TwitterUser: "some twitter user",
		TwitterID:   123123,
	}
	resp := []TweetResponse{
		TweetResponse(twTimeline[0]),
	}
	tests := []struct {
		name             string
		resp             []TweetResponse
		err              error
		getProfileIDResp Profile
		getProfileIDErr  error
		twitterResp      []twitter.Tweet
		twitterErr       error
	}{
		{
			name:             "Success",
			resp:             resp,
			err:              nil,
			getProfileIDResp: profileResp,
			getProfileIDErr:  nil,
			twitterResp:      twTimeline,
			twitterErr:       nil,
		},
		{
			name:             "Repo get profile fail",
			resp:             nil,
			err:              genericError,
			getProfileIDResp: Profile{},
			getProfileIDErr:  genericError,
			twitterResp:      twTimeline,
			twitterErr:       nil,
		},
		{
			name:             "Time Line Fail",
			resp:             nil,
			err:              genericError,
			getProfileIDResp: profileResp,
			getProfileIDErr:  nil,
			twitterResp:      twTimeline,
			twitterErr:       genericError,
		},
	}
	l := logger.Mock()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rm := getTweetListMock{
				resp: test.getProfileIDResp,
				err:  test.getProfileIDErr,
			}
			tm := getUserTimeLineMock{
				resp: test.twitterResp,
				err:  test.twitterErr,
			}
			s := NewService(rm, tm, l)
			ctx := context.Background()
			resp, err := s.GetTweetList(ctx, "12", 12)
			assert.Equal(t, test.resp, resp)
			assert.Equal(t, test.err, err)
		})
	}
}

type updateProfileRepoMock struct {
	baseRepoMock
	profileRes Profile
	profileErr error
	updateErr  error
}

func (m updateProfileRepoMock) GetProfileByID(ctx context.Context, ID string) (Profile, error) {
	return m.profileRes, m.profileErr
}

func (m updateProfileRepoMock) UpdateProfile(ctx context.Context, ID string, p Profile) error {
	return m.updateErr
}

type updateProfileTwitterMock struct {
	baseTwitterMock
	res twitter.User
	err error
}

func (m updateProfileTwitterMock) GetUserProfile(userName string) (twitter.User, error) {
	return m.res, m.err
}

func TestService_UpdateProfile(t *testing.T) {
	genericError := errors.New("generic error")
	profileResp := Profile{
		ID:          "12",
		Title:       "some title",
		Name:        "some name",
		Description: "some description",
		Image:       "some image",
		TwitterUser: "some twitter user",
		TwitterID:   123123,
	}
	twUser := twitter.User{
		ID:       123,
		Username: "some user",
		Image:    "some image",
	}

	response := ProfileResponse{
		ID:          "12",
		Title:       "some title changed",
		Name:        "some name changed",
		Description: "some description changed",
		Image:       "some image changed",
		TwitterUser: "some user changed",
		TwitterID:   123,
	}

	request := ProfileRequest{
		Title:       "some title changed",
		Name:        "some name changed",
		Description: "some description changed",
		Image:       "some image changed",
		TwitterUser: "some user changed",
	}

	tests := []struct {
		name           string
		req            ProfileRequest
		res            ProfileResponse
		err            error
		repoProfileRes Profile
		repoProfileErr error
		repoUpdateErr  error
		twRes          twitter.User
		twErr          error
	}{
		{
			name:           "success",
			req:            request,
			res:            response,
			err:            nil,
			repoProfileRes: profileResp,
			repoProfileErr: nil,
			repoUpdateErr:  nil,
			twRes:          twUser,
			twErr:          nil,
		},
		{
			name:           "Repo get profile fail",
			req:            request,
			res:            ProfileResponse{},
			err:            genericError,
			repoProfileRes: Profile{},
			repoProfileErr: genericError,
			repoUpdateErr:  nil,
			twRes:          twUser,
			twErr:          nil,
		},
		{
			name:           "Repo update profile fail",
			req:            request,
			res:            ProfileResponse{},
			err:            genericError,
			repoProfileRes: profileResp,
			repoProfileErr: nil,
			repoUpdateErr:  genericError,
			twRes:          twUser,
			twErr:          nil,
		},
		{
			name:           "Twitter get profile fail",
			req:            request,
			res:            ProfileResponse{},
			err:            genericError,
			repoProfileRes: profileResp,
			repoProfileErr: nil,
			repoUpdateErr:  nil,
			twRes:          twitter.User{},
			twErr:          genericError,
		},
	}

	l := logger.Mock()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rm := updateProfileRepoMock{
				profileRes: test.repoProfileRes,
				profileErr: test.repoProfileErr,
				updateErr:  test.repoUpdateErr,
			}
			tm := updateProfileTwitterMock{
				res: test.twRes,
				err: test.twErr,
			}
			s := NewService(rm, tm, l)
			ctx := context.Background()
			res, err := s.UpdateProfile(ctx, "12", test.req)
			assert.Equal(t, test.res, res)
			assert.Equal(t, test.err, err)
		})
	}

}

type baseRepoMock struct{}

func (_ baseRepoMock) GetProfileByID(ctx context.Context, ID string) (Profile, error) {
	panic("implementme")
}

func (_ baseRepoMock) GetAllProfile(ctx context.Context) ([]Profile, error) {
	panic("implementme")
}
func (_ baseRepoMock) UpdateProfile(ctx context.Context, ID string, p Profile) error {
	panic("implementme")
}

type baseTwitterMock struct{}

func (_ baseTwitterMock) GetUserTimeLine(userID int64, size int) ([]twitter.Tweet, error) {
	panic("implementme")
}

func (_ baseTwitterMock) GetUserProfile(userName string) (twitter.User, error) {
	panic("implementme")
}

package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/sail3/zemoga_test/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Service interface {
	GetUserTimeLine(userID int64, size int) ([]Tweet, error)
	GetUserProfile(userName string) (User, error)
}
type service struct {
	clientID     string
	clientSecret string
	tokenURL     string
}

func NewService(conf config.Config) Service {
	return service{
		clientID:     conf.TwitterID,
		clientSecret: conf.TwitterSecret,
		tokenURL:     conf.TwitterTokenURL,
	}
}

func (s service) getClient() *twitter.Client {
	config := &clientcredentials.Config{
		ClientID:     s.clientID,
		ClientSecret: s.clientSecret,
		TokenURL:     s.tokenURL,
	}
	hc := config.Client(oauth2.NoContext)
	c := twitter.NewClient(hc)
	return c
}

func (s service) GetUserProfile(userName string) (User, error) {
	client := s.getClient()
	user, _, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: userName,
	})
	if err != nil {
		return User{}, err
	}

	return User{
		ID:       user.ID,
		Username: user.Name,
		Image:    user.ProfileImageURL,
	}, nil
}

func (s service) GetUserTimeLine(userID int64, size int) ([]Tweet, error) {
	client := s.getClient()
	twitts, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	res := make([]Tweet, 0)
	for _, v := range twitts[0:size] {
		var t Tweet
		t.Text = v.Text
		t.ID = v.ID
		t.Image = v.User.ProfileImageURL
		res = append(res, t)
	}

	return res, nil
}

package portfolio

import "github.com/sail3/zemoga_test/pkg/twitter"

type GetProfileRequest struct {
	ID int `json:"id,omitempty"`
}

type GetProfileResponse struct {
	ID              string          `json:"id,omitempty"`
	Title           string          `json:"title,omitempty"`
	Name            string          `json:"name,omitempty"`
	Description     string          `json:"description,omitempty"`
	Image           string          `json:"image,omitempty"`
	TwitterUser     string          `json:"twitter_user,omitempty"`
	TwitterID       int64           `json:"twitter_id,omitempty"`
	TwitterTimeLine []twitter.Tweet `json:"twitter_time_line,omitempty"`
}

type ProfileResponse struct {
	ID              string          `json:"id,omitempty"`
	Title           string          `json:"title,omitempty"`
	Name            string          `json:"name,omitempty"`
	Description     string          `json:"description,omitempty"`
	Image           string          `json:"image,omitempty"`
	TwitterUser     string          `json:"twitter_user,omitempty"`
	TwitterID       int64           `json:"twitter_id,omitempty"`
	TwitterTimeLine []twitter.Tweet `json:"twitter_time_line,omitempty"`
}

type TweetResponse struct {
	ID    int64  `json:"id,omitempty"`
	Image string `json:"image,omitempty"`
	Text  string `json:"text,omitempty"`
}

type ProfileRequest struct {
	Title       string `json:"title,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
	TwitterUser string `json:"twitter_user,omitempty"`
}

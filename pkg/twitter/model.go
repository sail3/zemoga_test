package twitter

type Tweet struct {
	ID    int64  `json:"id,omitempty"`
	Image string `json:"image,omitempty"`
	Text  string `json:"text,omitempty"`
}

type User struct {
	ID       int64  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Image    string `json:"image,omitempty"`
}

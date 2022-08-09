package analytics

type ResumeResponse struct {
	Url      string `json:"url,omitempty"`
	Method   string `json:"path,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

type RegisterCallRequest struct {
	URL    string
	Method string
}

package analytics

type Call struct {
	Method   string `bson:"method,omitempty"`
	URL      string `bson:"url,omitempty"`
	Quantity int    `bson:"quantity,omitempty"`
}

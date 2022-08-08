package portfolio

import "go.mongodb.org/mongo-driver/bson/primitive"

type Profile struct {
	ID          string `bson:"_id,omitempty"`
	Title       string `bson:"title,omitempty"`
	Name        string `bson:"name,omitempty"`
	Description string `bson:"description,omitempty"`
	Image       string `bson:"image,omitempty"`
	TwitterUser string `bson:"twitter_user,omitempty"`
	TwitterID   int64  `bson:"twitter_id,omitempty"`
}

type Filter struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

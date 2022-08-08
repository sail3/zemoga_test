package portfolio

import (
	"context"
	"time"

	"github.com/sail3/zemoga_test/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetProfileByID(ctx context.Context, ID string) (Profile, error)
	GetAllProfile(ctx context.Context) ([]Profile, error)
	UpdateProfile(ctx context.Context, ID string, p Profile) error
}

func NewRepository(cl *mongo.Client, DBName string, log logger.Logger) Repository {
	return &repository{
		client:   cl,
		database: cl.Database(DBName),
		log:      log,
	}
}

type repository struct {
	client   *mongo.Client
	database *mongo.Database
	log      logger.Logger
}

func (r *repository) GetProfileByID(ctx context.Context, ID string) (Profile, error) {
	coll := r.database.Collection("profile")
	log := r.log.WithCorrelation(ctx)
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Error(err)
		return Profile{}, err
	}
	filter := bson.M{"_id": objectID}
	var p Profile
	err = coll.FindOne(ctx, filter).Decode(&p)
	if err != nil {
		log.Error(err)
		return Profile{}, nil
	}

	return p, nil
}

func (r *repository) GetAllProfile(ctx context.Context) ([]Profile, error) {
	coll := r.database.Collection("profile")
	log := r.log.WithCorrelation(ctx)
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		log.Error(err)
		return nil, nil
	}
	var res = make([]Profile, 0)
	for cursor.Next(ctx) {
		var p Profile
		err := cursor.Decode(&p)
		if err != nil {
			return nil, nil
		}
		res = append(res, p)
	}
	return res, nil
}

func (r *repository) UpdateProfile(ctx context.Context, ID string, p Profile) error {
	coll := r.database.Collection("profile")
	log := r.log.WithCorrelation(ctx)
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Error(err)
		return err
	}
	filter := bson.M{"_id": objectID}
	u := bson.D{primitive.E{Key: "$set", Value: bson.M{
		"name":         p.Name,
		"title":        p.Title,
		"image":        p.Image,
		"twitter_user": p.TwitterUser,
		"twitter_id":   p.TwitterID,
		"description":  p.Description,
	}}}
	resp := coll.FindOneAndUpdate(ctx, filter, u)
	if resp.Err() != nil {
		log.Error(resp.Err())
		return resp.Err()
	}
	return nil
}

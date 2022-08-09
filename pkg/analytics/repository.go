package analytics

import (
	"context"
	"time"

	"github.com/sail3/zemoga_test/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	IncrementRequestPath(ctx context.Context, call Call) error
	GetCallResume(ctx context.Context) ([]Call, error)
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

func (r repository) IncrementRequestPath(ctx context.Context, call Call) error {
	coll := r.database.Collection("analytics")
	log := r.log.WithCorrelation(ctx)
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	filter := bson.D{{"url", call.URL}, {"method", call.Method}}
	values := bson.D{{"$inc", bson.D{{"quantity", call.Quantity}}}}
	opt := options.Update().SetUpsert(true)
	_, err := coll.UpdateOne(ctx, filter, values, opt)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r repository) GetCallResume(ctx context.Context) ([]Call, error) {
	coll := r.database.Collection("analytics")
	log := r.log.WithCorrelation(ctx)
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		log.Error(err)
		return nil, nil
	}
	var res = make([]Call, 0)
	for cursor.Next(ctx) {
		var p Call
		err := cursor.Decode(&p)
		if err != nil {
			return nil, nil
		}
		res = append(res, p)
	}
	return res, nil
}

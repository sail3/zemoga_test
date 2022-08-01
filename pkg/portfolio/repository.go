package portfolio

import (
	"context"
	"database/sql"

	"github.com/sail3/zemoga_test/internal/logger"
)

type Repository interface {
	GetProfileByID(ctx context.Context, ID int) (Profile, error)
}

func NewRepository(db *sql.DB, log logger.Logger) Repository {
	return &repository{
		db:  db,
		log: log,
	}
}

type repository struct {
	db  *sql.DB
	log logger.Logger
}

func (r *repository) GetProfileByID(ctx context.Context, ID int) (Profile, error) {
	q := `SELECT id, name, title, image, twitter_username, twitter_id, description FROM profile WHERE id = $1`
	log := r.log.WithCorrelation(ctx)

	ctx, cancel := context.WithTimeout(ctx, 200000)
	defer cancel()

	resp := r.db.QueryRow(q, ID)

	var p Profile
	err := resp.Scan(
		&p.ID, &p.Name, &p.Title, &p.Image, &p.TwitterUser, &p.TwitterID, &p.Description,
	)
	if err != nil {
		log.Error(err)
		return Profile{}, err
	}

	return p, nil
}

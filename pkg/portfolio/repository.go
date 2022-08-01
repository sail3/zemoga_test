package portfolio

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sail3/zemoga_test/internal/logger"
)

type Repository interface {
	GetProfileByID(ctx context.Context, ID int) (Profile, error)
	GetAllProfile(ctx context.Context) ([]Profile, error)
	UpdateProfile(ctx context.Context, ID int, p Profile) error
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

func (r *repository) GetAllProfile(ctx context.Context) ([]Profile, error) {
	q := "SELECT id, name, title, image, twitter_username, twitter_id, description FROM profile"
	log := r.log.WithCorrelation(ctx)
	profiles := make([]Profile, 0)
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	for rows.Next() {
		var p Profile
		err := rows.Scan(
			&p.ID, &p.Name, &p.Title, &p.Image, &p.TwitterUser, &p.TwitterID, &p.Description,
		)
		if err != nil {
			log.Error(err)
			continue
		}
		profiles = append(profiles, p)
	}
	return profiles, nil
}

func (r *repository) UpdateProfile(ctx context.Context, ID int, p Profile) error {
	q := "UPDATE profile SET name= $1, title= $2 , image= $3, twitter_username= $4, twitter_id= $5, description= $6 WHERE id = $7"
	log := r.log.WithCorrelation(ctx)
	st, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		log.Error(err)
		return err
	}
	defer st.Close()

	result, err := st.ExecContext(ctx,
		p.Name,
		p.Title,
		p.Image,
		p.TwitterUser,
		p.TwitterID,
		p.Description,
		p.ID,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	ra, err := result.RowsAffected()
	if err != nil {
		log.Error(err)
		return err
	}
	if ra < 1 {
		err = fmt.Errorf("psql: expected 1 row affected, got %d", ra)
		log.Error(err)
		return err
	}
	return nil
}

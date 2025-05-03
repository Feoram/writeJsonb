package data

import (
	"database/sql"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) InsertJSONPayload(payload []byte) error {
	_, err := r.db.Exec(
		`INSERT INTO data (payload, created_at) VALUES ($1, $2)`,
		payload,
		time.Now(),
	)
	return err
}

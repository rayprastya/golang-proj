package repository

import (
	"context"
	"database/sql"
	"project/shopifyx/internal/model"
)

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{
		db: db,
	}
}

// GetAlbums retrieves all albums from the database
func (r *AlbumRepository) GetAlbums(ctx context.Context) ([]model.Album, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, artist, price FROM albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []model.Album
	for rows.Next() {
		var a model.Album
		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
			return nil, err
		}
		albums = append(albums, a)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return albums, nil
}

// AddAlbum adds a new album to the database
func (r *AlbumRepository) AddAlbum(ctx context.Context, a model.Album) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO albums (id, title, artist, price) VALUES ($1, $2, $3, $4)",
		a.ID, a.Title, a.Artist, a.Price)
	return err
}

package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) CreateNewList(ctx context.Context, ListTitle string) (*models.List, error) {
	return nil, nil
}

func (d *DAO) InsertMovieToList(ctx context.Context, listID, MovieID, userID uint) error {
	return nil
}

func (d *DAO) RemoveMovieFromList(ctx context.Context, listID, MovieID uint) error {
	return nil
}

func (d *DAO) DeleteList(ctx context.Context, listID uint) error {
	return nil
}

func (d *DAO) GetListInfo(ctx context.Context, listID uint) (*models.List, error) {
	return nil, nil
}

func (d *DAO) GetAllListInfoByUserID(ctx context.Context, userID uint) ([]*models.List, error) {
	return nil, nil
}

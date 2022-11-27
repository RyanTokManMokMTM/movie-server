package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) FindMovieListByGenreID(ctx context.Context, genreID uint) (*models.GenreInfo, error) {
	genreModel := &models.GenreInfo{
		GenreId: genreID,
	}

	if err := genreModel.GetMovieListByGenreID(ctx, d.engine); err != nil {
		return nil, err
	}
	return genreModel, nil
}

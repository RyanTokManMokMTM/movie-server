package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) CountMovieCollected(ctx context.Context, movieId uint) (int64, error) {
	lm := &models.ListMovie{
		MovieInfoId: movieId,
	}

	return lm.CountMovieCollected(ctx, d.engine)
}

func (d *DAO) FindOneMovieFormAnyList(ctx context.Context, movieID, userID uint) (*models.ListMovie, error) {
	lm := &models.ListMovie{
		MovieInfoId: movieID,
	}

	if err := lm.FindOneMovieFromAnyList(ctx, d.engine, userID); err != nil {
		return nil, err
	}
	return lm, nil
}

package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) FindOneUserLikedMovie(ctx context.Context, movieID, userID uint) (*models.UserMovie, error) {
	um := &models.UserMovie{UserId: userID, MovieInfoId: movieID}
	if err := um.FindOneLikedMovie(ctx, d.engine); err != nil {
		return nil, err
	}
	return um, nil
}

func (d *DAO) UpdateUserLikedMovieState(ctx context.Context, um *models.UserMovie) error {
	if err := um.UpdateLikedMovieState(ctx, d.engine); err != nil {
		return err
	}
	return nil
}

func (d *DAO) CountLikesOfMovie(ctx context.Context, movieID uint) (int64, error) {
	um := &models.UserMovie{
		MovieInfoId: movieID,
	}

	return um.CountLikesOfMovie(ctx, d.engine)
}

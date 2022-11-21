package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) FindOneUserLikedMovie(ctx context.Context, movieID, userID uint) (*models.User, error) {
	u := &models.User{
		ID: userID,
	}
	err := u.FindOneLikedMovie(ctx, d.engine, movieID)
	return u, err
}

func (d *DAO) CountLikesOfMovie(ctx context.Context, movieID uint) (int64, error) {
	um := &models.UserMovie{
		MovieInfoId: movieID,
	}

	return um.CountLikesOfMovie(ctx, d.engine)
}

func (d *DAO) RemoveUserLikedMovie(ctx context.Context, movieID, userID uint) error {
	u := &models.User{
		ID: userID,
	}

	return u.RemoveOneLikedMovie(ctx, d.engine, movieID)
}

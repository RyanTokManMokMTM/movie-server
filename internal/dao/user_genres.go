package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) UpdateUserGenres(ctx context.Context, ids []uint, userId uint) error {
	u := models.User{
		Id: userId,
	}

	return u.UpdateUserGenreTrans(ctx, d.engine, ids)
}

func (d *DAO) FindUserGenres(ctx context.Context, userId uint) (*[]models.GenreInfo, error) {
	u := models.User{
		Id: userId,
	}

	return u.FindUserGenres(ctx, d.engine)
}

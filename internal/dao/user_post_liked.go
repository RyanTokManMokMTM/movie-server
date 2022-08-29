package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) DeletePostLikes(ctx context.Context, postLiked *models.PostLiked) error {
	return postLiked.RemovePostLikes(ctx, d.engine)
}

func (d *DAO) FindOnePostLiked(ctx context.Context, userId, postId uint) (*models.PostLiked, error) {
	pl := &models.PostLiked{
		UserId:     userId,
		PostPostId: postId,
	}

	if err := pl.FindOnePostLiked(ctx, d.engine); err != nil {
		return nil, err
	}
	return pl, nil
}

func (d *DAO) CountPostLikes(ctx context.Context, postId uint) (int64, error) {
	pl := &models.PostLiked{
		PostPostId: postId,
	}
	return pl.CountPostLikes(ctx, d.engine)
}

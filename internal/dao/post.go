package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) CreateNewPost(ctx context.Context, title, desc string, userID, movieID uint) (*models.Post, error) {
	newPost := &models.Post{
		PostTitle:   title,
		PostDesc:    desc,
		UserId:      userID,
		MovieInfoId: movieID,
		PostLike:    0,
	}

	if err := newPost.CreateNewPost(ctx, d.engine); err != nil {
		return nil, err
	}
	return newPost, nil
}

func (d *DAO) UpdatePostInfo(ctx context.Context, title, desc string, postID uint) error {
	return nil
}

func (d *DAO) DeletePost(ctx context.Context, postID uint) error {
	return nil
}

func (d *DAO) GetPostInfo(ctx context.Context, postID uint) error {
	return nil
}

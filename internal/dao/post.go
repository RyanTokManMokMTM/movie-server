package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) CreateNewPost(ctx context.Context, post *models.Post) error {
	if err := post.CreateNewPost(ctx, d.engine); err != nil {
		return err
	}
	return nil
}

func (d *DAO) UpdatePostInfo(ctx context.Context, post *models.Post) error {
	return post.UpdatePost(ctx, d.engine)
}

func (d *DAO) DeletePost(ctx context.Context, postID, userID uint) error {
	post := &models.Post{
		PostId: postID,
		UserId: userID,
	}
	return post.DeletePost(ctx, d.engine)
}

func (d *DAO) GetPostInfo(ctx context.Context, postID uint) (*models.Post, error) {
	post := &models.Post{
		PostId: postID,
	}

	if err := post.GetPostInfo(ctx, d.engine); err != nil {
		return nil, err
	}
	return post, nil
}

func (d *DAO) GetAllPostInfo(ctx context.Context) ([]*models.Post, error) {
	post := &models.Post{}
	resp, err := post.GetAllPostInfoByCreateTimeDesc(ctx, d.engine)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (d *DAO) GetAllUserPosts(ctx context.Context, userID uint) ([]*models.Post, error) {
	post := &models.Post{
		UserId: userID,
	}
	resp, err := post.GetUserPostsByCreateTimeDesc(ctx, d.engine)
	if err != nil {
		return nil, err
	}
	return resp, err
}

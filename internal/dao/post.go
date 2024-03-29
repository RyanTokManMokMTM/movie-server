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

func (d *DAO) FindOnePostInfoWithUserLiked(ctx context.Context, postID, userID uint) (*models.Post, error) {
	post := &models.Post{
		PostId: postID,
	}

	if err := post.GetPostInfoWithUserLiked(ctx, d.engine, userID); err != nil {
		return nil, err
	}
	return post, nil
}

func (d *DAO) FindOnePostInfo(ctx context.Context, postID uint) (*models.Post, error) {
	post := &models.Post{
		PostId: postID,
	}

	if err := post.GetPostInfo(ctx, d.engine); err != nil {
		return nil, err
	}
	return post, nil
}

func (d *DAO) FindAllPosts(ctx context.Context, userID uint, limit, offset int) ([]*models.Post, int64, error) {
	post := &models.Post{}
	return post.GetAllPostInfoByCreateTimeDesc(ctx, d.engine, userID, limit, offset)
}

func (d *DAO) FindFollowingPosts(ctx context.Context, userID uint, limit, pageOffset int) ([]*models.Post, int64, error) {
	post := &models.Post{}
	return post.GetFollowPostInfoByCreateTimeDesc(ctx, d.engine, userID, limit, pageOffset)
}

func (d *DAO) FindUserPosts(ctx context.Context, userID, likedBy uint, limit, pageOffset int) ([]*models.Post, int64, error) {
	post := &models.Post{
		UserId: userID,
	}
	return post.GetUserPostsByCreateTimeDesc(ctx, d.engine, likedBy, limit, pageOffset)

}

func (d *DAO) CountUserPosts(ctx context.Context, userID uint) (int64, error) {
	post := &models.Post{UserId: userID}

	return post.CountUserPosts(ctx, d.engine)
}

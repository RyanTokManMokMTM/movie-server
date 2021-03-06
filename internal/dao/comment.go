package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) CreatePostComment(ctx context.Context, userID, PostID uint, comment string) (*models.Comment, error) {
	newComment := &models.Comment{
		PostID:  PostID,
		UserID:  userID,
		Comment: comment,
	}
	if err := newComment.CreatePostComment(ctx, d.engine); err != nil {
		return nil, err
	}
	return newComment, nil
}

func (d *DAO) UpdateComment(ctx context.Context, comment *models.Comment) error {
	return comment.UpdatePostComment(ctx, d.engine)
}

func (d *DAO) DeleteComment(ctx context.Context, commentID uint) error {
	comment := models.Comment{
		CommentID: commentID,
	}

	return comment.DeletePostComment(ctx, d.engine)
}

func (d *DAO) FindPostComments(ctx context.Context, postID uint) ([]*models.Comment, error) {
	comment := models.Comment{
		PostID: postID,
	}

	list, err := comment.FindOnePostComments(ctx, d.engine)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (d *DAO) FindOneComment(ctx context.Context, commentID uint) (*models.Comment, error) {
	comment := &models.Comment{
		CommentID: commentID,
	}

	if err := comment.FindOneComment(ctx, d.engine); err != nil {
		return nil, err
	}

	return comment, nil
}

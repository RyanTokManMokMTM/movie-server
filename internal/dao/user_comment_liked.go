package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

//
//func (d *DAO) UpdateCommentLiked(ctx context.Context, commentLiked *models.CommentLiked, comment *models.Comment) error {
//	return commentLiked.UpdateCommentLiked(ctx, d.engine, comment)
//}
//
//func (d *DAO) FindOneCommentLiked(ctx context.Context, userId, commentId uint) (*models.CommentLiked, error) {
//	cl := &models.CommentLiked{
//		UserId:           userId,
//		CommentCommentId: commentId,
//	}
//
//	if err := cl.FindOneCommentLike(ctx, d.engine); err != nil {
//		return nil, err
//	}
//	return cl, nil
//}
//
func (d *DAO) CountCommentLikes(ctx context.Context, commentId uint) (int64, error) {
	cl := &models.CommentLiked{
		CommentCommentId: commentId,
	}
	return cl.CountCommentLikes(ctx, d.engine)
}

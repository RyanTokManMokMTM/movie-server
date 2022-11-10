package dao

import (
	"context"
	"database/sql"
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

func (d *DAO) CreatePostReplyComment(ctx context.Context, userID, PostID, replyCommentId, parentID, replyUserID uint, comment string) (*models.Comment, error) {
	newComment := &models.Comment{
		PostID:      PostID,
		UserID:      userID,
		Comment:     comment,
		ParentID:    sql.NullInt64{Int64: int64(parentID), Valid: true},
		ReplyTo:     sql.NullInt64{Int64: int64(replyCommentId), Valid: true}, //reply to which comment
		ReplyUserID: sql.NullInt64{Int64: int64(replyUserID), Valid: true},    //reply to who
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

<<<<<<< HEAD
func (d *DAO) FindPostComments(ctx context.Context, postID uint, limit, pageOffset int) ([]*models.Comment, int64, error) {
=======
func (d *DAO) FindPostComments(ctx context.Context, postID, checkUser uint) ([]*models.Comment, error) {
>>>>>>> d91c53fa9425adcc67a3b3f94b556c84f2a1a718
	comment := models.Comment{
		PostID: postID,
	}

<<<<<<< HEAD
	return comment.FindOnePostComments(ctx, d.engine, limit, pageOffset)

}

func (d *DAO) FindReplyComments(ctx context.Context, commentID uint, limit, pageOffset int) ([]*models.Comment, int64, error) {
=======
	list, err := comment.FindOnePostComments(ctx, d.engine, checkUser)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (d *DAO) FindReplyComments(ctx context.Context, parentID, checkUser uint) ([]*models.Comment, error) {
>>>>>>> d91c53fa9425adcc67a3b3f94b556c84f2a1a718
	comment := models.Comment{
		ParentID: sql.NullInt64{Int64: int64(parentID), Valid: true},
	}

<<<<<<< HEAD
	return comment.FindReplyComments(ctx, d.engine, limit, pageOffset)
=======
	return comment.FindReplyParentComments(ctx, d.engine, checkUser)
>>>>>>> d91c53fa9425adcc67a3b3f94b556c84f2a1a718
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

func (d *DAO) UpdateCommentCount(ctx context.Context, comment *models.Comment, updateCount uint) error {
	comment.LikesCount = comment.LikesCount + updateCount

	return comment.UpdateCommentLiked(ctx, d.engine)
}

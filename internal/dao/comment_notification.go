package dao

import (
	"context"
	"database/sql"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"time"
)

func (d *DAO) InsertOneCommentNotification(ctx context.Context, receiverID, commentBy, postID, commentID uint, commentTime time.Time) error {
	notify := &models.CommentNotification{
		PostID:      postID,
		ReceiverId:  receiverID,
		CommentId:   commentID,
		CommentBy:   commentBy,
		CommentTime: commentTime,
		Type:        1,
	}

	return notify.InsertOne(d.engine, ctx)
}

func (d *DAO) InsertOneReplyCommentNotification(ctx context.Context, receiverID, commentBy, postID, commentID, replyCommentID uint, commentTime time.Time) error {
	notify := &models.CommentNotification{
		PostID:         postID,
		ReceiverId:     receiverID,
		CommentId:      commentID,                                                //what is the comment info
		ReplyCommentId: sql.NullInt64{Int64: int64(replyCommentID), Valid: true}, //reply to which comment
		CommentBy:      commentBy,
		CommentTime:    commentTime,
		Type:           2,
	}

	return notify.InsertOne(d.engine, ctx)
}

func (d *DAO) FindOneCommentNotification(ctx context.Context, receiverID uint, limit, pageOffset int) ([]*models.CommentNotification, int64, error) {
	notify := &models.CommentNotification{
		ReceiverId: receiverID,
	}

	return notify.FindNotificationsByReceiver(d.engine, ctx, limit, pageOffset)
}

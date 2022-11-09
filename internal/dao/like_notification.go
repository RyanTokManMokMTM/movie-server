package dao

import (
	"context"
	"database/sql"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"time"
)

func (d *DAO) InsertOnePostLikeNotification(ctx context.Context, postID, likedBy, Receiver uint, likedTime time.Time) error {
	notify := &models.LikeNotification{
		PostID:     postID,
		ReceiverID: Receiver,
		LikedBy:    likedBy,
		LikedTime:  likedTime,
		Type:       1,
	}
	return notify.InsertOne(d.engine, ctx)
}

//InsertOneCommentLikeNotification -liked comment
func (d *DAO) InsertOneCommentLikeNotification(ctx context.Context, postID, likedBy, commentID, Receiver uint, likedTime time.Time) error {
	notify := &models.LikeNotification{
		PostID:     postID,
		CommentId:  sql.NullInt64{Int64: int64(commentID), Valid: true},
		ReceiverID: Receiver,
		LikedBy:    likedBy,
		LikedTime:  likedTime,
		Type:       2,
	}
	return notify.InsertOne(d.engine, ctx)
}

func (d *DAO) FindLikesNotificationByReceiver(ctx context.Context, receiverID uint, limit, pageOffset int) ([]*models.LikeNotification, int64, error) {
	notify := &models.LikeNotification{
		ReceiverID: receiverID,
	}

	return notify.FindNotificationsByReceiver(d.engine, ctx, limit, pageOffset)
}

func (d *DAO) FindOneLikePostNotification(ctx context.Context, receiverID, likedBy, postID uint) error {
	notify := &models.LikeNotification{
		ReceiverID: receiverID,
		LikedBy:    likedBy,
		PostID:     postID,
	}

	return notify.FindOneLikePostNotification(d.engine, ctx)
}

func (d *DAO) FindOneLikeCommentNotification(ctx context.Context, receiverID, likedBy, commentID uint) error {
	notify := &models.LikeNotification{
		ReceiverID: receiverID,
		LikedBy:    likedBy,
		CommentId:  sql.NullInt64{Int64: int64(commentID), Valid: true},
	}

	return notify.FindOneLikeCommentNotification(d.engine, ctx)
}

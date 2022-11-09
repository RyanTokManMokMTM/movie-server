package models

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"time"
)

//There are 2 types of like notification
//1. post like
//2. comment like
//use a `type` field to group/identify them

type LikeNotification struct {
	//PostID
	//CommentID if type is 2
	ID         uint `gorm:"primaryKey"`
	ReceiverID uint
	PostID     uint
	CommentId  sql.NullInt64
	LikedBy    uint
	LikedTime  time.Time
	Type       uint //comment or post
	//State     uint //is liked or removed
	DefaultModel

	PostInfo     Post    `gorm:"foreignKey:PostID;references:PostId"`
	LikedUser    User    `gorm:"foreignKey:LikedBy;references:ID"`
	ReceiverInfo User    `gorm:"foreignKey:ReceiverID;references:ID"`
	CommentInfo  Comment `gorm:"foreignKey:CommentId;references:CommentID"`
}

func (m *LikeNotification) TableName() string {
	return "like_notification"
}

func (m *LikeNotification) InsertOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Create(&m).Error
}

func (m *LikeNotification) FindNotificationsByReceiver(db *gorm.DB, ctx context.Context, limit, pageOffset int) ([]*LikeNotification, int64, error) {
	var list []*LikeNotification
	var count int64 = 0
	if err := db.WithContext(ctx).Debug().Model(&m).Where("receiver_id = ?", m.ReceiverID).
		Preload("PostInfo").
		Preload("PostInfo.MovieInfo").
		Preload("LikedUser").
		Preload("CommentInfo").
		Order("created_at desc").Count(&count).Offset(pageOffset).Limit(limit).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil

}

//FindOneLikePostNotification for post like
func (m *LikeNotification) FindOneLikePostNotification(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Where("receiver_id = ? AND liked_by = ? AND post_id = ? AND type = ?", m.ReceiverID, m.LikedBy, m.PostID, 1).First(&m).Error
}

//FindOneLikeCommentNotification for comment like
func (m *LikeNotification) FindOneLikeCommentNotification(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Where("receiver_id = ? AND liked_by = ? AND comment_id = ? AND type = ?", m.ReceiverID, m.LikedBy, m.CommentId, 2).First(&m).Error
}

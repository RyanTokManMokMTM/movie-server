package models

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"time"
)

//There are 2 different comment
//1.reply to post
//2.reply to comment of comment

//post->comment
//post->comment->reply

type CommentNotification struct {
	//PostID
	//CommentID if type is 2
	ID             uint `gorm:"primaryKey"`
	ReceiverId     uint
	PostID         uint          //comment to which post
	CommentId      uint          //comment to what commentID
	ReplyCommentId sql.NullInt64 `gorm:"null"`
	//commentIDContent     string        //what is the comment ???comment
	CommentBy   uint
	CommentTime time.Time
	Type        uint //comment or post
	DefaultModel

	PostInfo        Post    `gorm:"foreignKey:PostID;references:PostId"`
	CommentInfo     Comment `gorm:"foreignKey:CommentId;references:CommentID"`
	RelyCommentInfo Comment `gorm:"foreignKey:ReplyCommentId;references:CommentID"`
	CommentUser     User    `gorm:"foreignKey:CommentBy;references:ID"`
	ReceiverInfo    User    `gorm:"foreignKey:ReceiverId;references:ID"`
}

func (m *CommentNotification) TableName() string {
	return "comment_notification"
}

func (m *CommentNotification) InsertOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Create(&m).Error
}

func (m *CommentNotification) FindNotificationsByReceiver(db *gorm.DB, ctx context.Context) ([]*CommentNotification, error) {
	var list []*CommentNotification
	if err := db.WithContext(ctx).Debug().Where("receiver_id = ?", m.ReceiverId).
		Preload("PostInfo").
		Preload("PostInfo.MovieInfo").
		Preload("CommentInfo").
		Preload("RelyCommentInfo").
		Preload("CommentUser").
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil

}

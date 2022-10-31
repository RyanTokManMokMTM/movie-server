package models

import (
	"database/sql"
	"time"
)

//There are 2 different comment
//1.reply to post
//2.reply to comment of comment

type CommentNotification struct {
	//PostID
	//CommentID if type is 2
	ID          uint          `gorm:"primaryKey"`
	PostID      uint          //comment to which post
	CommentID   sql.NullInt64 `gorm:"null"` //comment to what commentID
	Content     string        //what is the comment
	CommentBy   uint
	CommentTime time.Time
	Type        uint //comment or post
	DefaultModel

	PostInfo    Post    `gorm:"foreignKey:PostID;references:PostId"`
	CommentInfo Comment `gorm:"foreignKey:CommentID;references:CommentID"`
	LikedUser   User    `gorm:"foreignKey:CommentBy;references:ID"`
}

func (m *CommentNotification) TableName() string {
	return "comment_notification"
}

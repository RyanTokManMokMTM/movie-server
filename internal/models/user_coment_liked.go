package models

import (
	"context"
	"gorm.io/gorm"
)

type CommentLiked struct {
	UserId           uint `gorm:"primaryKey"` // User is following FriendTemp
	CommentCommentId uint `gorm:"primaryKey"` // FriendTemp is followed by User
	//State            uint `gorm:"not null;unsigned;type:tinyint(1)"` //remove it?????????
	//DefaultModel
}

func (m *CommentLiked) TableName() string {
	return "comment_liked"
}

//func (m *CommentLiked) BeforeCreate(db *gorm.DB) error {
//	m.State = 1 //when create set to follow
//	return nil
//}
//
//func (m *CommentLiked) FindOneCommentLike(ctx context.Context, db *gorm.DB) error {
//	return db.Debug().WithContext(ctx).Model(&m).First(&m).Error
//}
//

func (m *CommentLiked) CountCommentLikes(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Debug().WithContext(ctx).Model(&m).Where("comment_comment_id = ?", m.CommentCommentId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

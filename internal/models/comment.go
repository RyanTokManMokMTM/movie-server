package models

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Comment struct {
	CommentID uint          `gorm:"primaryKey;not null;autoIncrement"`
	PostID    uint          `gorm:"not null;type:bigint;unsigned;"`
	UserID    uint          `gorm:"not null;type:bigint;unsigned"`
	Comment   string        `gorm:"not null;type:longtext"`
	ReplyTo   sql.NullInt64 `gorm:"null;type:bigint;unsigned"` //if this field is null ,it means not a reply message

	User     User      `gorm:"foreignKey:UserID;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments []Comment `gorm:"foreignKey:ReplyTo;references:CommentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // a list of reply comment

	DefaultModel
}

func (m *Comment) TableName() string {
	return "comments"
}

func (m *Comment) CreatePostComment(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Create(m).Error
}

func (m *Comment) UpdatePostComment(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Updates(m).Error
}

func (m *Comment) DeletePostComment(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Where("comment_id = ?", m.CommentID).Delete(m).Error
}

func (m *Comment) FindOnePostComments(ctx context.Context, db *gorm.DB) ([]*Comment, error) {
	var comments []*Comment
	if err := db.Debug().WithContext(ctx).Model(m).Where("post_id = ? AND reply_to IS NULL", m.PostID).Preload("User").Preload("Comments", func(tx *gorm.DB) *gorm.DB {
		return db.Preload("User")
	}).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (m *Comment) FindReplyComments(ctx context.Context, db *gorm.DB) ([]*Comment, error) {
	var replyComments []*Comment
	if err := db.Debug().WithContext(ctx).Where("reply_to = ?", m.ReplyTo).Preload("User").Find(&replyComments).Error; err != nil {
		return nil, err
	}
	return replyComments, nil
}

func (m *Comment) FindOneComment(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Model(&m).First(&m).Error
}

//Upcoming Feature
//Reply Comment
//Like Comment

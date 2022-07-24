package models

import (
	"context"
	"gorm.io/gorm"
)

type Comment struct {
	CommentID uint   `gorm:"primaryKey;not null;autoIncrement"`
	PostID    uint   `gorm:"not null;type:bigint;unsigned;"`
	UserID    uint   `gorm:"not null;type:bigint;unsigned"`
	Comment   string `gorm:"not null;type:longtext"`

	//Post Post `gorm:"foreignKey:PostID;references:PostId"`
	User User `gorm:"foreignKey:UserID;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
	if err := db.Debug().WithContext(ctx).Model(m).Where("post_id = ?", m.PostID).Preload("User").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (m *Comment) FindOneComment(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Where("comment_id = ?", m.CommentID).First(m).Error
}

//Upcoming Feature
//Reply Comment
//Like Comment

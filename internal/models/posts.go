package models

import (
	"context"
	"gorm.io/gorm"
)

type Post struct {
	PostId      uint   `gorm:"primaryKey;not null;autoIncrement"`
	PostTitle   string `gorm:"not null;type:varchar(255)"`
	PostDesc    string `gorm:"not null;type:varchar(255)"`
	UserId      uint   `gorm:"not null;type:bigint;unsigned;"`
	MovieInfoId uint   `gorm:"not null;type:bigint;unsigned;"`
	PostLike    int64  `gorm:"not null;type:bigint;unsigned;default:0"`

	MovieInfo MovieInfo `gorm:"foreignKey:MovieInfoId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments  []Comment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DefaultModel
}

func (m *Post) TableName() string {
	return "posts"
}

func (m *Post) CreateNewPost(ctx context.Context, db *gorm.DB) error {
	if err := db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *Post) UpdatePost(ctx context.Context, db *gorm.DB) error {
	if err := db.WithContext(ctx).Where("post_id = ?", m.PostId).Updates(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *Post) DeletePost(ctx context.Context, db *gorm.DB) error {
	if err := db.WithContext(ctx).Where("post_id=?", m.PostId).Delete(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *Post) GetPostInfo(ctx context.Context, db *gorm.DB) error {
	if err := db.WithContext(ctx).Where("post_id = ?", m.PostId).First(&m).Error; err != nil {
		return err
	}
	return nil
}

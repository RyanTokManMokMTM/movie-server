package models

import (
	"context"
	"gorm.io/gorm"
)

type PostLiked struct {
	UserId     uint `gorm:"primaryKey"` // User is following FriendTemp
	PostPostId uint `gorm:"primaryKey"` // FriendTemp is followed by User
	//State      uint `gorm:"not null;unsigned;type:tinyint(1)"`
	//DefaultModel
}

func (m *PostLiked) TableName() string {
	return "post_liked"
}

//func (m *PostLiked) BeforeCreate(db *gorm.DB) error {
//	m.State = 1 //when create set to follow
//	return nil
//}

//func (m *PostLiked) UpdatePostLiked(ctx context.Context, db *gorm.DB) error {
//	return db.Debug().WithContext(ctx).Model(&m).Update("state", m.State).Error
//}

func (m *PostLiked) RemovePostLikes(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Model(&m).Delete(&m).Error
}

func (m *PostLiked) FindOnePostLiked(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Model(&m).First(&m).Error
}

func (m *PostLiked) CountPostLikes(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	//if err := db.Debug().WithContext(ctx).Model(&m).Where("post_post_id = ? AND state = 1", m.PostPostId).Count(&count).Error; err != nil {
	//	return 0, err
	//}
	if err := db.Debug().WithContext(ctx).Model(&m).Where("post_post_id = ?", m.PostPostId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

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
	UserInfo  User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments  []Comment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	DefaultModel
}

func (m *Post) TableName() string {
	return "posts"
}

//Create a new Post
func (m *Post) CreateNewPost(ctx context.Context, db *gorm.DB) error {
	if err := db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}
	return nil
}

//Update an existing post
func (m *Post) UpdatePost(ctx context.Context, db *gorm.DB) error {
	if err := db.WithContext(ctx).Where("post_id = ? and user_id = ?", m.PostId, m.UserId).Updates(&m).Error; err != nil {
		return err
	}
	return nil
}

//Delete an existing post
func (m *Post) DeletePost(ctx context.Context, db *gorm.DB) error {
	if err := db.Debug().WithContext(ctx).Model(&m).Where("post_id= ? and user_id = ?", m.PostId, m.UserId).Delete(&m).Error; err != nil {
		return err
	}
	return nil
}

//Get PostInfo by postID
func (m *Post) GetPostInfo(ctx context.Context, db *gorm.DB) error {
	if err := db.Debug().WithContext(ctx).Model(&m).Preload("MovieInfo").Preload("UserInfo").Preload("Comments").First(&m).Error; err != nil {
		return err
	}
	return nil
}

//Get All PostInfo - 10 record by recent 10
func (m *Post) GetAllPostInfoByCreateTimeDesc(ctx context.Context, db *gorm.DB, userID uint) ([]*Post, error) {
	var resp []*Post
	if err := db.Debug().WithContext(ctx).Model(&m).Preload("MovieInfo").Preload("UserInfo").Preload("Comments").Where("user_id != ?", userID).Order("created_at desc").Limit(10).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *Post) GetFollowPostInfoByCreateTimeDesc(ctx context.Context, db *gorm.DB, userID uint) ([]*Post, error) {
	var resp []*Post
	if err := db.Debug().WithContext(ctx).Model(&m).Preload("MovieInfo").Preload("UserInfo").Preload("Comments").Where("user_id = ?", userID).Order("created_at desc").Limit(10).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

//Get All PostInfo - 10 record by recent 10
func (m *Post) GetPostInfoByPostID(ctx context.Context, db *gorm.DB) error {
	if err := db.Debug().WithContext(ctx).Model(&m).Preload("MovieInfo").Preload("UserInfo").Preload("Comments").Where("post_id = ?", m.PostId).First(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *Post) GetUserPostsByCreateTimeDesc(ctx context.Context, db *gorm.DB) ([]*Post, error) {
	var resp []*Post
	if err := db.Debug().WithContext(ctx).Preload("MovieInfo").Preload("UserInfo").Preload("Comments").Where("user_id = ?", m.UserId).Order("created_at desc").Limit(10).Find(&resp).Limit(10).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *Post) CountUserPosts(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Debug().WithContext(ctx).Model(&m).Where("user_id = ?", m.UserId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

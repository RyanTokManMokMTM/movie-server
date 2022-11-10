package models

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Comment struct {
	CommentID   uint          `gorm:"primaryKey;not null;autoIncrement"`
	PostID      uint          `gorm:"not null;type:bigint;unsigned;"`
	UserID      uint          `gorm:"not null;type:bigint;unsigned"`
	Comment     string        `gorm:"not null;type:longtext"`
	ReplyTo     sql.NullInt64 `gorm:"null;type:bigint;unsigned"` //if this field is null ,it means not a reply message
	ReplyUserID sql.NullInt64 `gorm:"null;type:bigint;unsigned"`
	ParentID    sql.NullInt64 `gorm:"null;type:bigint;unsigned"` //this field uses to group reply message to its parent, root won't have parent node
	LikesCount  uint

	User        User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReplyToInfo User      `gorm:"foreignKey:ReplyUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:set null"`
	Comments    []Comment `gorm:"foreignKey:ParentID;references:CommentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // a list of reply comment
	PostInfo    Post      `gorm:"foreignKey:PostID;references:PostId ;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LikedUser   []User    `gorm:"many2many:comment_liked"`
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

func (m *Comment) FindOnePostComments(ctx context.Context, db *gorm.DB, limit, pageOffset int) ([]*Comment, int64, error) {
	var comments []*Comment
	var count int64 = 0
	if err := db.Debug().WithContext(ctx).Model(m).
		Where("post_id = ? AND reply_to IS NULL", m.PostID).
		Preload("User").
		Preload("Comments", func(tx *gorm.DB) *gorm.DB {
			return db.Preload("User")
		}).
		Order("created_at desc").
		Count(&count).Offset(pageOffset).Limit(limit).
		Find(&comments).Error; err != nil {
		return nil, 0, err
	}
	return comments, count, nil
}

func (m *Comment) FindReplyComments(ctx context.Context, db *gorm.DB, limit, pageOffset int) ([]*Comment, int64, error) {
	var replyComments []*Comment
	var count int64 = 0
	if err := db.Debug().WithContext(ctx).Where("reply_to = ?", m.ReplyTo).Preload("User").Order("created_at desc").Count(&count).Offset(pageOffset).Limit(limit).Find(&replyComments).Error; err != nil {
		return nil, 0, err
	}
	return replyComments, count, nil
}

func (m *Comment) FindOneComment(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Model(&m).Preload("PostInfo").First(&m).Error
}

func (m *Comment) UpdateCommentLiked(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Model(&m).Update("LikesCount", m.LikesCount).Error
}

//Upcoming Feature
//Reply Comment
//Like Comment

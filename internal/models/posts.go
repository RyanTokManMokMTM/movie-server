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
	//PostLike    int64  `gorm:"not null;type:bigint;unsigned;default:0"`

	MovieInfo  MovieInfo `gorm:"foreignKey:MovieInfoId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserInfo   User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments   []Comment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PostsLiked []User    `gorm:"many2many:post_liked;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
	return db.Transaction(func(tx *gorm.DB) error {
		////we need to remove all comments and all liked belongs to this post
		//
		//TODO: Delete all likes belongs to this post
		if err := tx.Debug().WithContext(ctx).Model(&m).Association("PostsLiked").Clear(); err != nil {
			return err
		}

		//TODO: Delete all comments belongs to this post
		//TODO: Get all comments
		var allComments []Comment
		if err := tx.Debug().WithContext(ctx).Model(&Comment{PostID: m.PostId}).Find(&allComments).Error; err != nil {
			return err
		}

		//TODO:Remove all comments
		if err := tx.Debug().WithContext(ctx).Delete(&Comment{}, allComments).Error; err != nil {
			return err
		}

		//TODO: Delete the post!
		return tx.Debug().WithContext(ctx).Model(&m).Delete(&m).Error

	})

}

//Get PostInfo by postID
func (m *Post) GetPostInfoWithUserLiked(ctx context.Context, db *gorm.DB, userID uint) error {
	if err := db.Debug().
		WithContext(ctx).
		Model(&m).
		Preload("MovieInfo").
		Preload("UserInfo").
		Preload("Comments").
		Preload("PostsLiked", func(db *gorm.DB) *gorm.DB {
			return db.Find(&User{ID: userID})
		}).First(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *Post) GetPostInfo(ctx context.Context, db *gorm.DB) error {
	if err := db.Debug().
		WithContext(ctx).
		Model(&m).
		Preload("MovieInfo").
		Preload("UserInfo").
		Preload("Comments").
		Preload("PostsLiked").First(&m).Error; err != nil {
		return err
	}
	return nil
}

//Get All PostInfo - 10 record by recent 10
func (m *Post) GetAllPostInfoByCreateTimeDesc(ctx context.Context, db *gorm.DB, userID uint, limit, pageOffset int) ([]*Post, int64, error) {
	//exclude following user
	var friends []uint
	var count int64 = 0
	//get friend list
	fd := &User{ID: userID}
	list, err := fd.GetFriendsList(db, ctx)
	if err != nil {
		return nil, 0, err
	}

	for _, info := range list {
		friends = append(friends, info.ID)
	}

	friends = append(friends, userID) //include itself???

	//logx.Info("friend ids", list)
	//TODO: Get total
	var resp []*Post
	if err := db.Debug().
		WithContext(ctx).
		Model(&m).
		Preload("MovieInfo").
		Preload("UserInfo").
		Preload("Comments").
		Preload("PostsLiked", func(db *gorm.DB) *gorm.DB {
			return db.Find(&User{ID: userID})
		}).
		Where("user_id NOT IN(?)", friends).
		Order("created_at desc").Count(&count).Offset(pageOffset).Limit(limit).Omit("state").Find(&resp).Error; err != nil {
		return nil, 0, err
	}
	return resp, count, nil
}

func (m *Post) GetFollowPostInfoByCreateTimeDesc(ctx context.Context, db *gorm.DB, userID uint, limit, pageOffset int) ([]*Post, int64, error) {
	var resp []*Post
	var count int64 = 0

	var friends []uint
	fd := &User{ID: userID}
	list, err := fd.GetFriendsList(db, ctx)
	if err != nil {
		return nil, 0, err
	}

	for _, info := range list {
		friends = append(friends, info.ID)
	}

	friends = append(friends, userID)

	if err := db.Debug().WithContext(ctx).Model(&m).
		Preload("MovieInfo").
		Preload("UserInfo").
		Preload("Comments").
		Preload("PostsLiked", func(db *gorm.DB) *gorm.DB {
			return db.Find(&User{ID: userID})
		}).
		Where("user_id IN (?)", friends).
		Count(&count).Offset(pageOffset).Order("created_at desc").Limit(limit).Find(&resp).Error; err != nil {
		return nil, 0, err
	}
	return resp, count, nil
}

//Get All PostInfo - 10 record by recent 10
func (m *Post) GetPostInfoByPostID(ctx context.Context, db *gorm.DB) error {
	if err := db.Debug().WithContext(ctx).Model(&m).Preload("MovieInfo").Preload("UserInfo").Preload("Comments").Preload("PostsLiked", func(tx *gorm.DB) *gorm.DB {
		return db.Debug().WithContext(ctx).Where("state = ?", 1)
	}).Where("post_id = ?", m.PostId).First(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *Post) GetUserPostsByCreateTimeDesc(ctx context.Context, db *gorm.DB, likedBy uint, limit, pageOffset int) ([]*Post, int64, error) {
	var resp []*Post
	var count int64 = 0
	if err := db.Debug().
		WithContext(ctx).
		Model(&m).
		Preload("MovieInfo").
		Preload("UserInfo").
		Preload("Comments").
		Preload("PostsLiked", func(db *gorm.DB) *gorm.DB {
			return db.Find(&User{ID: likedBy})
		}).Where("user_id = ?", m.UserId).Count(&count).Order("created_at desc").Offset(pageOffset).Limit(limit).Find(&resp).Limit(10).Error; err != nil {
		return nil, 0, err
	}
	return resp, count, nil
}

func (m *Post) CountUserPosts(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Debug().WithContext(ctx).Model(&m).Where("user_id = ?", m.UserId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

package models

import (
	"context"
	"gorm.io/gorm"
)

type ListMovie struct {
	ListListId  uint `gorm:"primaryKey"`
	MovieInfoId uint `gorm:"primaryKey"`
	//DefaultModel
}

func (m *ListMovie) TableName() string {
	return "lists_movies"
}

func (m *ListMovie) CountMovieCollected(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Debug().WithContext(ctx).Model(&m).Where("movie_info_id = ?", m.MovieInfoId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (m *ListMovie) CountMovieCollectedByUser(ctx context.Context, db *gorm.DB, userID uint) (int64, error) {
	var count int64

	listIds, err := (&List{UserId: userID}).GetUserListsID(ctx, db)
	if err != nil {
		return 0, err
	}

	if err := db.Debug().WithContext(ctx).Model(&m).Where("list_list_id IN (?)", listIds).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (m *ListMovie) FindOneMovieFromAnyList(ctx context.Context, db *gorm.DB, userId uint) error {
	//we need to know all list user have
	//get list id from list model ->
	var userListID []uint
	if err := db.Debug().WithContext(ctx).Model(&List{}).Select("list_id").Where("user_id = ?", userId).Find(&userListID).Error; err != nil {
		return err
	}

	//this query will query all list of the user and get a movie by movie id
	return db.Debug().WithContext(ctx).Model(&m).Where("list_list_id  IN (?)", userListID).Where("movie_info_id = ?", m.MovieInfoId).First(&m).Error
}

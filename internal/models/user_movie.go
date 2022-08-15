package models

import (
	"context"
	"gorm.io/gorm"
)

type UserMovie struct {
	UserId      uint `gorm:"primaryKey,not null"`
	MovieInfoId uint `gorm:"primaryKey,not null"`
	State       uint `gorm:"not null;unsigned;type:tinyint(1)"`
	DefaultModel
}

func (m *UserMovie) TableName() string {
	return "users_movies"
}

func (m *UserMovie) BeforeCreate(db *gorm.DB) error {
	m.State = 1
	return nil
}

func (m *UserMovie) FindOneLikedMovie(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Model(&UserMovie{}).Where("user_id = ? and movie_info_id = ?", m.UserId, m.MovieInfoId).First(&m).Error
}

func (m *UserMovie) UpdateLikedMovieState(ctx context.Context, db *gorm.DB) error {
	//here we just need to update our state
	return db.Debug().WithContext(ctx).Model(&m).Where("user_id = ? and movie_info_id = ?", m.UserId, m.MovieInfoId).Update("State", m.State).Error
}

func (m *UserMovie) CountLikesOfMovie(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Debug().WithContext(ctx).Model(&m).Where("movie_info_id = ?", m.MovieInfoId).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

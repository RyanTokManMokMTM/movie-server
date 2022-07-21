package models

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GenreInfo struct {
	GenreId uint   `gorm:"primaryKey;not null;autoIncrement"`
	Name    string `gorm:"not null"`

	MovieInfo []MovieInfo `gorm:"many2many:genres_movies"`
	DefaultModel
}

func (m *GenreInfo) TableName() string {
	return "genre_infos"
}

func (m *GenreInfo) GetGenreInfos(ctx context.Context, db gorm.DB) ([]*GenreInfo, error) {
	var resp []*GenreInfo
	if err := db.WithContext(ctx).Model(m).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *GenreInfo) GetGenreInfoByID(ctx context.Context, db gorm.DB) (*GenreInfo, error) {
	var resp *GenreInfo

	if err := db.WithContext(ctx).Model(m).Where("genre_id = ?", m.GenreId).First(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *GenreInfo) GetMovieListByGenreID(ctx context.Context, db *gorm.DB) error {
	logx.Info("Genre DB - Get Movie List By GenreID")
	if err := db.Debug().WithContext(ctx).Where("genre_id = ?", m.GenreId).Preload("MovieInfo", func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(10)
	}).Find(&m).Error; err != nil {
		logx.Info("Genre DB - Get Movie List By GenreID err :", err)
		return err
	}
	return nil
}

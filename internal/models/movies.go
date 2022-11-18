package models

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type MovieInfo struct {
	Adult            bool        `json:"adult" gorm:"not null"`
	BackdropPath     string      `json:"backdrop_path" gorm:"not null"`
	Id               uint        `json:"id" gorm:"primaryKey" gorm:"not null"`
	OriginalLanguage string      `json:"original_language" gorm:"not null"`
	OriginalTitle    string      `json:"original_title" gorm:"not null"`
	Overview         string      `json:"overview" gorm:"not null"`
	Popularity       float64     `json:"popularity" gorm:"not null"`
	PosterPath       string      `json:"poster_path" gorm:"not null"`
	ReleaseDate      string      `json:"release_date" gorm:"not null"`
	Title            string      `json:"title" gorm:"not null"`
	RunTime          int64       `json:"runtime" gorm:"not null"`
	Video            bool        `json:"video" gorm:"not null"`
	VoteAverage      float64     `json:"vote_average" gorm:"not null"`
	VoteCount        int64       `json:"vote_count" gorm:"not null"`
	GenreInfo        []GenreInfo `gorm:"many2many:genres_movies;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Lists            []List      `gorm:"many2many:lists_movies;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DefaultModel
}

func (m *MovieInfo) TableName() string {
	return "movie_infos"
}

func (m *MovieInfo) GetMoviesInfoByID(ctx context.Context, db *gorm.DB) ([]*MovieInfo, error) {
	return nil, nil
}

func (m *MovieInfo) FindOneMovieWithGenres(ctx context.Context, db *gorm.DB) error {
	logx.Info("MovieDB - Get Movie Detail")
	if err := db.Debug().WithContext(ctx).Model(&m).Where("movie_id = ?", m.Id).Preload("GenreInfo").Find(&m).Error; err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (m *MovieInfo) FindOneMovie(ctx context.Context, db *gorm.DB) error {
	logx.Info("MovieDB - Get Movie")
	if err := db.Debug().WithContext(ctx).Model(&m).First(&m).Error; err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

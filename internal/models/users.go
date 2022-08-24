package models

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"primaryKey;not null;autoIncrement"`
	Name     string `gorm:"not null;type:varchar(64)"`
	Email    string `gorm:"not null;type:varchar(64)"`
	Password string `gorm:"not null;type:varchar(64)"`
	Avatar   string `gorm:"not null;type:varchar(255)"`
	Cover    string `gorm:"not null;type:varchar(255)"`

	//can have a lot of list
	List       []List      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	MovieInfos []MovieInfo `gorm:"many2many:users_movies;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	//this relationship for following and follower?
	Friends []User `gorm:"many2many:friends;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	//Genres []GenreInfo `gorm` //one u
	//use may have a lot of post
	Posts []Post `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DefaultModel
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) Insert(ctx context.Context, db *gorm.DB) error {
	logx.Infof("UserDB - Create User:%+v \n", m)
	if err := db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *User) FindOneByID(ctx context.Context, db *gorm.DB) error {
	logx.Infof("UserDB - Find One By ID:%+v \n", m)
	if err := db.WithContext(ctx).Model(&m).Where("id = ?", m.Id).First(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *User) FindOneByEmail(ctx context.Context, db *gorm.DB) error {
	logx.Infof("UserDB - Find One By Email:%+v\n", m)
	if err := db.WithContext(ctx).Model(&m).Where("email = ?", m.Email).First(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *User) UpdateInfo(ctx context.Context, db *gorm.DB) error {
	logx.Infof("UserDB - Update Info:%+v \n", m)
	if err := db.WithContext(ctx).Model(&m).Where("id = ?", m.Id).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (m *User) CreateLikedMovie(ctx context.Context, db *gorm.DB, movie *MovieInfo) error {
	logx.Infof("UserDB - Create User Liked Movie:%+v \n", m)
	return db.WithContext(ctx).Model(&m).Association("MovieInfos").Append(movie)

}

func (m *User) UpdateLikedMovie(ctx context.Context, db *gorm.DB, movie *MovieInfo) error {
	logx.Infof("UserDB - Remove User Liked Movie:%+v \n", m)
	return db.WithContext(ctx).Model(&m).Association("MovieInfos").Delete(movie)
}

func (m *User) GetUserLikedMovies(ctx context.Context, db *gorm.DB) error {
	logx.Infof("UserDB - User Liked Movies:%+v \n", m)
	if err := db.Debug().WithContext(ctx).Preload("MovieInfos", func(db *gorm.DB) *gorm.DB {
		return db.Select("movie_infos.*").Joins("left join users_movies on users_movies.movie_info_id = movie_infos.id").Where("users_movies.state = ?", 1)
	}).Preload("MovieInfos.GenreInfo").Where("id = ?", m.Id).Find(&m).Error; err != nil {
		return err
	}
	return nil
}

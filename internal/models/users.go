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
	List        []List      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LikedMovies []MovieInfo `gorm:"many2many:liked_movies;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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

func (m *User) AddLikedMovie(ctx context.Context, db *gorm.DB, movie *MovieInfo) error {
	logx.Infof("UserDB - Add User Liked Movie:%+v \n", m)
	return db.WithContext(ctx).Model(&m).Association("LikedMovies").Append(movie)
}

func (m *User) RemoveLikedMovie(ctx context.Context, db *gorm.DB, movie *MovieInfo) error {
	logx.Infof("UserDB - Remove User Liked Movie:%+v \n", m)
	return db.WithContext(ctx).Model(&m).Association("LikedMovies").Delete(movie)
}

func (m *User) GetUserLikedMovies(ctx context.Context, db *gorm.DB) error {
	logx.Infof("UserDB - User Liked Movies:%+v \n", m)
	if err := db.Debug().WithContext(ctx).Preload("LikedMovies").Preload("LikedMovies.GenreInfo").Where("id = ?", m.Id).Find(&m).Error; err != nil {
		return err
	}
	return nil
}

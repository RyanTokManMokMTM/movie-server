package models

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type User struct {
	Id       uint `gorm:"primaryKey;not null;autoIncrement"`
	Name     string
	Email    string
	Password string
	Avatar   string
	Cover    string

	//can have a lot of list
	List []List `gorm:"foreignKey:UserId;references:Id"`

	//use may have a lot of post
	//Post []Post `gorm:"foreignKey:UserID;references:Id"`
	DefaultModel
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Insert(ctx context.Context, engine *gorm.DB) error {
	logx.Infof("UserDB - Create User:%+v \n", u)
	if err := engine.WithContext(ctx).Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) FindOneByID(ctx context.Context, engine *gorm.DB) error {
	logx.Infof("UserDB - Find One By ID:%+v \n", u)
	if err := engine.WithContext(ctx).Model(&u).Where("id = ?", u.Id).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) FindOneByEmail(ctx context.Context, engine *gorm.DB) error {
	logx.Infof("UserDB - Find One By Email:%+v\n", u)
	if err := engine.WithContext(ctx).Model(&u).Where("email = ?", u.Email).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateInfo(ctx context.Context, engine *gorm.DB) error {
	logx.Infof("UserDB - Update Info:%+v \n", u)
	if err := engine.WithContext(ctx).Model(&u).Where("id = ?", u.Id).Updates(u).Error; err != nil {
		return err
	}
	return nil
}

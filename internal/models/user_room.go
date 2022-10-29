package models

import (
	"context"
	"gorm.io/gorm"
)

type UsersRooms struct {
	RoomID   uint `gorm:"primaryKey"`
	UserID   uint `gorm:"primaryKey"`
	IsActive bool `gorm:"is_active;default:false"`
}

//There may a lot of user inside the same room
/*
For example
RoomID : x1
A :x1
B :x1
C :x1

that means A,B and C in the same chat group


*/

func (ur *UsersRooms) TableName() string {
	return "users_rooms"
}

func (ur *UsersRooms) GetRoomUsers(db *gorm.DB, ctx context.Context) ([]uint, error) {
	var allUser []uint
	if err := db.WithContext(ctx).Debug().Model(ur).Select("user_id").Where("room_id = ?", ur.RoomID).Find(&allUser).Error; err != nil {
		return nil, err
	}

	return allUser, nil
}

func (ur *UsersRooms) FindOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().First(&ur).Error
}

func (ur *UsersRooms) SetActiveState(db *gorm.DB, ctx context.Context) error {
	//only update is not the same state
	return db.WithContext(ctx).Debug().Model(&ur).Where("is_active = ?", !ur.IsActive).Update("IsActive", ur.IsActive).Error
}

func (ur *UsersRooms) GetUserActiveRoom(db *gorm.DB, ctx context.Context) ([]*UsersRooms, error) {
	var activeRooms []*UsersRooms
	if err := db.WithContext(ctx).Debug().Model(&UsersRooms{}).Where("user_id = ? AND is_active = ? ", ur.UserID, true).Find(&activeRooms).Error; err != nil {
		return nil, err
	}

	return activeRooms, nil
}

//
//func (ur *UsersRooms) GetUserRoomState(db *gorm.DB, ctx context.Context) (bool, error) {
//	if err := db.WithContext(ctx).Debug().First(&ur).Error; err != nil {
//		return false, err
//	}
//
//	return ur.IsActive, nil
//}

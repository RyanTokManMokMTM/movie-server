package models

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type Room struct {
	ID       uint   `gorm:"primary_key"` //room id
	Name     string `gorm:"null"`        //room name
	Info     string `gorm:"null"`        //room info
	OwnerRef uint
	Owner    User   `gorm:"foreignKey:OwnerRef"`
	Users    []User `gorm:"many2many:users_rooms;"`

	LastSender sql.NullInt64 `gorm:"null"` // if the field is null,it means no sender in this room
	Sender     User          `gorm:"foreignKey:LastSender;constraint:OnUpdate:CASCADE,OnDelete:set null"`
	Messages   []Message

	IsRead bool // is receiver read the message? - default is false
	DefaultModel
}

/*
Single Chat(Peer to Peer)

*/

func (r *Room) TableName() string {
	return "rooms"
}

func (r *Room) InsertOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Create(&r).Error
}

func (r *Room) RemoveOne(db *gorm.DB, ctx context.Context) error {

	return db.WithContext(ctx).Debug().Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Debug().Model(&r).Association("Users").Clear(); err != nil {
			logx.Error(err)
			return err
		}

		if err := tx.WithContext(ctx).Debug().Delete(&r).Error; err != nil {
			logx.Error(err)
			return err
		}
		return nil
	})
}

func (r *Room) FindOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().First(&r).Error
}

func (r *Room) FindOneWithInfo(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Model(&r).Preload("Users").Preload("Messages", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("sent_time desc").Limit(10)
	}).Preload("Messages.SendUser").First(&r).Error
}

func (r *Room) InsertOneUser(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Debug().Model(&r).Association("Users").Append(user)
}

func (r *Room) RemoveOneUser(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Debug().Model(&r).Association("Users").Delete(user)
}

func (r *Room) FindRoomMembers(db *gorm.DB, ctx context.Context) ([]*User, error) {
	var members []*User
	err := db.WithContext(ctx).Debug().Model(&r).Association("Users").Find(&members)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (r *Room) FindOneRoomMember(db *gorm.DB, ctx context.Context, userID uint) (*User, error) {
	var members *User
	err := db.WithContext(ctx).Debug().Model(&r).Where("user_id = ?", userID).Association("Users").Find(&members)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (r *Room) UpdateLastSender(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Model(&r).Update("LastSender", r.LastSender).Error
}

func (r *Room) UpdateIsReadState(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Model(&r).Where("is_read != ?", r.IsRead).Update("IsRead", r.IsRead).Error
}

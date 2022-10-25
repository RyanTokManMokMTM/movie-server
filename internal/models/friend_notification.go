package models

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type FriendNotification struct {
	ID       uint `gorm:"primary_key"`
	Sender   uint
	Receiver uint
	State    bool //0: accepted/declined/canceled 1:waiting/sent
	DefaultModel

	SenderInfo   User `gorm:"foreignKey:Sender;references:ID"`
	ReceiverInfo User `gorm:"foreignKey:Receiver;references:ID"`
}

func (f *FriendNotification) TableName() string {
	return "friend_notification"
}

func (f *FriendNotification) InsertOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Create(&f).Error
}

func (f *FriendNotification) FineOneByID(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Where("ID = ? AND State = ?", f.ID, f.State).First(&f).Error
}

func (f *FriendNotification) FineOneBySenderAndReceiver(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Where("(Sender = ? AND Receiver = ?) OR (Sender = ? AND Receiver = ?) AND State = ?", f.Sender, f.Receiver, f.Receiver, f.Sender, f.State).First(&f).Error
}

func (f *FriendNotification) Accept(db *gorm.DB, ctx context.Context) error {
	//calling FriendTemp model

	return db.WithContext(ctx).Debug().Transaction(func(tx *gorm.DB) error {
		/*
			TODO: Update Notification State
			TODO: Add FriendTemp to Friendship of both of them
		*/

		if err := tx.WithContext(ctx).Debug().Model(&f).Update("State", false).Error; err != nil {
			logx.Error(err)
			return err
		}

		if err := tx.WithContext(ctx).Debug().Model(&User{
			ID: f.Sender,
		}).Association("Friends").Append(&User{ID: f.Receiver}); err != nil {
			logx.Error(err)
			return err
		}

		if err := tx.WithContext(ctx).Debug().Model(&User{
			ID: f.Receiver,
		}).Association("Friends").Append(&User{ID: f.Sender}); err != nil {
			logx.Error(err)
			return err
		}

		//TODO: Creating the room
		//TODO: Insert both user with new roomID!
		r := &Room{OwnerRef: f.Sender}
		if err := r.InsertOne(tx, ctx); err != nil {
			logx.Error("create room error : ", err)
			return err
		}

		if err := r.InsertOneUser(tx, ctx, &User{ID: f.Sender}); err != nil {
			logx.Error("Insert user(sender) into room err :", err)
			return err
		}

		if err := r.InsertOneUser(tx, ctx, &User{ID: f.Receiver}); err != nil {
			logx.Error("Insert user(receiver) into room err :", err)
			return err
		}

		return nil
	})
}

func (f *FriendNotification) Cancel(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Model(&f).Update("state", false).Error
}
func (f *FriendNotification) Decline(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Model(&f).Update("state", false).Error
}

func (f *FriendNotification) GetNotifications(db *gorm.DB, ctx context.Context) ([]*FriendNotification, error) {
	var resp []*FriendNotification
	if err := db.WithContext(ctx).Debug().Model(&f).Preload("SenderInfo").Where("Receiver = ? AND State = ?", f.Receiver, f.State).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}
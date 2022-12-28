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
	State    uint //there are 3 start . 0 represent decline and cancel ,1 represent send a request and 2 represent accepted
	DefaultModel

	SenderInfo   User `gorm:"foreignKey:Sender;references:ID"`
	ReceiverInfo User `gorm:"foreignKey:Receiver;references:ID"`
}

func (f *FriendNotification) TableName() string {
	return "friend_notifications"
}

func (f *FriendNotification) InsertOne(db *gorm.DB, ctx context.Context, receiver *User) error {
	return db.WithContext(ctx).Debug().Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Debug().Create(&f).Error; err != nil {
			return err
		}

		//TODO: add 1 friend notification count for the receiver
		receiver.FriendNotificationCount = receiver.FriendNotificationCount + 1
		if err := tx.WithContext(ctx).Debug().Model(&receiver).UpdateColumn("FriendNotificationCount", receiver.FriendNotificationCount).Error; err != nil {
			return err
		}
		return nil
	})
}

func (f *FriendNotification) FineOneByID(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Where("ID = ? AND State = ?", f.ID, f.State).First(&f).Error
}

func (f *FriendNotification) FineOneBySenderAndReceiver(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Where("(Sender = ? AND Receiver = ? AND State = ?) OR (Sender = ? AND Receiver = ? AND State = ? )", f.Sender, f.Receiver, f.State, f.Receiver, f.Sender, f.State).First(&f).Error
}

func (f *FriendNotification) Accept(db *gorm.DB, ctx context.Context) error {
	//calling FriendTemp model

	return db.WithContext(ctx).Debug().Transaction(func(tx *gorm.DB) error {
		/*
			TODO: Update Notification State
			TODO: Add FriendTemp to Friendship of both of them
		*/

		if err := tx.WithContext(ctx).Debug().Model(&f).Update("State", 2).Error; err != nil {
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
		//TODO: CreateOne both user with new roomID!
		r := &Room{OwnerRef: f.Sender}
		if err := r.InsertOne(tx, ctx); err != nil {
			logx.Error("create room error : ", err)
			return err
		}

		if err := r.InsertOneUser(tx, ctx, &User{ID: f.Sender}); err != nil {
			logx.Error("CreateOne user(sender) into room err :", err)
			return err
		}

		if err := r.InsertOneUser(tx, ctx, &User{ID: f.Receiver}); err != nil {
			logx.Error("CreateOne user(receiver) into room err :", err)
			return err
		}

		//TODO: update user friend notification count
		receiverInfo := &User{
			ID: f.Sender,
		}

		if err := tx.WithContext(ctx).Debug().First(&receiverInfo).Error; err != nil {
			return err
		}

		receiverInfo.FriendNotificationCount = receiverInfo.FriendNotificationCount + 1
		return receiverInfo.UpdateFriendNotification(tx, ctx)

	})
}

func (f *FriendNotification) Cancel(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Model(&f).Update("state", 0).Error
}
func (f *FriendNotification) Decline(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Model(&f).Update("state", 0).Error
}

func (f *FriendNotification) GetNotifications(db *gorm.DB, ctx context.Context, limit, pageOffset int) ([]*FriendNotification, int64, error) {
	var resp []*FriendNotification

	var count int64 = 0
	if err := db.WithContext(ctx).Debug().Model(&f).
		Preload("SenderInfo").
		Where("Receiver = ? AND State between ? and ?", f.Receiver, 1, 2).
		Order("created_at desc").
		Count(&count).
		Offset(pageOffset).
		Limit(limit).
		Find(&resp).Error; err != nil {
		return nil, 0, err

	}
	return resp, count, nil
}

package models

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        uint   `gorm:"primaryKey"`
	MessageID string //From Client
	RoomID    uint
	Sender    uint
	Content   string
	SentTime  time.Time `gorm:"type:timestamp"` //FromClient
	DefaultModel

	RoomInfo Room `gorm:"foreignKey:RoomID;references:ID"`
	SendUser User `gorm:"foreignKey:Sender;references:ID"`
}

/*
A -> RoomID x
x ,A,"message1",time
x ,B,"message2",time
x ,C,"message3",time

sorting by time
*/

func (m *Message) TableName() string {
	return "messages"
}

func (m *Message) InsertOne(db *gorm.DB, ctx context.Context) error {
	return db.WithContext(ctx).Debug().Create(&m).Error
}

func (m *Message) GetRoomMessages(db *gorm.DB, ctx context.Context, lastID uint, limit, pageOffset int) ([]*Message, int64, error) {
	var record []*Message
	var count int64 = 0
	//if we use offset directly,it'll cause a bug
	//example: we now have message as [7,8,9,10] ,and the offset is 4. Now, a user send a new message which message id is 11
	//then we got some older message by offset ,and we are expecting the result is [3,4,5,6] . But we got a new message ,so the result will be [4,5,6,7]
	//the message id which is 7 is duplicated in client side : db:[1,2,3,4,5,6,7,8,9,10,(11) new one]
	if err := db.WithContext(ctx).Debug().
		Model(&m).
		Where("id < ? AND room_id = ?", lastID, m.RoomID). //using last ID for offset -> smaller id
		Preload("SendUser").
		Order("sent_time desc").
		Count(&count).
		Limit(limit).
		Find(&record).Error; err != nil {
		return nil, 0, err
	}
	return record, count, nil
}

func (m *Message) GetRoomLatestMessages(db *gorm.DB, ctx context.Context) ([]Message, error) {
	var record []Message
	if err := db.WithContext(ctx).Debug().
		Model(&m).
		Preload("SendUser").
		Where("room_id = ?", m.RoomID).
		Order("sent_time desc").Limit(10).
		Find(&record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (m *Message) CountMessage(db *gorm.DB, ctx context.Context) (int64, error) {
	var count int64 = 0
	if err := db.WithContext(ctx).Debug().Model(&m).Where("room_id = ?", m.RoomID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil

}

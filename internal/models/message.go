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

func (m *Message) GetRoomMessages(db *gorm.DB, ctx context.Context, limit, pageOffset int) ([]*Message, int64, error) {
	var record []*Message
	var count int64 = 0
	if err := db.WithContext(ctx).Debug().
		Model(&m).
		Where("room_id = ?", m.RoomID).
		Preload("SendUser").
		Order("sent_time desc").
		Count(&count).
		Offset(pageOffset).
		Limit(limit).
		Find(&record).Error; err != nil {
		return nil, 0, err
	}
	return record, count, nil
}

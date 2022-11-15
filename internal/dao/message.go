package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"time"
)

func (d *DAO) InsertOneMessage(ctx context.Context, roomID, userId uint, message, messageID string, sentTime int64) error {

	msg := &models.Message{
		RoomID:    roomID,
		Sender:    userId,
		Content:   message,
		MessageID: messageID,
		SentTime:  time.Unix(sentTime, 0),
	}

	return msg.InsertOne(d.engine, ctx)
}

func (d *DAO) GetRoomMessage(ctx context.Context, roomID, lastID uint, limit, pageOffset int) ([]*models.Message, int64, error) {
	msg := &models.Message{
		RoomID: roomID,
	}

	return msg.GetRoomMessages(d.engine, ctx, lastID, limit, pageOffset)
}

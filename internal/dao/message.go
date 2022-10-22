package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) InsertOneMessage(ctx context.Context, roomID, userId uint, message string) error {
	msg := &models.Message{
		RoomID: roomID,
		Sender: userId,
		Data:   message,
	}

	return msg.InsertOne(d.engine, ctx)
}

func (d *DAO) GetRoomMessage(ctx context.Context, roomID uint) ([]*models.Message, error) {
	msg := &models.Message{
		RoomID: roomID,
	}

	return msg.GetRoomMessages(d.engine, ctx)
}

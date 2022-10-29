package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) ExistInTheRoom(ctx context.Context, userID, roomID uint) error {
	ur := &models.UsersRooms{
		RoomID: roomID,
		UserID: userID,
	}
	return ur.FindOne(d.engine, ctx)
}

func (d *DAO) GetRoomUsers(ctx context.Context, roomID uint) ([]uint, error) {
	ur := &models.UsersRooms{
		RoomID: roomID,
	}
	return ur.GetRoomUsers(d.engine, ctx)
}

func (d *DAO) UpdateRoomActiveState(ctx context.Context, roomID uint, userID uint, state bool) error {
	ur := &models.UsersRooms{
		RoomID:   roomID,
		UserID:   userID,
		IsActive: state,
	}

	return ur.SetActiveState(d.engine, ctx)
}

//func (d *DAO) GetUserRoomState(ctx context.Context, roomID uint, userID uint) (bool, error) {
//	ur := &models.UsersRooms{
//		RoomID: roomID,
//		UserID: userID,
//	}
//
//	return ur.GetUserRoomState(d.engine, ctx)
//}

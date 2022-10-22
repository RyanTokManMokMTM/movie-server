package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) InsertOneRoom(ctx context.Context, name, info string, userID uint) (*models.Room, error) {
	r := &models.Room{
		Name:     name,
		Info:     info,
		OwnerRef: userID,
	}

	if err := r.InsertOne(d.engine, ctx); err != nil {
		return nil, err
	}

	return r, nil
}

func (d *DAO) RemoveOneRoom(ctx context.Context, roomID uint) error {
	r := &models.Room{
		ID: roomID,
	}

	return r.RemoveOne(d.engine, ctx)
}

func (d *DAO) FindOneOwnerRoom(ctx context.Context, roomID, userID uint) (*models.Room, error) {
	r := &models.Room{
		ID:       roomID,
		OwnerRef: userID,
	}

	if err := r.FindOne(d.engine, ctx); err != nil {
		return nil, err
	}

	return r, nil
}

func (d *DAO) FindOneByRoomID(ctx context.Context, roomID uint) (*models.Room, error) {
	r := &models.Room{
		ID: roomID,
	}

	if err := r.FindOne(d.engine, ctx); err != nil {
		return nil, err
	}

	return r, nil
}

func (d *DAO) JoinOneRoom(ctx context.Context, roomID uint, u *models.User) error {
	r := models.Room{
		ID: roomID,
	}
	return r.InsertOneUser(d.engine, ctx, u)
}

func (d *DAO) LeaveOneRoom(ctx context.Context, roomID uint, u *models.User) error {
	r := models.Room{
		ID: roomID,
	}
	return r.RemoveOneUser(d.engine, ctx, u)
}

func (d *DAO) FindRoomMembers(ctx context.Context, roomID uint) ([]*models.User, error) {
	r := models.Room{
		ID: roomID,
	}
	return r.FindRoomMembers(d.engine, ctx)
}

func (d *DAO) FindOneRoomMember(ctx context.Context, roomID, userID uint) (*models.User, error) {
	r := models.Room{
		ID: roomID,
	}
	return r.FindOneRoomMember(d.engine, ctx, userID)
}

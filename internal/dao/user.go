package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if err := user.Insert(ctx, d.engine); err != nil {
		return nil, err
	}
	return user, nil
}

func (d *DAO) FindUserByID(ctx context.Context, userID uint) (*models.User, error) {
	user := &models.User{
		Id: userID,
	}
	if err := user.FindOneByID(ctx, d.engine); err != nil {
		return nil, err
	}

	return user, nil
}

func (d *DAO) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{
		Email: email,
	}
	if err := user.FindOneByEmail(ctx, d.engine); err != nil {
		return nil, err
	}
	return user, nil
}

func (d *DAO) UpdateUser(ctx context.Context, user *models.User) error {
	if err := user.UpdateInfo(ctx, d.engine); err != nil {
		return err
	}
	return nil
}

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

func (d *DAO) AddLikedMovie(ctx context.Context, movieID, userID uint) error {
	user := &models.User{Id: userID}
	movie := &models.MovieInfo{MovieId: movieID}

	return user.AddLikedMovie(ctx, d.engine, movie)
}

func (d *DAO) RemoveLikedMovie(ctx context.Context, movieID, userID uint) error {
	user := &models.User{Id: userID}
	movie := &models.MovieInfo{MovieId: movieID}

	return user.RemoveLikedMovie(ctx, d.engine, movie)
}

func (d *DAO) GetUserLikedMovies(ctx context.Context, userID uint) (*models.User, error) {
	user := &models.User{Id: userID}

	err := user.GetUserLikedMovies(ctx, d.engine)
	if err != nil {
		return nil, err
	}
	return user, err
}

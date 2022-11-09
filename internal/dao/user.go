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
		ID: userID,
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

func (d *DAO) GetUserLikedMovies(ctx context.Context, userID uint, limit, pageOffset int) (*models.User, int64, error) {
	user := &models.User{ID: userID}

	err, count := user.GetUserLikedMovies(ctx, d.engine, limit, pageOffset)
	if err != nil {
		return nil, 0, err
	}
	return user, count, nil
}

func (d *DAO) CreateLikedMovie(ctx context.Context, movieID, userID uint) error {
	user := &models.User{ID: userID}
	movie := &models.MovieInfo{Id: movieID}

	return user.CreateLikedMovie(ctx, d.engine, movie)
}

//func (d *DAO) FindUserFollowingList(ctx context.Context, userId uint) ([]*models.User, error) {
//	user := &models.User{
//		ID: userId,
//	}
//
//	return user.GetFollowingList(ctx, d.engine)
//}
//
//func (d *DAO) FindUserFollowedList(ctx context.Context, userId uint) ([]*models.User, error) {
//	f := &models.User{
//		ID: userId,
//	}
//
//	return f.GetFollowedList(ctx, d.engine)
//}

func (d *DAO) CreatePostLiked(ctx context.Context, userId uint, postId *models.Post) error {
	u := &models.User{ID: userId}
	return u.CreateUserPostLiked(ctx, d.engine, postId)
}

func (d *DAO) CreateCommentLiked(ctx context.Context, userId uint, comment *models.Comment) error {
	u := &models.User{ID: userId}

	return u.CreateUserCommentLiked(ctx, d.engine, comment)
}

func (d *DAO) GetUserRooms(ctx context.Context, userID uint) ([]*models.Room, error) {
	u := &models.User{
		ID: userID,
	}
	return u.GetUserRooms(ctx, d.engine)
}

func (d *DAO) GetUserActiveRooms(ctx context.Context, userID uint) ([]models.Room, error) {
	u := &models.User{
		ID: userID,
	}
	return u.GetUserActiveRooms(ctx, d.engine)
}

func (d *DAO) GetUserRoomsWithMembers(ctx context.Context, userID uint) (*models.User, error) {
	u := &models.User{
		ID: userID,
	}
	if err := u.GetUserRoomsWithMembers(ctx, d.engine); err != nil {
		return nil, err
	}
	return u, nil
}

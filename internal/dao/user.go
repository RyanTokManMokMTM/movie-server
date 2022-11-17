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

func (d *DAO) RemoveLikedMovie(ctx context.Context, movieID, userID uint) error {
	user := &models.User{ID: userID}
	movie := &models.MovieInfo{Id: movieID}

	return user.RemoveLikedMovie(ctx, d.engine, movie)
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

//
//func (d *DAO) CreateCommentLiked(ctx context.Context, userId uint, comment *models.Comment) error {
//	u := &models.User{ID: userId}
//
//	return u.CreateUserCommentLiked(ctx, d.engine, comment)
//}

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

func (d *DAO) InsertOneCommentLike(ctx context.Context, userID uint, commentID, count uint) error {
	u := &models.User{
		ID: userID,
	}

	return u.InsertOneCommentLikes(ctx, d.engine, commentID, count)
}

func (d *DAO) RemoveOneCommentLike(ctx context.Context, userID uint, commentID, count uint) error {
	u := &models.User{
		ID: userID,
	}

	return u.RemoveOneCommentLikes(ctx, d.engine, commentID, count)
}

//FriendNotification

func (d *DAO) UpdateFriendNotification(ctx context.Context, u *models.User, count uint) error {
	u.FriendNotificationCount = u.FriendNotificationCount + count
	return u.UpdateFriendNotification(d.engine, ctx)
}

func (d *DAO) ResetFriendNotification(ctx context.Context, u *models.User) error {
	u.FriendNotificationCount = 0
	return u.UpdateFriendNotification(d.engine, ctx)
}

//LikesNotification

func (d *DAO) UpdateLikesNotification(ctx context.Context, u *models.User, count uint) error {
	u.LikeNotificationCount = u.LikeNotificationCount + count
	return u.UpdateFriendNotification(d.engine, ctx)
}

func (d *DAO) ResetLikesNotification(ctx context.Context, u *models.User) error {
	u.LikeNotificationCount = 0
	return u.UpdateLikesNotification(d.engine, ctx)
}

//CommentNotification

func (d *DAO) UpdateCommentNotification(ctx context.Context, u *models.User, count uint) error {
	u.CommentNotificationCount = u.CommentNotificationCount + count
	return u.UpdateFriendNotification(d.engine, ctx)
}

func (d *DAO) ResetCommentNotification(ctx context.Context, u *models.User) error {
	u.CommentNotificationCount = 0
	return u.UpdateCommentNotification(d.engine, ctx)
}

package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

func (d *DAO) CreateNewFriend(ctx context.Context, userId, friendId uint) error {
	f := &models.Friend{
		UserId:   userId,
		FriendID: friendId,
	}

	return f.AddNewsFriend(ctx, d.engine)
}

func (d *DAO) UpdateFriendState(ctx context.Context, f *models.Friend) error {
	return f.UpdateFriendState(ctx, d.engine)
}

func (d *DAO) FindOneFriend(ctx context.Context, userId, friendId uint) (*models.Friend, error) {
	f := &models.Friend{
		UserId:   userId,
		FriendID: friendId,
	}
	//if it returns not found -> not friend?
	if err := f.FindOneUserFromFriendList(ctx, d.engine); err != nil {
		return nil, err
	}
	return f, nil
}

func (d *DAO) CountFollowingUser(ctx context.Context, userId uint) (int64, error) {
	f := &models.Friend{
		UserId: userId,
	}
	return f.CountFollowingUser(ctx, d.engine)
}

func (d *DAO) CountFollowedUser(ctx context.Context, userId uint) (int64, error) {
	f := &models.Friend{
		UserId: userId,
	}
	return f.CountFollowedUser(ctx, d.engine)
}

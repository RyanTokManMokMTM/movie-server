package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
)

//
//func (d *DAO) CreateNewFriend(ctx context.Context, userId, friendId uint) error {
//	f := &models.FriendTemp{
//		UserId:   userId,
//		FriendID: friendId,
//	}
//
//	return f.AddNewsFriend(ctx, d.engine)
//}
//
//func (d *DAO) UpdateFriendState(ctx context.Context, f *models.FriendTemp) error {
//	return f.UpdateFriendState(ctx, d.engine)
//}
//
//func (d *DAO) FindOneFriend(ctx context.Context, userId, friendId uint) (*models.FriendTemp, error) {
//	f := &models.FriendTemp{
//		UserId:   userId,
//		FriendID: friendId,
//	}
//	//if it returns not found -> not friend?
//	if err := f.FindOneUserFromFriendList(ctx, d.engine); err != nil {
//		return nil, err
//	}
//	return f, nil
//}
//

//func (d *DAO) CountFollowingUser(ctx context.Context, userId uint) (int64, error) {
//	f := &models.Friend{
//		UserID: userId,
//	}
//	return f.CountFollowingUser(ctx, d.engine)
//}
//
//func (d *DAO) CountFollowedUser(ctx context.Context, userId uint) (int64, error) {
//	f := &models.FriendTemp{
//		UserId: userId,
//	}
//	return f.CountFollowedUser(ctx, d.engine)
//}

func (d *DAO) CountFriends(ctx context.Context, friendID uint) int64 {
	f := &models.Friend{
		ID: friendID,
	}
	return f.CountFriend(d.engine, ctx)
}

func (d *DAO) GetFriendsList(ctx context.Context, friendID uint) ([]*models.User, error) {
	f := &models.Friend{
		ID: friendID,
	}
	return f.GetFriendsList(d.engine, ctx)
}

func (d *DAO) GetUserFriendRecord(ctx context.Context, userId uint) (*models.Friend, error) {
	f := &models.Friend{
		UserID: userId,
	}
	if err := f.GetUserFriend(d.engine, ctx); err != nil {
		return nil, err
	}

	return f, nil
}

func (d *DAO) InsertOneFriendNotification(ctx context.Context, sender, receiver uint) error {
	fr := &models.FriendNotification{
		Sender:   sender,
		Receiver: receiver,
		State:    true,
	}

	return fr.InsertOne(d.engine, ctx)
}
func (d *DAO) FindOneFriendNotification(ctx context.Context, sender, receiver uint) (*models.FriendNotification, error) {
	fr := &models.FriendNotification{
		Sender:   sender,
		Receiver: receiver,
		State:    true,
	}
	err := fr.FineOne(d.engine, ctx)
	if err != nil {
		return nil, err
	}
	return fr, nil
}

func (d *DAO) FindOneFriendNotificationByID(ctx context.Context, requestID uint) (*models.FriendNotification, error) {
	fr := &models.FriendNotification{
		ID:    requestID,
		State: true,
	}

	err := fr.FineOne(d.engine, ctx)
	if err != nil {
		return nil, err
	}
	return fr, nil
}

func (d *DAO) AcceptFriendNotification(ctx context.Context, fr *models.FriendNotification) error {
	return fr.Accept(d.engine, ctx)
}
func (d *DAO) CancelFriendNotification(ctx context.Context, requestID uint) error {
	f := &models.FriendNotification{
		ID: requestID,
	}

	return f.Cancel(d.engine, ctx)
}

func (d *DAO) DeclineFriendNotification(ctx context.Context, requestID uint) error {
	f := &models.FriendNotification{
		ID: requestID,
	}

	return f.Decline(d.engine, ctx)
}

func (d *DAO) HasFriendShip(ctx context.Context, userID, friendID uint) (*models.User, error) {
	f := &models.Friend{
		UserID: userID,
	}

	return f.FindOneFriend(d.engine, ctx, friendID)
}

func (d *DAO) InsertOneFriendInstance(ctx context.Context, userID uint) error {
	f := &models.Friend{
		UserID: userID,
	}

	return f.InsertOne(d.engine, ctx)
}

func (d *DAO) RemoveFriend(ctx context.Context, userID, friendID uint) error {
	f := &models.Friend{}

	return f.RemoveOne(d.engine, ctx, userID, friendID)
}

func (d *DAO) GetFriendRequest(ctx context.Context, userID uint) ([]*models.FriendNotification, error) {
	notification := &models.FriendNotification{
		Receiver: userID,
		State:    true,
	}

	return notification.GetNotifications(d.engine, ctx)
}

func (d *DAO) IsFriend(ctx context.Context, userID, friendID uint) (bool, error) {
	f := &models.Friend{
		UserID: userID,
	}

	return f.IsFriend(d.engine, ctx, friendID)
}

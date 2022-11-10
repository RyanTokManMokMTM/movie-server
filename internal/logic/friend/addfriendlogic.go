package friend

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/internal/logic/serverWs"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFriendLogic) AddFriend(req *types.AddFriendReq) (resp *types.AddFriendResp, err error) {
	// todo: add your logic here and delete this line

	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	if userId == 0 {
		return nil, fmt.Errorf("user_id is missing")
	}
	//TODO: Check - FriendID can't be it self or we can?????
	if req.UserID == userId {
		return nil, fmt.Errorf("you can't add your own as friend")
	}

	//TODO: whether Your is exist
	u, err := l.svcCtx.DAO.FindUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not exist")
		}
		return nil, err
	}

	//TODO: Check that user is exist?
	receiver, err := l.svcCtx.DAO.FindUserByID(l.ctx, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("friend is not exist")
		}
		return nil, err
	}

	//TODO: Is Friend already?
	isFriend, err := l.svcCtx.DAO.IsFriend(l.ctx, userId, req.UserID)
	if err != nil {
		return nil, err
	}

	if isFriend {
		return &types.AddFriendResp{
			SenderID:  0,
			RequestID: 0,
		}, nil
	}
	//TODO:Add to the notification if it hasn't sent a request.
	_, err = l.svcCtx.DAO.FindOneFriendNotification(l.ctx, userId, req.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//TODO: insert 1 notification and add the friend count
		id, err := l.svcCtx.DAO.InsertOneFriendNotification(l.ctx, userId, receiver)
		if err != nil {
			return nil, err
		}
		go func() {
			//Send the notification via websocket
			_ = serverWs.SendNotificationToUserWithUserInfo(req.UserID, u, fmt.Sprintf("%s傳送了好友邀請給您", u.Name), serverWs.FRIEND_NOTIFICATION)
			//_ = serverWs.SendNotificationToUser(u.ID, req.UserID, fmt.Sprintf("%s傳送了好友邀請給您", u.Name))
		}()

		return &types.AddFriendResp{
			SenderID:  userId,
			RequestID: id,
		}, nil
	}

	return nil, fmt.Errorf("friend request had been sent")
}

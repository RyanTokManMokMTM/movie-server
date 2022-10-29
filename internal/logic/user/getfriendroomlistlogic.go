package user

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendRoomListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendRoomListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendRoomListLogic {
	return &GetFriendRoomListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendRoomListLogic) GetFriendRoomList(req *types.GetUserFriendRoomListReq) (resp *types.GetUserFriendRoomListResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	list, err := l.svcCtx.DAO.GetFriendRoomList(l.ctx, userID)
	if err != nil {
		return nil, err
	}

	var friendList []types.FriendRoomInfo
	for _, v := range list.Rooms {
		//logx.Infof("%+v", v.Users)
		friendList = append(friendList, types.FriendRoomInfo{
			RoomID: v.ID,
			Info: types.UserInfo{
				ID:     v.Users[0].ID,
				Name:   v.Users[0].Name,
				Avatar: v.Users[0].Avatar,
			},
		})
	}

	return &types.GetUserFriendRoomListResp{
		Lists: friendList,
	}, nil
}

package UserMovieList

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"
	"strconv"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	// todo: add your logic here and delete this line
	userID := fmt.Sprintf("%v", l.ctx.Value("userID"))
	id, _ := strconv.Atoi(userID)
	_, err = l.svcCtx.User.FindOne(l.ctx, int64(id))
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	//Return all user movie list by user_id
	res, err := l.svcCtx.List.FindAllByUserID(l.ctx, req.Id)
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	var ListInfos []*types.ListInfo
	for _, v := range res {
		ListInfos = append(ListInfos, &types.ListInfo{
			Id:         v.ListId,
			ListTitle:  v.ListTitle,
			UserId:     v.UserId,
			CreateTime: v.CreateTime.Unix(),
			UpdateTime: v.UpdateTime.Unix(),
		})
	}

	return &types.UserListResp{
		Lists: ListInfos,
	}, nil
}

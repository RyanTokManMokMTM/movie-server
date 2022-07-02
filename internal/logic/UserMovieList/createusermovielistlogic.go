package UserMovieList

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"
	"github.com/ryantokmanmokmtm/movie-server/model/list"
	"strconv"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserMovieListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserMovieListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserMovieListLogic {
	return &CreateUserMovieListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserMovieListLogic) CreateUserMovieList(req *types.CreateNewUserListReq) (resp *types.CreateNewUserListResp, err error) {
	// todo: add your logic here and delete this line
	userID := fmt.Sprintf("%v", l.ctx.Value("userID"))
	id, _ := strconv.Atoi(userID)
	_, err = l.svcCtx.User.FindOne(l.ctx, int64(id))
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	newModel := list.Lists{
		ListTitle: req.Title,
		UserId:    int64(id),
	}
	res, err := l.svcCtx.List.Insert(l.ctx, &newModel)

	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	newModel.ListId, err = res.LastInsertId()
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	return &types.CreateNewUserListResp{
		Id:        newModel.ListId,
		UserId:    newModel.UserId,
		ListTitle: newModel.ListTitle,
	}, nil
}

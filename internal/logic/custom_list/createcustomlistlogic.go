package custom_list

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/model/list"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCustomListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCustomListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCustomListLogic {
	return &CreateCustomListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCustomListLogic) CreateCustomList(req *types.CreateCustomListReq) (resp *types.CreateCustomListResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	//find user
	user, err := l.svcCtx.User.FindOne(l.ctx, userID)
	if err != nil && err != sqlx.ErrNotFound {
		return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("CreateCustomList - user db err:%v, userID:%v", err, userID))
	}

	if user == nil {
		return nil, errors.Wrap(errx.NewErrCode(errx.USER_NOT_EXIST), fmt.Sprintf("CreateCustomList - user db USER NOT FOUND err: %v, userID: %v", err, userID))
	}
	//Do we need to check title exits????

	newList := list.Lists{
		UserId:    userID,
		ListTitle: req.Title,
	}
	sqlRes, err := l.svcCtx.List.Insert(l.ctx, &newList)
	if err != nil {
		return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("CreateCustomList - List db err:%v, req:%+v", err, req))
	}

	newList.ListId, err = sqlRes.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(errx.NewErrCode(errx.DB_AFFECTED_ZERO_ERROR), fmt.Sprintf("CreateCustomList - List db INSERT.LastInsertId err: %v, req: %+v", err, req))
	}

	return &types.CreateCustomListResp{
		ID:       newList.ListId,
		Title:    newList.ListTitle,
		UpdateOn: newList.UpdateTime.Unix(),
	}, nil
}

package custom_list

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
	userID := fmt.Sprintf("%v", l.ctx.Value("userID"))
	id, _ := strconv.Atoi(userID)
	_, err = l.svcCtx.User.FindOne(l.ctx, int64(id))
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	//Do we need to check title exits????

	newList := list.Lists{
		UserId:    int64(id),
		ListTitle: req.Title,
	}
	sqlRes, err := l.svcCtx.List.Insert(l.ctx, &newList)
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	newList.ListId, err = sqlRes.LastInsertId()
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	return &types.CreateCustomListResp{
		ID:       newList.ListId,
		Title:    newList.ListTitle,
		UpdateOn: newList.UpdateTime.Unix(),
	}, nil
}

package custom_list

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUserListLogic {
	return &GetAllUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUserListLogic) GetAllUserList(req *types.AllCustomListReq) (resp *types.AllCustomListResp, err error) {
	// todo: add your logic here and delete this line
	logx.Info("Get All User List")
	logx.Info(req.ID)
	listInfos, err := l.svcCtx.List.FindAllByUserID(l.ctx, req.ID)
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	var listResp []types.ListInfo
	for _, v := range listInfos {
		listResp = append(listResp, types.ListInfo{
			ID:       v.ListId,
			Title:    v.ListTitle,
			UpdateOn: v.UpdateTime.Unix(),
		})
	}

	return &types.AllCustomListResp{
		Lists: listResp,
	}, nil
	return
}

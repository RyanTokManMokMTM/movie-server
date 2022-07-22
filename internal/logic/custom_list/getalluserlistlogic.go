package custom_list

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
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

	lists, err := l.svcCtx.DAO.GetUserLists(l.ctx, req.ID)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	var userLists []types.ListInfo
	for _, v := range lists {
		userLists = append(userLists, types.ListInfo{
			ID:       v.ListId,
			Title:    v.ListTitle,
			UpdateOn: v.UpdatedAt.Unix(),
		})
	}
	return &types.AllCustomListResp{
		Lists: userLists,
	}, nil
}

//
//func (l *GetAllUserListLogic) GetAllUserList(req *types.AllCustomListReq) (resp *types.AllCustomListResp, err error) {
//	// todo: add your logic here and delete this line
//	//logx.Info("Get All User List")
//	//logx.Info(req.ID)
//	//listInfos, err := l.svcCtx.List.FindAllByUserID(l.ctx, req.ID)
//	//if err != nil && err != sqlx.ErrNotFound {
//	//	//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("GetAllUserList - list db find by user id err: %v, userID: %v", err, req.ID))
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//	//
//	//var listResp []types.ListInfo
//	//for _, v := range listInfos {
//	//	listResp = append(listResp, types.ListInfo{
//	//		ID:       v.ListId,
//	//		Title:    v.ListTitle,
//	//		UpdateOn: v.UpdateTime.Unix(),
//	//	})
//	//}
//	//
//	//return &types.AllCustomListResp{
//	//	Lists: listResp,
//	//}, nil
//	return
//}

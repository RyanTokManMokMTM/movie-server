package custom_list

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCustomListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomListLogic {
	return &UpdateCustomListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCustomListLogic) UpdateCustomList(req *types.UpdateCustomListReq) (resp *types.UpdateCustomListResp, err error) {
	// todo: add your logic here and delete this line

	logx.Info("Update Custom List")
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	list, err := l.svcCtx.DAO.FindOneList(l.ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.LIST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if req.Title != "" {
		list.ListTitle = req.Title
	}

	if req.Intro != "" {
		list.ListIntro = req.Intro
	}

	if err := l.svcCtx.DAO.UpdateList(l.ctx, list); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return
}

//
//func (l *UpdateCustomListLogic) UpdateCustomList(req *types.UpdateCustomListReq) (resp *types.UpdateCustomListResp, err error) {
//	// todo: add your logic here and delete this line
//	//userID := ctxtool.GetUserIDFromCTX(l.ctx)
//	//
//	////find user
//	//user, err := l.svcCtx.User.FindOne(l.ctx, userID)
//	//if err != nil && err != sqlx.ErrNotFound {
//	//	//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UpdateCustomList - user db err:%v, userID:%v", err, userID))
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//	//
//	//if user == nil {
//	//	//return nil, errors.Wrap(errx.NewErrCode(errx.USER_NOT_EXIST), fmt.Sprintf("UpdateCustomList - user db FINDgot NOT FOUND err: %v, userID: %v", err, userID))
//	//	return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
//	//}
//	//
//	//res, err := l.svcCtx.List.FindOne(l.ctx, req.ID)
//	//if err != nil && err != sqlx.ErrNotFound {
//	//	//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UpdateCustomList - list db FIND err: %v, ListID: %v", err, req.ID))
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//	//
//	//if res == nil {
//	//	return nil, errx.NewErrCode(errx.LIST_NOT_EXIST)
//	//}
//	//
//	////title is a required field
//	//res.ListTitle = req.Title
//	//
//	//err = l.svcCtx.List.Update(l.ctx, res)
//	//if err != nil {
//	//	//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UpdateCustomList - list db UPDATE  err: %v, req: %+v", err, req))
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//	//return &types.UpdateCustomListResp{}, nil
//	return
//}

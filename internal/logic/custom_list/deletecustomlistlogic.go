package custom_list

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteCustomListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCustomListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCustomListLogic {
	return &DeleteCustomListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCustomListLogic) DeleteCustomList(req *types.DeleteCustomListReq) (resp *types.DeleteCustomListResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if err := l.svcCtx.DAO.DeleteList(l.ctx, req.ID, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.LIST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return
}

//func (l *DeleteCustomListLogic) DeleteCustomList(req *types.DeleteCustomListReq) (resp *types.DeleteCustomListResp, err error) {
//	// todo: add your logic here and delete this line
//	//userID := ctxtool.GetUserIDFromCTX(l.ctx)
//	//
//	////find user
//	//user, err := l.svcCtx.User.FindOne(l.ctx, userID)
//	//if err != nil && err != sqlx.ErrNotFound {
//	//	//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("DeleteCustomList - user db err:%v, userID:%v", err, userID))
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//	//
//	//if user == nil {
//	//	//return nil, errors.Wrap(errx.NewErrCode(errx.USER_NOT_EXIST), fmt.Sprintf("DeleteCustomList - user db USER NOT FOUND err: %v, userID: %v", err, userID))
//	//	return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
//	//}
//	//
//	//err = l.svcCtx.List.Delete(l.ctx, req.ID)
//	//if err != nil {
//	//	//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("DeleteCustomList - List db delete err:%v, req:%v", err, req))
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//	//
//	//return &types.DeleteCustomListResp{}, nil
//	return
//}

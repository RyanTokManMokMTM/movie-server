package posts

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

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.DeletePostReq) (resp *types.DeletePostResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	//find post
	err = l.svcCtx.DAO.DeletePost(l.ctx, req.PostID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return
}

//func (l *DeletePostLogic) DeletePost(req *types.DeletePostReq) (resp *types.DeletePostResp, err error) {
//	// todo: add your logic here and delete this line
//	//userID := ctxtool.GetUserIDFromCTX(l.ctx)
//	//user, err := l.svcCtx.User.FindOne(l.ctx, userID)
//	//if err != nil && err != sqlx.ErrNotFound {
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//	//
//	//if user == nil {
//	//	return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
//	//}
//	//
//	//err = l.svcCtx.PostModel.Delete(l.ctx, req.PostID)
//	//if err != nil {
//	//	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	//}
//
//	return
//}

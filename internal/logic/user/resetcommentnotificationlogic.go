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

type ResetCommentNotificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetCommentNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetCommentNotificationLogic {
	return &ResetCommentNotificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetCommentNotificationLogic) ResetCommentNotification(req *types.CommentNotificationReq) (resp *types.CommentNotificationResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	user, err := l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if err := l.svcCtx.DAO.ResetCommentNotification(l.ctx, user); err != nil {
		return nil, err
	}

	return
}

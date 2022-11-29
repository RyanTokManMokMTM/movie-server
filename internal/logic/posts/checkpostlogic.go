package posts

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPostLogic {
	return &CheckPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckPostLogic) CheckPost(req *types.CheckPostReq) (resp *types.CheckPostResp, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DAO.FindOnePostInfo(l.ctx, req.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &types.CheckPostResp{
				IsExist: false,
			}, nil
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.CheckPostResp{
		IsExist: true,
	}, nil
}

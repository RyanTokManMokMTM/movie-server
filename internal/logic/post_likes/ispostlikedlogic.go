package post_likes

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsPostLikedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsPostLikedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsPostLikedLogic {
	return &IsPostLikedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsPostLikedLogic) IsPostLiked(req *types.IsPostLikedReq) (resp *types.IsPostLikedResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	//find that user
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.DAO.FindOnePostInfo(l.ctx, req.PostId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewErrCode(errx.POST_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	postLiked, err := l.svcCtx.DAO.FindOnePostLiked(l.ctx, userID, req.PostId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &types.IsPostLikedResp{
				IsLiked: false,
			}, nil
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if postLiked.State == 0 {
		return &types.IsPostLikedResp{
			IsLiked: false,
		}, nil
	}

	return &types.IsPostLikedResp{
		IsLiked: true,
	}, nil

}

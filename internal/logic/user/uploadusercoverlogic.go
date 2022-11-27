package user

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/common/uploadx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"net/http"
)

type UploadUserCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadUserCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadUserCoverLogic {
	return &UploadUserCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadUserCoverLogic) UploadUserCover(req *types.UploadImageReq) (resp *types.UploadImageResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	user, err := l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	fileName, err := uploadx.UploadFile(l.r, l.svcCtx.Config.MaxBytes, "uploadCover", l.svcCtx.Config.Path)
	if err != nil {
		return nil, errx.NewErrCode(errx.USER_UPLOAD_USER_AVATAR_FAILED)
	}

	//remove the original one?
	//_ = os.Remove(user.Cover)

	//update user avatar path
	user.Cover = fmt.Sprintf("/%s", fileName)

	if err := l.svcCtx.DAO.UpdateUser(l.ctx, user); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	return &types.UploadImageResp{
		Path: user.Cover,
	}, nil
}

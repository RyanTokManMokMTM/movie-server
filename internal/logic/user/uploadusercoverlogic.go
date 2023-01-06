package user

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/common/uploadx"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
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

	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		logx.Error(err)
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	err = l.r.ParseMultipartForm(l.svcCtx.Config.MaxBytes)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.USER_UPLOAD_USER_AVATAR_FAILED, err.Error())
	}

	file, handler, err := l.r.FormFile("uploadCover")
	if err != nil {
		return nil, errx.NewCommonMessage(errx.USER_UPLOAD_USER_AVATAR_FAILED, err.Error())
	}
	defer file.Close()

	fileName, err := uploadx.UploadFile(file, handler, l.svcCtx.Config.Path)
	if err != nil {
		logx.Error(err.Error())
		return nil, errx.NewErrCode(errx.USER_UPLOAD_USER_AVATAR_FAILED)
	}

	//update user avatar path
	cover := fmt.Sprintf("/%s", fileName)
	updatedCover := &models.User{Cover: cover}
	if err := l.svcCtx.DAO.UpdateUser(l.ctx, userID, updatedCover); err != nil {
		logx.Error(err)
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	return &types.UploadImageResp{
		Path: cover,
	}, nil
}

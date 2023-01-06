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

type UploadUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadUserAvatarLogic {
	return &UploadUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadUserAvatarLogic) UploadUserAvatar(req *types.UploadImageReq) (resp *types.UploadImageResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	//logx.Info("Path: ", l.svcCtx.Config.Path)
	//fileName, err := uploadx.UploadFile(l.r, l.svcCtx.Config.MaxBytes, "uploadAvatar", l.svcCtx.Config.Path)
	//if err != nil {
	//
	//	return nil, errx.NewCommonMessage(errx.USER_UPLOAD_USER_AVATAR_FAILED, err.Error())
	//}

	err = l.r.ParseMultipartForm(l.svcCtx.Config.MaxBytes)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.USER_UPLOAD_USER_AVATAR_FAILED, err.Error())
	}

	file, handler, err := l.r.FormFile("uploadAvatar")
	if err != nil {
		return nil, errx.NewCommonMessage(errx.USER_UPLOAD_USER_AVATAR_FAILED, err.Error())
	}
	defer file.Close()

	fileName, err := uploadx.UploadFile(file, handler, l.svcCtx.Config.Path)
	if err != nil {
		return nil, errx.NewCommonMessage(errx.USER_UPLOAD_USER_AVATAR_FAILED, err.Error())
	}

	//update user avatar path
	avatar := fmt.Sprintf("/%s", fileName)

	if err := l.svcCtx.DAO.UpdateUser(l.ctx, userID, &models.User{Avatar: avatar}); err != nil {
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}
	return &types.UploadImageResp{
		Path: avatar,
	}, nil
}

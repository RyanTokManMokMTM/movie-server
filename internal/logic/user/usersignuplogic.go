package user

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/crytox"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserSignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignUpLogic {
	return &UserSignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSignUpLogic) UserSignUp(req *types.UserSignUpReq) (resp *types.UserSignUpResp, err error) {
	// todo: add your logic here and delete this line

	//Check User

	user, err := l.svcCtx.DAO.FindUserByEmail(l.ctx, req.Email)
	if user != nil && user.Id != 0 {
		logx.Info(user)
		return nil, errx.NewErrCode(errx.EMAIL_HAS_BEEN_REGISTERED)
	}

	if err != nil && err == gorm.ErrRecordNotFound {
		newUser := &models.User{
			Name:     req.UserName,
			Email:    req.Email,
			Password: crytox.PasswordEncrypt(req.Password, l.svcCtx.Config.Salt),
			Avatar:   "/defaultAvatar.jpeg", //TODO: Upload User avatar
			Cover:    "/defaultCover.jpeg",  //TODO: Upload User Cover
		}

		user, err := l.svcCtx.DAO.CreateUser(l.ctx, newUser)
		if err != nil {
			return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
		}

		return &types.UserSignUpResp{
			ID:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}, nil
	}
	logx.Info(err.Error())
	return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
}

//func (l *UserSignUpLogic) UserSignUp(req *types.UserSignUpRequest) (resp *types.UserSignUpResponse, err error) {
//	// todo: add your logic here and delete this line
//	lowEmail := strings.ToLower(req.Email)
//	userInfo, err := l.svcCtx.User.FindOneByEmail(l.ctx, lowEmail)
//	if err != nil && err != sqlx.ErrNotFound {
//		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UserSignUp - user db err:%v, req:%+v", err, req))
//		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	}
//
//	if userInfo != nil {
//		return nil, errx.NewErrCode(errx.EMAIL_HAS_BEEN_REGISTERED)
//
//	}
//	//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UserSignUp - user db err: %v, req:%+v", err, req))
//	newUser := user.User{
//		Name:     req.UserName,
//		Email:    lowEmail,
//		Password: crytox.PasswordEncrypt(req.Password, l.svcCtx.Config.Salt),
//		//avatar for testing
//		Avatar: "https://upload.cc/i1/2022/07/03/MJIXkd.jpeg", //TODO: Upload User avatar
//		Cover:  "https://upload.cc/i1/2022/07/04/yQN7tU.jpeg", //TODO: Upload User Cover
//	}
//
//	res, err := l.svcCtx.User.Insert(l.ctx, &newUser)
//	if err != nil {
//		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UserSignUp - user db Insert err:%v, req:%+v", err, req))
//		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	}
//
//	newUser.Id, err = res.LastInsertId()
//	if err != nil {
//		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_AFFECTED_ZERO_ERROR), fmt.Sprintf("UserSignUp - user db Insert.LastInsertID err:%v, req:%+v", err, req))
//		return nil, errx.NewErrCode(errx.DB_AFFECTED_ZERO_ERROR)
//	}
//
//	return &types.UserSignUpResponse{
//		ID:    newUser.Id,
//		Name:  newUser.Name,
//		Email: lowEmail,
//	}, nil
//}

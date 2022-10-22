package user

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/crytox"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/common/jwtx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	// todo: add your logic here and delete this line

	user, err := l.svcCtx.DAO.FindUserByEmail(l.ctx, req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	hashedPassword := crytox.PasswordEncrypt(req.Password, l.svcCtx.Config.Salt)
	if hashedPassword != user.Password {
		return nil, errx.NewErrCode(errx.USER_PASSWORD_INCORRECT)
	}

	payload := map[string]interface{}{
		ctxtool.CTXJWTUserID: user.ID,
	}

	now := time.Now().Unix()
	exp := l.svcCtx.Config.Auth.AccessExpire
	key := l.svcCtx.Config.Auth.AccessSecret

	token, err := jwtx.GetToken(now, exp, key, payload)
	if err != nil {
		return nil, errx.NewErrCode(errx.TOKEN_GENERATE_ERROR)
	}

	return &types.UserLoginResp{
		Token:   token,
		Expired: now + exp,
	}, nil
}

//
//func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
//	// todo: add your logic here and delete this line
//
//	res, err := l.svcCtx.User.FindOneByEmail(l.ctx, req.Email)
//	if err != nil && err != sqlx.ErrNotFound {
//		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UserLogin - user db FindByEmail err:%v, Email:%v", err, req.Email))
//		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	}
//
//	if res == nil {
//		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UserLogin - user db FindByEmail NotFound err:%v, Email:%v", err, req.Email))
//		return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
//	}
//
//	hashedPassword := crytox.PasswordEncrypt(req.Password, l.svcCtx.Config.Salt)
//	if hashedPassword != res.Password {
//		//return nil, errors.Wrap(errx.NewErrCode(errx.USER_PASSWORD_INCORRECT), fmt.Sprintf("UserLogin - Password err:%v, Email:%v", err, req.Email))
//		return nil, errx.NewErrCode(errx.USER_PASSWORD_INCORRECT)
//	}
//
//	payload := map[string]interface{}{
//		ctxtool.CTXJWTUserID: res.ID,
//	}
//
//	now := time.Now().Unix()
//	exp := l.svcCtx.Config.Auth.AccessExpire
//	key := l.svcCtx.Config.Auth.AccessSecret
//
//	token, err := jwtx.GetToken(now, exp, key, payload)
//	if err != nil {
//		//return nil, errors.Wrap(errx.NewErrCode(errx.TOKEN_GENERATE_ERROR), fmt.Sprintf("UserLogin - Token Generate err:%v", err))
//		return nil, errx.NewErrCode(errx.TOKEN_GENERATE_ERROR)
//	}
//
//	return &types.UserLoginResp{
//		Token:   token,
//		Expired: now + exp,
//	}, nil
//}

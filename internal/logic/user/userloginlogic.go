package user

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/movie-server/common/crytox"
	"github.com/ryantokmanmokmtm/movie-server/common/jwtx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.User.FindOneByEmail(l.ctx, req.Email)
	if err == sqlx.ErrNotFound {
		return nil, errors.New("email not exist")
	} else if err != nil {
		return nil, errors.New(err.Error())
	}
	hashedPassword := crytox.PasswordEncrypt(req.Password, l.svcCtx.Config.Salt)
	if string(hashedPassword) != res.Password {
		return nil, errors.New("password incorrect")
	}

	payload := map[string]interface{}{
		"user_id": res.Id,
	}
	now := time.Now().Unix()
	exp := l.svcCtx.Config.Auth.AccessExpire
	key := l.svcCtx.Config.Auth.AccessSecret
	token, err := jwtx.GetToken(now, exp, key, payload)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &types.UserLoginResponse{
		Token:   token,
		Expired: now + exp,
	}, nil
}

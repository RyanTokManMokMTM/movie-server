package user

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/crytox"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/ryantokmanmokmtm/movie-server/model/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *UserSignUpLogic) UserSignUp(req *types.UserSignUpRequest) (resp *types.UserSignUpResponse, err error) {
	// todo: add your logic here and delete this line
	lowEmail := strings.ToLower(req.Email)
	userInfo, err := l.svcCtx.User.FindOneByEmail(l.ctx, lowEmail)
	if err != nil && err != sqlx.ErrNotFound {
		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UserSignUp - user db err:%v, req:%+v", err, req))
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	if userInfo != nil {
		return nil, errx.NewErrCode(errx.EMAIL_HAS_BEEN_REGISTERED)

	}
	//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UserSignUp - user db err: %v, req:%+v", err, req))
	newUser := user.Users{
		Name:     req.UserName,
		Email:    lowEmail,
		Password: crytox.PasswordEncrypt(req.Password, l.svcCtx.Config.Salt),
		//avatar for testing
		Avatar: "https://upload.cc/i1/2022/07/03/MJIXkd.jpeg", //TODO: Upload User avatar
		Cover:  "https://upload.cc/i1/2022/07/04/yQN7tU.jpeg", //TODO: Upload User Cover
	}

	res, err := l.svcCtx.User.Insert(l.ctx, &newUser)
	if err != nil {
		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UserSignUp - user db Insert err:%v, req:%+v", err, req))
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	newUser.Id, err = res.LastInsertId()
	if err != nil {
		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_AFFECTED_ZERO_ERROR), fmt.Sprintf("UserSignUp - user db Insert.LastInsertID err:%v, req:%+v", err, req))
		return nil, errx.NewErrCode(errx.DB_AFFECTED_ZERO_ERROR)
	}

	return &types.UserSignUpResponse{
		ID:    newUser.Id,
		Name:  newUser.Name,
		Email: lowEmail,
	}, nil
}

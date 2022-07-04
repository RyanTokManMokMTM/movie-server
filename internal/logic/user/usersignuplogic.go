package user

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/crytox"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"
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
	_, err = l.svcCtx.User.FindOneByEmail(l.ctx, lowEmail)
	if err == nil {
		return nil, errorx.NewDefaultCodeError("email exits")
	}
	if err == sqlx.ErrNotFound {
		newUser := user.Users{
			Name:     req.UserName,
			Email:    lowEmail,
			Password: crytox.PasswordEncrypt(req.Password, l.svcCtx.Config.Salt),
			//avatar for testing
			Avatar: "https://upload.cc/i1/2022/07/03/MJIXkd.jpeg",
			Cover:  "https://upload.cc/i1/2022/07/04/yQN7tU.jpeg",
		}

		res, err := l.svcCtx.User.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, errorx.NewDefaultCodeError(err.Error())
		}

		newUser.Id, err = res.LastInsertId()
		if err != nil {
			return nil, errorx.NewDefaultCodeError(err.Error())
		}

		return &types.UserSignUpResponse{
			ID:    newUser.Id,
			Name:  newUser.Name,
			Email: lowEmail,
		}, nil
	}
	return nil, errorx.NewDefaultCodeError(err.Error())
}

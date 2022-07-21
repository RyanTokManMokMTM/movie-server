package user

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserProfileLogic {
	return &UserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserProfileLogic) UserProfile(req *types.UserProfileReq) (resp *types.UserProfileResp, err error) {
	// todo: add your logic here and delete this line
	//logx.Infof("userId: %v", l.ctx.Value("userID"))
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	user, err := l.svcCtx.DAO.FindUserByID(l.ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
	}

	return &types.UserProfileResp{
		ID:     user.Id,
		Name:   user.Name,
		Email:  user.Email,
		Avatar: user.Avatar,
		Cover:  user.Cover,
	}, nil
}

//func (l *UserProfileLogic) UserProfile(req *types.UserProfileReq) (resp *types.UserProfileResp, err error) {
//	// todo: add your logic here and delete this line
//	//logx.Infof("userId: %v", l.ctx.Value("userID"))
//	userID := ctxtool.GetUserIDFromCTX(l.ctx)
//
//	//find user
//	user, err := l.svcCtx.User.FindOne(l.ctx, userID)
//	if err != nil && err != sqlx.ErrNotFound {
//		//return nil, errors.Wrap(errx.NewErrCode(errx.DB_ERROR), fmt.Sprintf("UserProfile - user db err:%v, userID:%v", err, userID))
//		return nil, errx.NewCommonMessage(errx.DB_ERROR, err.Error())
//	}
//
//	if user == nil {
//		//return nil, errors.Wrap(errx.NewErrCode(errx.USER_NOT_EXIST), fmt.Sprintf("UserProfile - user db USER NOT FOUND err: %v, userID: %v", err, userID))
//		return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
//	}
//	//list, err := l.svcCtx.LikedMovie.FindAllByUserIDWithMovieInfo(l.ctx, int64(id))
//	//if err != nil {
//	//	return nil, errorx.NewDefaultCodeError(err.Error())
//	//}
//
//	//var likedMovieList []*types.LikedMovieInfo
//	//for _, v := range list {
//	//	var genres []types.GenreInfo
//	//	json.Unmarshal(v.Genres, &genres)
//	//	likedMovieList = append(likedMovieList, &types.LikedMovieInfo{
//	//		Genres:       genres,
//	//		MovieId:      v.MovieId,
//	//		MovieName:    v.MovieTitle,
//	//		MoviePoster:  v.MoviePoster,
//	//		MovieVoteAvg: v.MovieVoteAvg,
//	//	})
//	//}
//
//	return &types.UserProfileResp{
//		ID:     user.Id,
//		Name:   user.Name,
//		Email:  user.Email,
//		Avatar: user.Avatar,
//		Cover:  user.Cover,
//		//LikedMovies: likedMovieList,
//	}, nil
//}

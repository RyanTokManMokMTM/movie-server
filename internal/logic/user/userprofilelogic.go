package user

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/errorx"
	"strconv"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *UserProfileLogic) UserProfile(req *types.UserProfileRequest) (resp *types.UserProfileResponse, err error) {
	// todo: add your logic here and delete this line
	//logx.Infof("userId: %v", l.ctx.Value("userID"))
	userID := fmt.Sprintf("%v", l.ctx.Value("userID"))
	id, _ := strconv.Atoi(userID)
	res, err := l.svcCtx.User.FindOne(l.ctx, int64(id))
	if err != nil {
		return nil, errorx.NewDefaultCodeError(err.Error())
	}

	//list, err := l.svcCtx.LikedMovie.FindAllByUserIDWithMovieInfo(l.ctx, int64(id))
	//if err != nil {
	//	return nil, errorx.NewDefaultCodeError(err.Error())
	//}

	//var likedMovieList []*types.LikedMovieInfo
	//for _, v := range list {
	//	var genres []types.GenreInfo
	//	json.Unmarshal(v.Genres, &genres)
	//	likedMovieList = append(likedMovieList, &types.LikedMovieInfo{
	//		Genres:       genres,
	//		MovieID:      v.MovieId,
	//		MovieName:    v.MovieTitle,
	//		MoviePoster:  v.MoviePoster,
	//		MovieVoteAvg: v.MovieVoteAvg,
	//	})
	//}

	return &types.UserProfileResponse{
		ID:     res.Id,
		Name:   res.Name,
		Email:  res.Email,
		Avatar: res.Avatar,
		Cover:  res.Cover,
		//LikedMovies: likedMovieList,
	}, nil
}

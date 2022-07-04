package liked_movie

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LikedMoviesModel = (*customLikedMoviesModel)(nil)

type (
	// LikedMoviesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLikedMoviesModel.
	LikedMoviesModel interface {
		likedMoviesModel
	}

	customLikedMoviesModel struct {
		*defaultLikedMoviesModel
	}
)

// NewLikedMoviesModel returns a model for the database table.
func NewLikedMoviesModel(conn sqlx.SqlConn, c cache.CacheConf) LikedMoviesModel {
	return &customLikedMoviesModel{
		defaultLikedMoviesModel: newLikedMoviesModel(conn, c),
	}
}

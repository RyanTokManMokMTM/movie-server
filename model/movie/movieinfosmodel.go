package movie

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MovieInfosModel = (*customMovieInfosModel)(nil)

type (
	// MovieInfosModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMovieInfosModel.
	MovieInfosModel interface {
		movieInfosModel
	}

	customMovieInfosModel struct {
		*defaultMovieInfosModel
	}
)

// NewMovieInfosModel returns a model for the database table.
func NewMovieInfosModel(conn sqlx.SqlConn, c cache.CacheConf) MovieInfosModel {
	return &customMovieInfosModel{
		defaultMovieInfosModel: newMovieInfosModel(conn, c),
	}
}

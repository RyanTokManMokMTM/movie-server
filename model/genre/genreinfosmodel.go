package genre

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GenreInfosModel = (*customGenreInfosModel)(nil)

type (
	// GenreInfosModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGenreInfosModel.
	GenreInfosModel interface {
		genreInfosModel
	}

	customGenreInfosModel struct {
		*defaultGenreInfosModel
	}
)

// NewGenreInfosModel returns a model for the database table.
func NewGenreInfosModel(conn sqlx.SqlConn, c cache.CacheConf) GenreInfosModel {
	return &customGenreInfosModel{
		defaultGenreInfosModel: newGenreInfosModel(conn, c),
	}
}

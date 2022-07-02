package list_movie

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ListsMoviesModel = (*customListsMoviesModel)(nil)

type (
	// ListsMoviesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customListsMoviesModel.
	ListsMoviesModel interface {
		listsMoviesModel
	}

	customListsMoviesModel struct {
		*defaultListsMoviesModel
	}
)

// NewListsMoviesModel returns a model for the database table.
func NewListsMoviesModel(conn sqlx.SqlConn, c cache.CacheConf) ListsMoviesModel {
	return &customListsMoviesModel{
		defaultListsMoviesModel: newListsMoviesModel(conn, c),
	}
}

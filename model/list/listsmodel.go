package list

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ListsModel = (*customListsModel)(nil)

type (
	// ListsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customListsModel.
	ListsModel interface {
		listsModel
	}

	customListsModel struct {
		*defaultListsModel
	}
)

// NewListsModel returns a model for the database table.
func NewListsModel(conn sqlx.SqlConn, c cache.CacheConf) ListsModel {
	return &customListsModel{
		defaultListsModel: newListsModel(conn, c),
	}
}

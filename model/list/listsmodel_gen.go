// Code generated by goctl. DO NOT EDIT!

package list

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	listsFieldNames          = builder.RawFieldNames(&Lists{})
	listsRows                = strings.Join(listsFieldNames, ",")
	listsRowsExpectAutoSet   = strings.Join(stringx.Remove(listsFieldNames, "`list_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	listsRowsWithPlaceHolder = strings.Join(stringx.Remove(listsFieldNames, "`list_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	JoinedUserlistsFieldNames     = builder.RawFieldNames(&JoinedUserList{})
	JoinedUserlistsRows           = strings.Join(JoinedUserlistsFieldNames, ",")
	JoinedUserlistsRowsExpectTime = strings.Join(stringx.Remove(JoinedUserlistsFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	JoinedUserListsAmbiguousRows  = strings.Join([]string{"lists.create_time", "lists.update_time"}, ",")
	cacheMovieListsListIdPrefix   = "cache:movie:lists:listId:"
)

type (
	listsModel interface {
		Insert(ctx context.Context, data *Lists) (sql.Result, error)
		FindOne(ctx context.Context, listId int64) (*Lists, error)
		Update(ctx context.Context, newData *Lists) error
		Delete(ctx context.Context, listId int64) error

		FindAllByUserID(ctx context.Context, userID int64) ([]*Lists, error)
		FindAllByUpdateTimeDESC(ctx context.Context) ([]*JoinedUserList, error)

		FindOneByUserIDAndListId(ctx context.Context, userID, listID int64) (*Lists, error)
	}

	defaultListsModel struct {
		sqlc.CachedConn
		table string
	}

	Lists struct {
		ListId     int64     `db:"list_id"`
		ListTitle  string    `db:"list_title"`
		UserId     int64     `db:"user_id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}

	JoinedUserList struct {
		ListId    int64  `db:"list_id"`
		ListTitle string `db:"list_title"`
		UserId    int64  `db:"user_id"`
		//CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		Name       string    `db:"name"`
		Avatar     string    `db:"avatar"`
		Email      string    `db:"email"`
	}
)

func newListsModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultListsModel {
	return &defaultListsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`lists`",
	}
}

func (m *defaultListsModel) Delete(ctx context.Context, listId int64) error {
	movieListsListIdKey := fmt.Sprintf("%s%v", cacheMovieListsListIdPrefix, listId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `list_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, listId)
	}, movieListsListIdKey)
	return err
}

func (m *defaultListsModel) FindOne(ctx context.Context, listId int64) (*Lists, error) {
	movieListsListIdKey := fmt.Sprintf("%s%v", cacheMovieListsListIdPrefix, listId)
	var resp Lists
	err := m.QueryRowCtx(ctx, &resp, movieListsListIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `list_id` = ? limit 1", listsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, listId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultListsModel) Insert(ctx context.Context, data *Lists) (sql.Result, error) {
	movieListsListIdKey := fmt.Sprintf("%s%v", cacheMovieListsListIdPrefix, data.ListId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, listsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ListTitle, data.UserId)
	}, movieListsListIdKey)
	return ret, err
}

func (m *defaultListsModel) Update(ctx context.Context, data *Lists) error {
	movieListsListIdKey := fmt.Sprintf("%s%v", cacheMovieListsListIdPrefix, data.ListId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `list_id` = ?", m.table, listsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ListTitle, data.UserId, data.ListId)
	}, movieListsListIdKey)
	return err
}

func (m *defaultListsModel) FindAllByUserID(ctx context.Context, userID int64) ([]*Lists, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE lists.user_id = ?;", listsRows, m.table)
	var res []*Lists
	err := m.QueryRowsNoCacheCtx(ctx, &res, query, userID)

	switch err {
	case nil:
		return res, nil
	default:
		return nil, err
	}
}

func (m *defaultListsModel) FindAllByUpdateTimeDESC(ctx context.Context) ([]*JoinedUserList, error) {
	query := fmt.Sprintf("SELECT %s,%s FROM %s INNER JOIN users ON lists.user_id = users.id ORDER BY lists.update_time ASC", JoinedUserlistsRowsExpectTime, JoinedUserListsAmbiguousRows, m.table)
	logx.Info(query)
	var res []*JoinedUserList
	err := m.QueryRowsNoCacheCtx(ctx, &res, query)
	switch err {
	case nil:
		return res, nil
	default:
		return nil, err
	}
}

func (m *defaultListsModel) FindOneByUserIDAndListId(ctx context.Context, userID, listID int64) (*Lists, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE list_id = ? AND user_id = ?", listsRows, m.table)
	logx.Info(query)
	var res Lists
	err := m.QueryRowNoCacheCtx(ctx, &res, query, listID, userID)
	switch err {
	case nil:
		return &res, nil
	default:
		return nil, err
	}
}

func (m *defaultListsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMovieListsListIdPrefix, primary)
}

func (m *defaultListsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `list_id` = ? limit 1", listsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultListsModel) tableName() string {
	return m.table
}

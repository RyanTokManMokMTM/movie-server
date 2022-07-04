// Code generated by goctl. DO NOT EDIT!

package liked_movie

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	likedMoviesFieldNames          = builder.RawFieldNames(&LikedMovies{})
	likedMoviesRows                = strings.Join(likedMoviesFieldNames, ",")
	likedMoviesRowsExpectAutoSet   = strings.Join(stringx.Remove(likedMoviesFieldNames, "`liked_movie_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	likedMoviesRowsWithPlaceHolder = strings.Join(stringx.Remove(likedMoviesFieldNames, "`liked_movie_id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	LikedMovieInfoField     = builder.RawFieldNames(&LikedMovieInfo{})
	LikedMovieInfoFieldRows = strings.Join(LikedMovieInfoField, ",")

	cacheMovieLikedMoviesLikedMovieIdPrefix  = "cache:movie:likedMovies:likedMovieId:"
	cacheMovieLikedMoviesUserIdMovieIdPrefix = "cache:movie:likedMovies:userId:movieId:"
)

type (
	likedMoviesModel interface {
		Insert(ctx context.Context, data *LikedMovies) (sql.Result, error)
		FindOne(ctx context.Context, likedMovieId int64) (*LikedMovies, error)
		FindOneByUserIdMovieId(ctx context.Context, userId int64, movieId int64) (*LikedMovies, error)
		Update(ctx context.Context, newData *LikedMovies) error
		Delete(ctx context.Context, likedMovieId int64) error

		FindAllByUserIDWithMovieInfo(ctx context.Context, userID int64) ([]*LikedMovieInfo, error)
	}

	defaultLikedMoviesModel struct {
		sqlc.CachedConn
		table string
	}

	LikedMovies struct {
		LikedMovieId int64     `db:"liked_movie_id"`
		UserId       int64     `db:"user_id"`
		MovieId      int64     `db:"movie_id"`
		CreateTime   time.Time `db:"create_time"`
		UpdateTime   time.Time `db:"update_time"`
	}

	JoinedMovies struct {
		LikedMovieId int64 `db:"liked_movie_id"`
		UserId       int64 `db:"user_id"`
		MovieId      int64 `db:"movie_id"`
		//CreateTime   time.Time `db:"create_time"`
		//UpdateTime   time.Time `db:"update_time"`

		MovieID     int64  `db:"movie_id"`
		MovieTitle  string `db:"title"`
		MoviePoster string `db:"poster_path"`
	}

	LikedMovieInfo struct {
		LikedMovieId int64   `db:"liked_movie_id"`
		MovieId      int64   `db:"movie_id"`
		MovieTitle   string  `db:"title"`
		MoviePoster  string  `db:"poster_path"`
		MovieVoteAvg float32 `db:"vote_average"`
		Genres       []uint8 `db:"genres"`
	}

	//GenreInfo struct {
	//	ID   int64  `db:"id"`
	//	Name string `db:"name"`
	//}
)

func newLikedMoviesModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultLikedMoviesModel {
	return &defaultLikedMoviesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`liked_movies`",
	}
}

func (m *defaultLikedMoviesModel) Delete(ctx context.Context, likedMovieId int64) error {
	data, err := m.FindOne(ctx, likedMovieId)
	if err != nil {
		return err
	}

	movieLikedMoviesLikedMovieIdKey := fmt.Sprintf("%s%v", cacheMovieLikedMoviesLikedMovieIdPrefix, likedMovieId)
	movieLikedMoviesUserIdMovieIdKey := fmt.Sprintf("%s%v:%v", cacheMovieLikedMoviesUserIdMovieIdPrefix, data.UserId, data.MovieId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `liked_movie_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, likedMovieId)
	}, movieLikedMoviesLikedMovieIdKey, movieLikedMoviesUserIdMovieIdKey)
	return err
}

func (m *defaultLikedMoviesModel) FindOne(ctx context.Context, likedMovieId int64) (*LikedMovies, error) {
	movieLikedMoviesLikedMovieIdKey := fmt.Sprintf("%s%v", cacheMovieLikedMoviesLikedMovieIdPrefix, likedMovieId)
	var resp LikedMovies
	err := m.QueryRowCtx(ctx, &resp, movieLikedMoviesLikedMovieIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `liked_movie_id` = ? limit 1", likedMoviesRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, likedMovieId)
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

func (m *defaultLikedMoviesModel) FindOneByUserIdMovieId(ctx context.Context, userId int64, movieId int64) (*LikedMovies, error) {
	movieLikedMoviesUserIdMovieIdKey := fmt.Sprintf("%s%v:%v", cacheMovieLikedMoviesUserIdMovieIdPrefix, userId, movieId)
	var resp LikedMovies
	err := m.QueryRowIndexCtx(ctx, &resp, movieLikedMoviesUserIdMovieIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `movie_id` = ? limit 1", likedMoviesRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, movieId); err != nil {
			return nil, err
		}
		return resp.LikedMovieId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLikedMoviesModel) Insert(ctx context.Context, data *LikedMovies) (sql.Result, error) {
	movieLikedMoviesLikedMovieIdKey := fmt.Sprintf("%s%v", cacheMovieLikedMoviesLikedMovieIdPrefix, data.LikedMovieId)
	movieLikedMoviesUserIdMovieIdKey := fmt.Sprintf("%s%v:%v", cacheMovieLikedMoviesUserIdMovieIdPrefix, data.UserId, data.MovieId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, likedMoviesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.MovieId)
	}, movieLikedMoviesLikedMovieIdKey, movieLikedMoviesUserIdMovieIdKey)
	return ret, err
}

func (m *defaultLikedMoviesModel) Update(ctx context.Context, newData *LikedMovies) error {
	data, err := m.FindOne(ctx, newData.LikedMovieId)
	if err != nil {
		return err
	}

	movieLikedMoviesLikedMovieIdKey := fmt.Sprintf("%s%v", cacheMovieLikedMoviesLikedMovieIdPrefix, data.LikedMovieId)
	movieLikedMoviesUserIdMovieIdKey := fmt.Sprintf("%s%v:%v", cacheMovieLikedMoviesUserIdMovieIdPrefix, data.UserId, data.MovieId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `liked_movie_id` = ?", m.table, likedMoviesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.MovieId, newData.LikedMovieId)
	}, movieLikedMoviesLikedMovieIdKey, movieLikedMoviesUserIdMovieIdKey)
	return err
}

func (m *defaultLikedMoviesModel) FindAllByUserIDWithMovieInfo(ctx context.Context, userID int64) ([]*LikedMovieInfo, error) {
	query := fmt.Sprintf("SELECT "+
		"likedMovies.*, JSON_ARRAYAGG(JSON_OBJECT( 'id', genre_infos.genre_id, 'name', genre_infos.`name` )) AS genres "+
		"FROM `genres_movies` "+
		"INNER JOIN ( "+
		"SELECT liked_movie_id, movie_infos.movie_id, movie_infos.`title`,vote_average,poster_path "+
		"FROM %s "+
		"INNER JOIN movie_infos ON movie_infos.movie_id = liked_movies.movie_id WHERE liked_movies.user_id = ? ) AS likedMovies "+
		"ON likedMovies.movie_id = genres_movies.movie_info_movie_id "+
		"INNER JOIN genre_infos ON genre_infos.genre_id = genres_movies.genre_info_genre_id GROUP BY likedMovies.movie_id LIMIT 10;", m.table)

	var res []*LikedMovieInfo
	err := m.QueryRowsNoCacheCtx(ctx, &res, query, userID)
	switch err {
	case nil:
		return res, nil
	default:
		return nil, err
	}
}

func (m *defaultLikedMoviesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMovieLikedMoviesLikedMovieIdPrefix, primary)
}

func (m *defaultLikedMoviesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `liked_movie_id` = ? limit 1", likedMoviesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLikedMoviesModel) tableName() string {
	return m.table
}

package svc

import (
	"github.com/ryantokmanmokmtm/movie-server/internal/config"
	"github.com/ryantokmanmokmtm/movie-server/model/genre"
	"github.com/ryantokmanmokmtm/movie-server/model/liked_movie"
	"github.com/ryantokmanmokmtm/movie-server/model/list"
	"github.com/ryantokmanmokmtm/movie-server/model/list_movie"
	"github.com/ryantokmanmokmtm/movie-server/model/movie"
	"github.com/ryantokmanmokmtm/movie-server/model/post"
	"github.com/ryantokmanmokmtm/movie-server/model/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config     config.Config
	User       user.UsersModel
	Movie      movie.MovieInfosModel
	Genre      genre.GenreInfosModel
	List       list.ListsModel
	ListMovie  list_movie.ListsMoviesModel
	LikedMovie liked_movie.LikedMoviesModel
	PostModel  post.PostsModel

	GormEngine *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:     c,
		User:       user.NewUsersModel(conn, c.CacheRedis),
		Movie:      movie.NewMovieInfosModel(conn, c.CacheRedis),
		Genre:      genre.NewGenreInfosModel(conn, c.CacheRedis),
		List:       list.NewListsModel(conn, c.CacheRedis),
		ListMovie:  list_movie.NewListsMoviesModel(conn, c.CacheRedis),
		LikedMovie: liked_movie.NewLikedMoviesModel(conn, c.CacheRedis),
		PostModel:  post.NewPostsModel(conn, c.CacheRedis),
		GormEngine: gormConfig(c.MySQL.DataSource),
	}
}

func gormConfig(mysqlConfig string) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: mysqlConfig,
	}))

	if err != nil {
		panic("GORM INIT ERROR")
	}

	return db
}

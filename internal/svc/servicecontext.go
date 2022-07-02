package svc

import (
	"github.com/ryantokmanmokmtm/movie-server/internal/config"
	"github.com/ryantokmanmokmtm/movie-server/model/genre"
	"github.com/ryantokmanmokmtm/movie-server/model/list"
	"github.com/ryantokmanmokmtm/movie-server/model/movie"
	"github.com/ryantokmanmokmtm/movie-server/model/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	User   user.UsersModel
	Movie  movie.MovieInfosModel
	Genre  genre.GenreInfosModel
	List   list.ListsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config: c,
		User:   user.NewUsersModel(conn, c.CacheRedis),
		Movie:  movie.NewMovieInfosModel(conn, c.CacheRedis),
		Genre:  genre.NewGenreInfosModel(conn, c.CacheRedis),
		List:   list.NewListsModel(conn, c.CacheRedis),
	}
}

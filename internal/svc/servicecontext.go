package svc

import (
	"github.com/ryantokmanmokmtm/movie-server/internal/config"
	"github.com/ryantokmanmokmtm/movie-server/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	DAO    dao.Store //USING DATABASE ACCESS AS MODEL LAYER
}

func NewServiceContext(c config.Config, dao dao.Store) *ServiceContext {
	//conn := sqlx.NewMysql(c.MySQL.DataSource)

	return &ServiceContext{
		Config: c,
		DAO:    dao,
	}
}

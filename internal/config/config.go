package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	MySQL struct {
		DataSource   string
		MaxIdleConns int
		MaxOpenConns int
	}

	CacheRedis cache.CacheConf
	Salt       string
}

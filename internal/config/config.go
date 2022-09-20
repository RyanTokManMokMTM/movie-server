package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Path string `json:",default=./resources"`
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	MySQL struct {
		DataSource   string
		MaxIdleConns int
		MaxOpenConns int
	}

	//CacheRedis cache.CacheConf
	Salt     string
	MaxBytes int64
}

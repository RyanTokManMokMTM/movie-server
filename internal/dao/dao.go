package dao

import "gorm.io/gorm"

type DAO struct {
	engine *gorm.DB
}

func NewDAO(engine *gorm.DB) *DAO {
	return &DAO{
		engine: engine,
	}
}

package gormModel

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id         uint      `gorm:"primary_key"`
	Name       string    `gorm:"type:varchar(64)"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Avatar     string    `db:"avatar"`
	Cover      string    `db:"cover"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`

	gorm.Model
}

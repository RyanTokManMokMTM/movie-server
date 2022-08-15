package models

import (
	"github.com/ryantokmanmokmtm/movie-server/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DefaultModel struct {
	CreatedAt time.Time      `json:"-" gorm:"type:timestamp"`
	UpdatedAt time.Time      `json:"-" gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:timestamp" json:"-"`
}

func NewEngine(c config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.MySQL.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //all table lower case User->use
		},
	})

	if err != nil {
		logx.Error(err)
		panic(err.(any))
	}

	sql, err := db.DB()
	if err != nil {
		logx.Error(err)
		panic(err.(any))
	}

	sql.SetMaxIdleConns(c.MySQL.MaxIdleConns)
	sql.SetMaxOpenConns(c.MySQL.MaxOpenConns)
	db.AutoMigrate(&GenreInfo{})
	db.AutoMigrate(&MovieInfo{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserMovie{})
	db.AutoMigrate(&Friend{})
	if err := db.SetupJoinTable(&User{}, "MovieInfos", &UserMovie{}); err != nil {
		panic(err.Error())
	}

	if err := db.SetupJoinTable(&User{}, "Friends", &Friend{}); err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&List{})
	db.AutoMigrate(&ListMovie{})
	if err := db.SetupJoinTable(&List{}, "MovieInfos", &ListMovie{}); err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

	return db
}

func CloseDB(db *gorm.DB) {
	sql, err := db.DB()
	if err != nil {
		logx.Error(err)
		panic(err.(any))
	}

	sql.Close()
}

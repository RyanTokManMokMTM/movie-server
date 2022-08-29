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
	db.AutoMigrate(&PostLiked{})
	db.AutoMigrate(&CommentLiked{})

	db.AutoMigrate(&List{})
	db.AutoMigrate(&ListMovie{})
	if err := db.SetupJoinTable(&List{}, "MovieInfos", &ListMovie{}); err != nil {
		panic(err.(any))
	}
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
	//db.AutoMigrate(&UserInterestedGenre{})
	if err := db.SetupJoinTable(&User{}, "MovieInfos", &UserMovie{}); err != nil {
		panic(err.(any))
	}
	if err := db.SetupJoinTable(&User{}, "Friends", &Friend{}); err != nil {
		panic(err.(any))
	}
	if err := db.SetupJoinTable(&User{}, "PostsLiked", &PostLiked{}); err != nil {
		panic(err.(any))
	}
	if err := db.SetupJoinTable(&Post{}, "PostsLiked", &PostLiked{}); err != nil {
		panic(err.(any))
	}

	if err := db.SetupJoinTable(&User{}, "CommentLiked", &CommentLiked{}); err != nil {
		panic(err.(any))
	}
	//if err := db.SetupJoinTable(&User{}, "InterestedGenre", &UserInterestedGenre{}); err != nil {
	//	panic(err.(any))
	//}

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
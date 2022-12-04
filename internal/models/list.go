package models

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type List struct {
	ListId    uint   `gorm:"primaryKey;not null;autoIncrement"`
	ListTitle string `gorm:"not null;type:varchar(255);"` //unique?
	ListIntro string `gorm:"null;type:longtext"`
	UserId    uint   `gorm:"not null;type:bigint;unsigned"`
	DefaultModel

	User User `gorm:"foreignKey:UserId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	//List has many movies
	//Movie can add to many list
	MovieInfos []MovieInfo `gorm:"many2many:lists_movies;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (m *List) TableName() string {
	return "lists"
}

func (m *List) CreateNewList(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Model(&m).Create(&m).Error
}

func (m *List) FindOneList(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Model(&m).Where("list_id = ?", m.ListId).Preload("MovieInfos", func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(20)
	}).First(&m).Error
}

func (m *List) FindOneUserList(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).First(&m).Error
}

func (m *List) FindAllList(ctx context.Context, db *gorm.DB, limit, pageOffset int) ([]*List, int64, error) {
	var lists []*List
	var count int64 = 0
	if err := db.Debug().WithContext(ctx).Model(&m).
		Where("user_id = ?", m.UserId).
		Count(&count).
		Offset(pageOffset).
		Limit(limit).
		Find(&lists).Error; err != nil {
		return nil, 0, err
	}
	return lists, count, nil
}

func (m *List) CountListMovies(ctx context.Context, db *gorm.DB) (int64, error) {
	return db.Debug().WithContext(ctx).Model(&m).Association("MovieInfos").Count(), nil
}

func (m *List) GetUserListsID(ctx context.Context, db *gorm.DB) ([]int, error) {
	var listsIDs []int
	err := db.Debug().WithContext(ctx).Model(&m).Select("list_id").Where("user_id = ?", m.UserId).Scan(&listsIDs).Error
	return listsIDs, err
}

func (m *List) UpdateList(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Model(&m).Where("list_id = ?", m.ListId).Updates(&m).Error
}

func (m *List) DeleteList(ctx context.Context, db *gorm.DB) error {
	//need to remove all collected record in this list
	return db.Transaction(func(tx *gorm.DB) error {

		//Remove all list movie
		if err := tx.Debug().WithContext(ctx).Model(&m).Association("MovieInfos").Clear(); err != nil {
			return err
		}

		//delete the list
		return tx.Debug().WithContext(ctx).Delete(&m).Error
	})
}

func (m *List) InsertMovieToList(ctx context.Context, db *gorm.DB, info *MovieInfo) error {
	return db.Debug().WithContext(ctx).Model(&m).Association("MovieInfos").Append(info)
}

func (m *List) RemoveMovieFromList(ctx context.Context, db *gorm.DB, info *MovieInfo) error {
	return db.Debug().WithContext(ctx).Model(&m).Association("MovieInfos").Delete(info)
}

func (m *List) RemoveMoviesFromList(ctx context.Context, db *gorm.DB, movieIds []uint) error {

	var infos []*MovieInfo
	for _, v := range movieIds {
		infos = append(infos, &MovieInfo{
			Id: v,
		})
	}

	return db.Debug().WithContext(ctx).Model(&m).Association("MovieInfos").Delete(&infos)

}

//TODO - Check Movie is already collected by user - return a list info
func (m *List) FindOneMovieFromList(ctx context.Context, db *gorm.DB, info *MovieInfo) error {
	return db.Debug().WithContext(ctx).Model(&m).Where("user_id = ?", m.ListId).Association("MovieInfos").Find(&info)
}

//TODO: Custom Data For Getting List Movie
type ListMovieInfoWithCreateTime struct {
	Id          uint      `json:"id"`
	PosterPath  string    `json:"poster_path"`
	Title       string    `json:"title"`
	VoteAverage float64   `json:"vote_average"`
	CreatedAt   time.Time `json:"created_at"`
}

func (m *List) FindListMovieByCreateTime(ctx context.Context, db *gorm.DB, createTime uint, limit int) ([]ListMovieInfoWithCreateTime, int64, error) {
	movies := make([]ListMovieInfoWithCreateTime, 0)

	if err := db.Debug().WithContext(ctx).
		Table((&MovieInfo{}).TableName()).
		Select("movie_infos.id,movie_infos.poster_path,movie_infos.title,movie_infos.vote_average,lists_movies.created_at").
		Joins("INNER JOIN lists_movies ON lists_movies.movie_info_id = movie_infos.id WHERE lists_movies.list_list_id = ? AND lists_movies.created_at > ?", m.ListId, time.Unix(int64(createTime), 0)).
		Limit(limit).Find(&movies).Error; err != nil {
		return nil, 0, err
	}

	count := db.Debug().WithContext(ctx).Model(&m).Association("MovieInfos").Count()
	//logx.Info(count)
	return movies, count, nil
}

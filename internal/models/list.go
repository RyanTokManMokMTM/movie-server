package models

import (
	"context"
	"gorm.io/gorm"
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
		Preload("MovieInfos").
		Count(&count).
		Offset(pageOffset).
		Limit(limit).
		Find(&lists).Error; err != nil {
		return nil, 0, err
	}
	return lists, count, nil
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
	l := List{}
	var infos []*MovieInfo
	for _, v := range movieIds {
		infos = append(infos, &MovieInfo{
			Id: v,
		})
	}

	return db.Debug().WithContext(ctx).Model(&l).Association("MovieInfos").Delete(infos)

}

//TODO - Check Movie is already collected by user - return a list info
func (m *List) FindOneMovieFromList(ctx context.Context, db *gorm.DB, info *MovieInfo) error {
	return db.Debug().WithContext(ctx).Model(&m).Where("user_id = ?", m.ListId).Association("MovieInfos").Find(&info)
}

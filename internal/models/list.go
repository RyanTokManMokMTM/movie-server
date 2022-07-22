package models

import (
	"context"
	"gorm.io/gorm"
)

type List struct {
	ListId    uint   `gorm:"primaryKey;not null;autoIncrement"`
	ListTitle string `gorm:"not null;type:varchar(255);"` //unique?
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
	return db.Debug().WithContext(ctx).Model(&m).Where("list_id = ?", m.ListId).First(&m).Error
}

func (m *List) FindAllList(ctx context.Context, db *gorm.DB) ([]*List, error) {
	var lists []*List
	if err := db.Debug().WithContext(ctx).Model(&m).Where("user_id = ?", m.UserId).Find(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func (m *List) UpdateList(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Model(&m).Where("list_id = ?", m.ListId).Updates(&m).Error
}

func (m *List) DeleteList(ctx context.Context, db *gorm.DB) error {
	return db.Debug().WithContext(ctx).Model(&m).Where("list_id = ? AND user_id = ?", m.ListId, m.UserId).Delete(&m).Error
}

func (m *List) InsertMovieToList(ctx context.Context, db *gorm.DB, info *MovieInfo) error {
	return db.Debug().WithContext(ctx).Model(&m).Where("list_id = ?", m.ListId).Association("MovieInfos").Append(&info)

}

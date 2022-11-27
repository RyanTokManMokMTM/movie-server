package models

import (
	"gorm.io/gorm"
)

//UserInterestedGenre - storing user interested movie genre for some purpose - Recommend movie? Recommend friend? etc...
type UserInterestedGenre struct {
	UserId           uint `gorm:"primaryKey,not null"`
	GenreInfoGenreId uint `gorm:"primaryKey,not null"`
	State            uint `gorm:"not null;unsigned;type:tinyint(1)"`
	DefaultModel
}

func (m *UserInterestedGenre) TableName() string {
	return "users_genres"
}

func (m *UserInterestedGenre) BeforeCreate(db *gorm.DB) error {
	m.State = 1 //when create set to follow
	return nil
}

package models

type Comment struct {
	CommentID uint `gorm:"primaryKey;not null;autoIncrement"`
	PostID    uint
	UserID    uint
	Comment   string

	Post Post `gorm:"foreignKey:PostID;references:PostId"`
	User User `gorm:"foreignKey:UserID;references:Id"`

	//reply?
	//like
	DefaultModel
}

func (m *Comment) TableName() string {
	return "comments"
}

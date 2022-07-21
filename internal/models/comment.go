package models

type Comment struct {
	CommentID uint   `gorm:"primaryKey;not null;autoIncrement"`
	PostID    uint   `gorm:"not null;type:bigint;unsigned;"`
	UserID    uint   `gorm:"not null;type:bigint;unsigned"`
	Comment   string `gorm:"not null;type:longtext"`

	//Post Post `gorm:"foreignKey:PostID;references:PostId"`
	User User `gorm:"foreignKey:UserID;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	//reply?
	//like
	DefaultModel
}

func (m *Comment) TableName() string {
	return "comments"
}

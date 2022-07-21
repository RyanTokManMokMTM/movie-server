package models

type Post struct {
	PostId      uint `gorm:"primaryKey;not null;autoIncrement"`
	PostTitle   string
	PostDesc    string
	UserId      uint
	MovieInfoId uint
	PostLike    int64

	User      User      `gorm:"foreignKey:UserId;references:Id"`
	MovieInfo MovieInfo `gorm:"foreignKey:MovieInfoId;references:MovieId"`
	DefaultModel
}

func (m *Post) TableName() string {
	return "posts"
}

func (m *Post) CreateNewPost() {

}

func (m *Post) UpdatePost() {

}

func (m *Post) DeletePost() {

}

func (m *Post) GetPostInfo() {}

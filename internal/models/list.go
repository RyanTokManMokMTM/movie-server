package models

type List struct {
	ListId    uint `gorm:"primaryKey;not null;autoIncrement"`
	ListTitle string
	UserId    uint
	DefaultModel

	User User `gorm:"foreignKey:UserId;references:id"`

	//List has many movies
	//Movie can add to many list
	MovieInfos []MovieInfo `gorm:"many2many:lists_movies"`
}

func (m *List) TableName() string {
	return "lists"
}

func (m *List) CreateNewList() string {
	return "lists"
}

func (m *List) FindOneList() string {
	return "lists"
}

func (m *List) FindAllList() string {
	return "lists"
}

func (m *List) UpdateList() string {
	return "lists"
}

func (m *List) DeleteList() string {
	return "lists"
}

func (m *List) InsertMovieToList() string {
	return "lists"
}
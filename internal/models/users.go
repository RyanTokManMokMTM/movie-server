package models

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey;not null;autoIncrement"`
	Name     string `gorm:"not null;type:varchar(64)"`
	Email    string `gorm:"not null;type:varchar(64)"`
	Password string `gorm:"not null;type:varchar(64)"`
	Avatar   string `gorm:"not null;type:varchar(255)"`
	Cover    string `gorm:"not null;type:varchar(255)"`

	//can have a lot of list
	List            []List      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	MovieInfos      []MovieInfo `gorm:"many2many:users_movies;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Posts           []Post      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PostsLiked      []Post      `gorm:"many2many:post_liked;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CommentLiked    []Comment   `gorm:"many2many:comment_liked;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	InterestedGenre []GenreInfo `gorm:"many2many:users_genres;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Rooms           []Room      `gorm:"many2many:users_rooms;"` //for all available room of that user
	Friends         []User      `gorm:"many2many:friendship"`
	DefaultModel
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) Insert(ctx context.Context, db *gorm.DB) error {
	logx.Infof("UserDB - Create User:%+v \n", m)
	return db.WithContext(ctx).Debug().Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Debug().Create(&m).Error; err != nil {
			logx.Error(err)
			return err
		}

		logx.Infof("User created %d", m.ID)
		//
		//if err := tx.WithContext(ctx).Debug().Create(&Friend{
		//	UserID: m.ID,
		//}).Error; err != nil {
		//	logx.Error(err)
		//	return err
		//}
		return nil
	})
}

func (m *User) FindOneByID(ctx context.Context, db *gorm.DB) error {
	logx.Infof("UserDB - Find One By ID:%+v \n", m)
	if err := db.WithContext(ctx).Model(&m).Where("id = ?", m.ID).First(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *User) FindOneByEmail(ctx context.Context, db *gorm.DB) error {
	logx.Infof("UserDB - Find One By Email:%+v\n", m)
	if err := db.WithContext(ctx).Model(&m).Where("email = ?", m.Email).First(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *User) UpdateInfo(ctx context.Context, db *gorm.DB) error {
	logx.Infof("UserDB - Update Info:%+v \n", m)
	if err := db.WithContext(ctx).Model(&m).Where("id = ?", m.ID).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (m *User) CreateLikedMovie(ctx context.Context, db *gorm.DB, movie *MovieInfo) error {
	logx.Infof("UserDB - Create User Liked Movie:%+v \n", m)
	return db.WithContext(ctx).Model(&m).Association("MovieInfos").Append(movie)

}

func (m *User) UpdateLikedMovie(ctx context.Context, db *gorm.DB, movie *MovieInfo) error {
	logx.Infof("UserDB - Remove User Liked Movie:%+v \n", m)
	return db.WithContext(ctx).Model(&m).Association("MovieInfos").Delete(movie)
}

func (m *User) GetUserLikedMovies(ctx context.Context, db *gorm.DB) error {
	logx.Infof("UserDB - User Liked Movies:%+v \n", m)
	if err := db.Debug().WithContext(ctx).Preload("MovieInfos", func(db *gorm.DB) *gorm.DB {
		return db.Select("movie_infos.*").Joins("left join users_movies on users_movies.movie_info_id = movie_infos.id").Where("users_movies.state = ?", 1)
	}).Preload("MovieInfos.GenreInfo").Where("id = ?", m.ID).Find(&m).Error; err != nil {
		return err
	}
	return nil
}

//Friend data

func (m *User) FindOneFriend(db *gorm.DB, ctx context.Context, friendID uint) (*User, error) {
	var friend User
	if err := db.WithContext(ctx).Debug().Model(&m).Where(User{
		ID: friendID,
	}).Association("Friends").Find(&friend); err != nil {
		return nil, err
	}

	return &friend, nil
}

//
//func (fd *Friend) GetUserFriend(db *gorm.DB, ctx context.Context) error {
//	return db.WithContext(ctx).Debug().Where("user_id = ?", fd.UserID).Find(&fd).Error
//}

func (m *User) CountFriend(db *gorm.DB, ctx context.Context) int64 {
	return db.WithContext(ctx).Debug().Model(&m).Association("Friends").Count()
}

func (m *User) GetFriendsList(db *gorm.DB, ctx context.Context) ([]*User, error) {
	var friends []*User
	if err := db.WithContext(ctx).Debug().Model(&m).Association("Friends").Find(&friends); err != nil {
		return nil, err
	}
	return friends, nil
}

func (m *User) GetFriendsRoomList(db *gorm.DB, ctx context.Context) error {
	//var friends []*User
	return db.WithContext(ctx).Debug().Model(&m).Preload("Rooms").Preload("Rooms.Users", func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id NOT IN (?)", []uint{m.ID})
	}).First(&m).Error
}

func (m *User) RemoveOne(db *gorm.DB, ctx context.Context, userID, friendID uint) error {
	//Remove an existing Friend
	return db.WithContext(ctx).Debug().Transaction(func(tx *gorm.DB) error {
		//Friendship : A -> B
		if err := tx.WithContext(ctx).Debug().Model(&User{ID: userID}).Association("Friends").Delete(&User{ID: friendID}); err != nil {
			return err
		}

		//Friendship : B -> A
		if err := tx.WithContext(ctx).Debug().Model(&User{ID: friendID}).Association("Friends").Delete(&User{ID: userID}); err != nil {
			return err
		}
		return nil
	})
}

func (m *User) IsFriend(db *gorm.DB, ctx context.Context, friendID uint) (bool, error) {
	var friend *User
	err := db.WithContext(ctx).Debug().Model(&m).Where("id = ?", friendID).Association("Friends").Find(&friend)
	if err != nil {
		return false, err
	}

	logx.Infof("%+v", friend)
	if friend.ID == 0 {
		return false, nil
	}

	return true, nil
}

//
////RETURN A LIST OF USERINFO
//func (m *User) GetFollowingList(ctx context.Context, db *gorm.DB) ([]*User, error) {
//
//	//get friend list id
//	var userId []uint
//	if err := db.Debug().WithContext(ctx).Model(&FriendTemp{}).Select("friend_id").Where("user_id = ?", m.ID).Find(&userId).Error; err != nil {
//		return nil, err
//	}
//
//	var users []*User
//	if err := db.Debug().WithContext(ctx).Model(&m).Where("id IN (?)", userId).Find(&users).Error; err != nil {
//		return nil, err
//	}
//
//	return users, nil
//}
//
////RETURN A LIST OF USERINFO
//func (m *User) GetFollowedList(ctx context.Context, db *gorm.DB) ([]*User, error) {
//
//	//get friend list id
//	var userId []uint
//	if err := db.Debug().WithContext(ctx).Model(&FriendTemp{}).Select("user_id").Where("friend_id = ?", m.ID).Find(&userId).Error; err != nil {
//		return nil, err
//	}
//
//	var users []*User
//	if err := db.Debug().WithContext(ctx).Model(&m).Where("id IN (?)", userId).Find(&users).Error; err != nil {
//		return nil, err
//	}
//	return users, nil
//}

//PostLiked
func (m *User) CreateUserPostLiked(ctx context.Context, db *gorm.DB, post *Post) error {
	return db.Debug().WithContext(ctx).Model(&m).Omit("PostsLiked.*").Association("PostsLiked").Append(post)
}

//CommentLiked
func (m *User) CreateUserCommentLiked(ctx context.Context, db *gorm.DB, comment *Comment) error {
	return db.WithContext(ctx).Debug().Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().WithContext(ctx).Model(&m).Omit("CommentLiked.*").Association("CommentLiked").Append(comment); err != nil {
			return err
		}

		//TODO: Add 1 count to db
		comment.LikesCount = comment.LikesCount + 1
		if err := tx.WithContext(ctx).Debug().Model(&comment).Update("LikesCount", comment.LikesCount).Error; err != nil {
			return err
		}

		return nil
	})
}

//CreateUserGenre
func (m *User) UpdateUserGenreTrans(ctx context.Context, db *gorm.DB, ids []uint) error {
	/*
		Getting All user genre first

		remove all existed,but will not exist after update~
	*/

	return db.Debug().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		user := &User{}
		logx.Infof("transaction begin...")

		//getting all existing genre
		if err := tx.Debug().WithContext(ctx).Preload("InterestedGenre").Take(&user, m.ID).Error; err != nil {
			return err
		}

		//getting genre by ids
		var genres []GenreInfo //all genreInfo need to be inserted into user genre db
		if err := tx.Debug().WithContext(ctx).Model(&GenreInfo{}).Where("genre_id IN (?)", ids).Find(&genres).Error; err != nil {
			return err
		}

		if len(genres) != len(ids) {
			return errors.New("some genre_id is not exist ")
		}

		//logx.Infof("Found Genres : %+v", genres)

		//user genre not contain in genres
		genresToRemove := filter(user.InterestedGenre, func(genre GenreInfo) bool {
			return !contains(genres, genre)
		})

		if len(genresToRemove) > 0 {
			if err := tx.Debug().WithContext(ctx).Model(&user).Association("InterestedGenre").Delete(&genresToRemove); err != nil {
				return err
			}
		}

		//now we need to update
		//getting new user Genres ->
		if err := tx.Debug().WithContext(ctx).Preload("InterestedGenre").Take(&user, m.ID).Error; err != nil {
			return err
		}

		genresToBeUpdate := filter(genres, func(genre GenreInfo) bool {
			return !contains(user.InterestedGenre, genre)
		})

		//logx.Infof("find genre to append to%+v", genresToBeUpdate)
		if len(genresToBeUpdate) > 0 {
			for _, ug := range genresToBeUpdate {
				if err := tx.Debug().WithContext(ctx).Model(&user).Omit("InterestedGenre.*").Association("InterestedGenre").Append(&ug); err != nil {
					return err
				}
			}
		}
		logx.Infof("transaction Completed...")
		return nil
	})

	//return db.Debug().WithContext(ctx).Model(&m).Omit("GenreInfo.*").Association("GenreInfo").Append(genre)
}

func (m *User) FindUserGenres(ctx context.Context, db *gorm.DB) (*[]GenreInfo, error) {
	var genreIds []uint

	if err := db.Debug().WithContext(ctx).Model(&m).Select("genre_info_genre_id").Association("InterestedGenre").Find(&genreIds); err != nil {
		return nil, err
	}

	var genreInfos []GenreInfo
	if err := db.Debug().WithContext(ctx).Model(&genreInfos).Where("genre_id IN (?)", genreIds).Find(&genreInfos).Error; err != nil {
		return nil, err
	}

	return &genreInfos, nil
}

func (m *User) GetUserRooms(ctx context.Context, db *gorm.DB) ([]*Room, error) {
	var rooms []*Room
	if err := db.WithContext(ctx).Debug().Model(&m).Association("Rooms").Find(&rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (m *User) GetUserActiveRooms(ctx context.Context, db *gorm.DB) ([]Room, error) {
	//var rooms []*Room
	//if err := db.WithContext(ctx).Debug().Model(&m).Preload("Rooms").Preload("")

	//get all user active room list
	ur := &UsersRooms{
		UserID: m.ID,
	}

	var roomsIDs []uint
	rooms, err := ur.GetUserActiveRoom(db, ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range rooms {
		roomsIDs = append(roomsIDs, v.RoomID)
	}

	logx.Infof("active room ids %v", roomsIDs)

	if err := db.WithContext(ctx).Debug().Model(&m).Preload("Rooms", func(tx *gorm.DB) *gorm.DB {
		return tx.Where("ID IN (?)", roomsIDs)
	}).First(&m).Error; err != nil {
		return nil, err
	}

	return m.Rooms, nil
}

func (m *User) GetUserRoomsWithMembers(ctx context.Context, db *gorm.DB) error {
	//get all user active room list
	ur := &UsersRooms{
		UserID: m.ID,
	}

	var roomsIDs []uint
	rooms, err := ur.GetUserActiveRoom(db, ctx)
	if err != nil {
		return nil
	}

	for _, v := range rooms {
		roomsIDs = append(roomsIDs, v.RoomID)
	}

	logx.Infof("active room ids %v", roomsIDs)

	return db.WithContext(ctx).Debug().Model(&m).Preload("Rooms", func(tx *gorm.DB) *gorm.DB {
		return tx.Where("ID IN (?)", roomsIDs)
	}).Preload("Rooms.Users").Preload("Rooms.Messages", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("sent_time desc").Limit(10)
	}).Preload("Rooms.Messages.SendUser").First(&m).Error

}

//Util tool
func filter(elements []GenreInfo, handler func(genre GenreInfo) bool) []GenreInfo {
	i := 0
	for _, ele := range elements {
		if handler(ele) {
			elements[i] = ele
			i++
		}
	}

	return elements[:i]
}

func contains(elements []GenreInfo, target GenreInfo) bool {

	//elements contain target??
	for _, ele := range elements {
		if ele.GenreId == target.GenreId {
			return true
		}
	}

	return false //not found
}

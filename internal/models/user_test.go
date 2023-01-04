package models

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

var (
	err    error
	mock   sqlmock.Sqlmock
	db     *sql.DB
	gormDB *gorm.DB
)

func TestMain(m *testing.M) {
	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})

	os.Exit(m.Run())
}

func TestCreateOne(t *testing.T) {
	u := User{
		Name:     "jackson_tmm",
		Email:    "admin@admin.com",
		Password: "31a5e6c3e319732cb1c723228e20cf3c659fd0fa601eb0a03e04e85791a357bb",
		Avatar:   "/test.png",
		Cover:    "/test.png",
		DefaultModel: DefaultModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` "+
		"(`name`,`email`,`password`,`avatar`,`cover`,`friend_notification_count`,`like_notification_count`,`comment_notification_count`,`created_at`,`updated_at`,`deleted_at`) "+
		"VALUES (?,?,?,?,?,?,?,?,?,?,?)").
		WithArgs(u.Name, u.Email, u.Password, u.Avatar, u.Cover, u.FriendNotificationCount, u.LikeNotificationCount, u.CommentNotificationCount, u.CreatedAt, u.UpdatedAt, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := u.CreateOne(context.Background(), gormDB)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindOneByID(t *testing.T) {
	u := User{
		ID:                       1,
		Name:                     "jackson_tmm",
		Email:                    "admin@admin.com",
		Password:                 "31a5e6c3e319732cb1c723228e20cf3c659fd0fa601eb0a03e04e85791a357bb",
		Avatar:                   "/test.png",
		Cover:                    "/test.png",
		FriendNotificationCount:  0,
		LikeNotificationCount:    0,
		CommentNotificationCount: 0,
	}

	mock.ExpectQuery("SELECT `users`.`id`,`users`.`name`,`users`.`email`,`users`.`password`,`users`.`avatar`,`users`.`cover`,`users`.`friend_notification_count`,`users`.`like_notification_count`,`users`.`comment_notification_count` FROM `users` WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1").
		WithArgs(u.ID).
		WillReturnRows(sqlmock.NewRows(u.GetFieldNames()).
			AddRow(u.ID, u.Name, u.Email, u.Password, u.Avatar, u.Cover, u.FriendNotificationCount, u.LikeNotificationCount, u.CommentNotificationCount))

	err := u.FindOneByID(context.Background(), gormDB)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindOneByEmail(t *testing.T) {
	u := User{
		ID:                       1,
		Name:                     "jackson_tmm",
		Email:                    "admin@admin.com",
		Password:                 "31a5e6c3e319732cb1c723228e20cf3c659fd0fa601eb0a03e04e85791a357bb",
		Avatar:                   "/test.png",
		Cover:                    "/test.png",
		FriendNotificationCount:  0,
		LikeNotificationCount:    0,
		CommentNotificationCount: 0,
	}
	//mock.ExpectBegin()
	mock.ExpectQuery("SELECT `users`.`id`,`users`.`name`,`users`.`email`,`users`.`password`,`users`.`avatar`,`users`.`cover`,`users`.`friend_notification_count`,`users`.`like_notification_count`,`users`.`comment_notification_count` FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL AND `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1").
		WithArgs(u.Email, u.ID).WillReturnRows(
		sqlmock.NewRows(u.GetFieldNames()).
			AddRow(u.ID, u.Name, u.Email, u.Password, u.Avatar, u.Cover, u.FriendNotificationCount, u.LikeNotificationCount, u.CommentNotificationCount))
	err := u.FindOneByEmail(context.Background(), gormDB)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateInfo(t *testing.T) {

	testCases := []struct {
		name        string
		userID      uint
		expectedSQL string
		data        any
		model       User
	}{
		{
			name:        "Update Name Only",
			userID:      1,
			expectedSQL: "UPDATE `users` SET `name`=?,`updated_at`=? WHERE `users`.`deleted_at` IS NULL AND `id` = ?",
			data:        "jackson.tmm",
			model: User{
				Name: "jackson.tmm",
			},
		},
		{
			name:        "Update Avatar",
			userID:      1,
			expectedSQL: "UPDATE `users` SET `avatar`=?,`updated_at`=? WHERE `users`.`deleted_at` IS NULL AND `id` = ?",
			data:        "/update-avatar.png",
			model: User{
				Avatar: "/update-avatar.png",
			},
		},

		{
			name:        "Update cover",
			userID:      1,
			expectedSQL: "UPDATE `users` SET `cover`=?,`updated_at`=? WHERE `users`.`deleted_at` IS NULL AND `id` = ?",
			data:        "/update-cover.png",
			model: User{
				Cover: "/update-cover.png",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(test.expectedSQL).WithArgs(test.data, sqlmock.AnyArg(), test.userID).WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
			err := test.model.UpdateInfo(context.Background(), test.userID, gormDB)
			assert.Nil(t, err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}

}

func TestCreateLikedMovie(t *testing.T) {
	//var userID uint = 1
	var movieID uint = 3
	u := User{ID: 1}
	/*
		INSERT INTO `users_movies` (`user_id`,`movie_info_id`) VALUES (5,66) ON DUPLICATE KEY UPDATE `user_id`=`user_id`
			UPDATE `users` SET `updated_at`='2022-12-26 15:32:57.636' WHERE `users`.`deleted_at` IS NULL AND `id` = 5
	*/
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users` SET `updated_at`=? WHERE `users`.`deleted_at` IS NULL AND `id` = ?").
		WithArgs(sqlmock.AnyArg(), u.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO `users_movies` (`user_id`,`movie_info_id`) VALUES (?,?) ON DUPLICATE KEY UPDATE `user_id`=`user_id`").
		WithArgs(u.ID, movieID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err := u.CreateLikedMovie(context.Background(), gormDB, movieID)
	assert.Nil(t, err)

	err = mock.ExpectationsWereMet() //exactly met all expected sql
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCountLikedMovie(t *testing.T) {
	u := User{ID: 1}
	var count int64
	mock.ExpectQuery("SELECT count(*) FROM `movie_infos` JOIN `users_movies` ON `users_movies`.`movie_info_id` = `movie_infos`.`id` AND `users_movies`.`user_id` = ? WHERE `movie_infos`.`deleted_at` IS NULL").WithArgs(u.ID)
	count = u.CountLikedMovie(context.Background(), gormDB)
	assert.GreaterOrEqual(t, count, int64(0))
}

func TestFindOneFriend(t *testing.T) {
	u := User{
		ID: 1,
	}
	var friendID uint = 2 //friend -> UserID?
	expectedSQL := "SELECT `id`,`name` FROM `users` JOIN `friendship` ON `friendship`.`friend_id` = `users`.`id` AND `friendship`.`user_id` = ? WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL"
	result := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "testing")
	mock.ExpectQuery(expectedSQL).
		WithArgs(u.ID, friendID).
		WillReturnRows(result)

	_, err := u.FindOneFriend(gormDB, context.Background(), friendID)
	assert.Nil(t, err)
}

func TestCountFriend(t *testing.T) {
	u := User{ID: 1}
	mock.ExpectQuery("SELECT count(*) FROM `users` JOIN `friendship` ON `friendship`.`friend_id` = `users`.`id` AND `friendship`.`user_id` = ? WHERE `users`.`deleted_at` IS NULL").
		WithArgs(u.ID)
	c := u.CountFriend(gormDB, context.Background())
	assert.GreaterOrEqual(t, c, int64(0))
}

func TestGetFriendsList(t *testing.T) {
	u := User{
		ID: 1,
	}
	// `users`.`id`,
	//`users`.`name`,

	mockResult := mock.NewRows([]string{"id"}).
		AddRow("1").
		AddRow("2")

	mock.ExpectQuery("SELECT `id` FROM `users`  JOIN `friendship` ON `friendship`.`friend_id` = `users`.`id` AND `friendship`.`user_id` = ? WHERE `users`.`deleted_at` IS NULL").
		WithArgs(u.ID).WillReturnRows(mockResult)

	_, err = u.GetFriendsList(gormDB, context.Background())
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//TODO : empty case, single data case need to be tested, cuz the sql quires of all those case are different
func TestGetUserLikedMovies(t *testing.T) {
	u := User{
		ID: 1,
	}
	limit := 10
	mock.ExpectQuery("SELECT `id`, `name` FROM `users` WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = ?").
		WithArgs(u.ID).
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "name"}).AddRow(1, "jacksontmm"))
	//
	////Get liked Movie ID
	mock.ExpectQuery("SELECT * FROM `users_movies` WHERE `users_movies`.`user_id` = ?").
		WithArgs(u.ID).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "movie_info_id"}).
			AddRow(1, 1).
			AddRow(1, 2))

	////Get Movie Info for movie IDs
	mock.ExpectQuery("SELECT `id`, `title`, `poster_path`, `vote_count` FROM `movie_infos` WHERE `movie_infos`.`id` IN (?,?) AND `movie_infos`.`deleted_at` IS NULL LIMIT 10 OFFSET 1").
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "poster_path", "vote_count"}).
			AddRow(1, "test", "/test.png", 10).AddRow(2, "test2", "/test.png", 10))
	//AddRow(false, "/test.png", 1, "en", "test_title", "test_view", 1.2, "/test.png", "2022-01-01", "test_ori_title", "160", false, 10, 10, time.Now(), time.Now(), nil).
	//AddRow(false, "/test.png", 2, "en", "test_title", "test_view", 1.2, "/test.png", "2022-01-01", "test_ori_title", "160", false, 10, 10, time.Now(), time.Now(), nil))

	mock.ExpectQuery("SELECT * FROM `genres_movies` WHERE `genres_movies`.`movie_info_id` IN (?,?)").
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"genre_id", "name"}).
			AddRow(3, "action").
			AddRow(4, "horror").
			AddRow(5, "romantic"))

	err = u.GetUserLikedMovies(context.Background(), gormDB, limit, 1)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

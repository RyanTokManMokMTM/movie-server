package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/ryantokmanmokmtm/movie-server/internal/config"
	mockdb "github.com/ryantokmanmokmtm/movie-server/internal/dao/mock"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"github.com/ryantokmanmokmtm/movie-server/server"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCountPostsLikes(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	daoMock := mockdb.NewMockStore(ctl)

	var configInfo = []byte(`{
	"Path" : "./resources",
	"Auth":{
			"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
			"AccessExpire": 86400
	},
	"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
	"MaxBytes": 1073741824
	}`)

	var c config.Config
	err := json.Unmarshal(configInfo, &c)
	assert.Nil(t, err)

	var userID uint = 1
	var postID uint = 1
	uri := fmt.Sprintf("/api/v1/liked/post/count/%d", postID)

	daoMock.EXPECT().FindOnePostInfo(gomock.Any(), postID).Times(1).Return(&models.Post{
		PostId:      postID,
		PostTitle:   "post_title_test",
		PostDesc:    "post_desc_test",
		UserId:      userID,
		MovieInfoId: 1,
	}, nil)

	r := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

// add more test cases
func TestCreatePostLikes(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	daoMock := mockdb.NewMockStore(ctl)

	var configInfo = []byte(`{
	"Path" : "./resources",
	"Auth":{
			"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
			"AccessExpire": 86400
	},
	"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
	"MaxBytes": 1073741824
	}`)

	var c config.Config
	err := json.Unmarshal(configInfo, &c)
	assert.Nil(t, err)

	var userID uint = 1
	var postID uint = 1
	uri := "/api/v1/liked/post"

	type CreatePost struct {
		PostId uint `json:"post_id"`
	}

	reqData := CreatePost{
		PostId: postID,
	}

	reqBytes, _ := json.Marshal(reqData)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOnePostInfo(gomock.Any(), postID).Times(1).Return(&models.Post{
		PostId:      postID,
		PostTitle:   "post_test_title",
		PostDesc:    "post_test_desc",
		UserId:      2,
		MovieInfoId: 1,
	}, nil)
	//daoMock.EXPECT().FindOnePostLiked(gomock.Any(), userID, postID).Times(1).Return(&models.PostLiked{
	//	UserId:     userID,
	//	PostPostId: postID,
	//}, nil)

	daoMock.EXPECT().FindOnePostLiked(gomock.Any(), userID, postID).Times(1).Return(nil, gorm.ErrRecordNotFound)
	daoMock.EXPECT().CreatePostLiked(gomock.Any(), userID, gomock.Any()).Times(1).Return(nil)
	daoMock.EXPECT().FindOneLikePostNotification(gomock.Any(), gomock.Any(), userID, postID).Times(1).Return(gorm.ErrRecordNotFound)
	daoMock.EXPECT().InsertOnePostLikeNotification(gomock.Any(), postID, userID, gomock.Any(), gomock.Any()).Times(1).Return(nil)

	r := httptest.NewRequest("POST", uri, bytes.NewBuffer(reqBytes))
	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	r.Header.Add("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestIsPostLiked(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	daoMock := mockdb.NewMockStore(ctl)

	var configInfo = []byte(`{
	"Path" : "./resources",
	"Auth":{
			"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
			"AccessExpire": 86400
	},
	"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
	"MaxBytes": 1073741824
	}`)

	var c config.Config
	err := json.Unmarshal(configInfo, &c)
	assert.Nil(t, err)

	var userID uint = 1
	var postID uint = 1
	uri := fmt.Sprintf("/api/v1/liked/post/%d", postID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOnePostInfo(gomock.Any(), postID).Times(1).Return(&models.Post{
		PostId:      postID,
		PostTitle:   "test_post",
		PostDesc:    "test_desc",
		UserId:      userID,
		MovieInfoId: 1,
	}, nil)

	daoMock.EXPECT().FindOnePostLiked(gomock.Any(), userID, postID).Times(1).Return(nil, gorm.ErrRecordNotFound)

	r := httptest.NewRequest("GET", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}
func TestRemovePostLikes(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	daoMock := mockdb.NewMockStore(ctl)

	var configInfo = []byte(`{
	"Path" : "./resources",
	"Auth":{
			"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
			"AccessExpire": 86400
	},
	"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
	"MaxBytes": 1073741824
	}`)

	var c config.Config
	err := json.Unmarshal(configInfo, &c)
	assert.Nil(t, err)

	var userID uint = 1
	var postID uint = 1
	uri := "/api/v1/liked/post"

	type deletePostLiked struct {
		PostId uint `json:"post_id"`
	}

	reqData := deletePostLiked{
		PostId: postID,
	}

	reqBytes, _ := json.Marshal(reqData)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOnePostInfo(gomock.Any(), postID).Times(1).Return(&models.Post{
		PostId:      postID,
		PostTitle:   "test_title",
		PostDesc:    "test_desc",
		UserId:      postID,
		MovieInfoId: 1,
	}, nil)

	daoMock.EXPECT().FindOnePostLiked(gomock.Any(), userID, postID).Times(1).Return(&models.PostLiked{
		UserId:     userID,
		PostPostId: postID,
	}, nil)
	daoMock.EXPECT().DeletePostLikes(gomock.Any(), gomock.Any()).Times(1).Return(nil)

	r := httptest.NewRequest("DELETE", uri, bytes.NewBuffer(reqBytes))
	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	r.Header.Add("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

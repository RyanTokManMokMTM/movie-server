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
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreatePost(t *testing.T) {
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

	uri := "/api/v1/post"

	var userID uint = 1
	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type postData struct {
		PostTitle string `json:"post_title"`
		PostDesc  string `json:"post_desc"`
		MovieID   uint   `json:"movie_id"`
	}

	reqData := postData{
		PostTitle: "test_post_data",
		PostDesc:  "test_post_desc",
		MovieID:   1,
	}

	reqBytes, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().CreateNewPost(gomock.Any(), &models.Post{
		PostTitle:   reqData.PostTitle,
		PostDesc:    reqData.PostDesc,
		MovieInfoId: reqData.MovieID,
		UserId:      userID,
	}).Times(1).Return(nil)

	r := httptest.NewRequest("POST", uri, bytes.NewBuffer(reqBytes))
	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestDeletePost(t *testing.T) {
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
	uri := "/api/v1/post"

	type PostData struct {
		PostID uint `json:"post_id"`
	}

	reqData := PostData{
		PostID: 1,
	}

	reqByte, _ := json.Marshal(reqData)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOnePostInfoWithUserLiked(gomock.Any(), reqData.PostID, userID).Times(1).Return(&models.Post{
		PostId: reqData.PostID,
	}, nil)
	daoMock.EXPECT().DeletePost(gomock.Any(), reqData.PostID, userID).Times(1).Return(nil)

	r := httptest.NewRequest("DELETE", uri, bytes.NewBuffer(reqByte))
	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckPost(t *testing.T) {
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

	var postID uint = 1
	uri := fmt.Sprintf("/api/v1/post/check/%d", postID)

	daoMock.EXPECT().FindOnePostInfo(gomock.Any(), postID).Times(1).Return(&models.Post{
		PostId: postID,
	}, nil)

	r := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCountAllUserPost(t *testing.T) {
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
	uri := fmt.Sprintf("/api/v1/posts/count/%d", userID)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().CountUserPosts(gomock.Any(), userID).Times(1).Return(int64(0), nil)

	r := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllPost(t *testing.T) {
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
	uri := "/api/v1/posts/all"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)

	dummyPostData := []*models.Post{
		{
			PostId:      1,
			PostTitle:   "test_title_1",
			PostDesc:    "test_desc_1",
			MovieInfoId: 1,
			UserId:      1,
		},
		{
			PostId:      2,
			PostTitle:   "test_title_2",
			PostDesc:    "test_desc_2",
			MovieInfoId: 4,
			UserId:      1,
		},
		{
			PostId:      3,
			PostTitle:   "test_title_3",
			PostDesc:    "test_desc_3",
			MovieInfoId: 6,
			UserId:      1,
		},
	}

	daoMock.EXPECT().
		FindAllPosts(gomock.Any(), userID, gomock.Any(), gomock.Any()).
		Times(1).
		Return(dummyPostData, int64(0), nil)

	r := httptest.NewRequest("GET", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetFollowingPost(t *testing.T) {
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
	uri := "/api/v1/posts/follow"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	dummyPostData := []*models.Post{
		{
			PostId:      1,
			PostTitle:   "test_title_1",
			PostDesc:    "test_desc_1",
			MovieInfoId: 1,
			UserId:      userID,
		},
		{
			PostId:      2,
			PostTitle:   "test_title_2",
			PostDesc:    "test_desc_2",
			MovieInfoId: 4,
			UserId:      userID,
		},
		{
			PostId:      3,
			PostTitle:   "test_title_3",
			PostDesc:    "test_desc_3",
			MovieInfoId: 6,
			UserId:      userID,
		},
	}

	daoMock.EXPECT().
		FindUserByID(gomock.Any(), userID).
		Times(1).
		Return(&models.User{
			ID: userID,
		}, nil)

	daoMock.EXPECT().
		FindFollowingPosts(gomock.Any(), userID, gomock.Any(), gomock.Any()).Times(1).
		Return(dummyPostData, int64(0), nil)

	r := httptest.NewRequest("GET", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetPostByPostID(t *testing.T) {
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
	uri := fmt.Sprintf("/api/v1/post/%d", postID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)

	daoMock.EXPECT().FindOnePostInfoWithUserLiked(gomock.Any(), postID, userID).Times(1).Return(&models.Post{
		PostId:      postID,
		PostTitle:   "test_post",
		PostDesc:    "test_desc",
		UserId:      userID,
		MovieInfoId: 1,
	}, nil)

	r := httptest.NewRequest("GET", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}
func TestGetUserPosts(t *testing.T) {
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
	var reqUserID uint = 2

	//get reqUserID posts -> current user is userID
	uri := fmt.Sprintf("/api/v1/posts/%d", reqUserID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	dummyPostData := []*models.Post{
		{
			PostId:      1,
			PostTitle:   "test_title_1",
			PostDesc:    "test_desc_1",
			MovieInfoId: 1,
			UserId:      reqUserID,
		},
		{
			PostId:      2,
			PostTitle:   "test_title_2",
			PostDesc:    "test_desc_2",
			MovieInfoId: 4,
			UserId:      reqUserID,
		},
		{
			PostId:      3,
			PostTitle:   "test_title_3",
			PostDesc:    "test_desc_3",
			MovieInfoId: 6,
			UserId:      reqUserID,
		},
	}

	daoMock.EXPECT().FindUserByID(gomock.Any(), reqUserID).Times(1).Return(&models.User{
		ID: reqUserID,
	}, nil)

	daoMock.EXPECT().
		FindUserPosts(gomock.Any(), reqUserID, userID, gomock.Any(), gomock.Any()).Times(1).
		Return(dummyPostData, int64(len(dummyPostData)), nil)

	r := httptest.NewRequest("GET", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestUpdatePost(t *testing.T) {
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
	uri := "/api/v1/posts"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type UpdatePost struct {
		PostID    uint   `json:"post_id"`
		PostTitle string `json:"post_title"`
		PostDesc  string `json:"post_desc"`
	}

	reqData := UpdatePost{
		PostID:    postID,
		PostTitle: "Test_title",
		PostDesc:  "Test_Desc",
	}

	reqByte, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOnePostInfo(gomock.Any(), postID).Times(1).Return(&models.Post{
		PostId:      postID,
		PostTitle:   "post_test_1",
		PostDesc:    "post_desc_1",
		UserId:      userID,
		MovieInfoId: 1,
	}, nil)

	daoMock.EXPECT().UpdatePostInfo(gomock.Any(), gomock.Any()).Times(1).Return(nil)

	r := httptest.NewRequest("PATCH", uri, bytes.NewBuffer(reqByte))
	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	r.Header.Add("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

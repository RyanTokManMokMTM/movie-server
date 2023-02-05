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

func TestCountCollectedMovie(t *testing.T) {
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
	uri := fmt.Sprintf("/api/v1/list/movies/count/%d", userID)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().CountCollectedMovie(gomock.Any(), userID).Times(1).Return(int64(0), nil)

	r := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateCustomList(t *testing.T) {
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
	uri := "/api/v1/list"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type createList struct {
		Title string `json:"title"`
		Intro string `json:"intro"`
	}

	reqData := createList{
		Title: "test",
		Intro: "test",
	}

	reqBytes, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().CreateNewList(gomock.Any(), gomock.Any(), gomock.Any(), userID).Times(1).Return(&models.List{
		UserId: userID,
	}, nil)

	r := httptest.NewRequest("POST", uri, bytes.NewBuffer(reqBytes))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteCustomList(t *testing.T) {
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
	var listID uint = 1
	uri := "/api/v1/list"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type deleteList struct {
		ID uint `json:"list_id"`
	}

	reqData := deleteList{
		ID: listID,
	}

	reqBytes, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneList(gomock.Any(), listID).Times(1).Return(&models.List{
		ListId: listID,
		UserId: userID,
	}, nil)
	daoMock.EXPECT().DeleteList(gomock.Any(), listID, userID).Times(1).Return(nil)

	r := httptest.NewRequest("DELETE", uri, bytes.NewBuffer(reqBytes))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllUserList(t *testing.T) {
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

	uri := fmt.Sprintf("/api/v1/lists/%d", userID)

	dummyList := []*models.List{
		{
			ListId:    1,
			ListTitle: "Test",
			ListIntro: "Test",
			UserId:    userID,
		},
		{
			ListId:    2,
			ListTitle: "Test2",
			ListIntro: "Test2",
			UserId:    userID,
		},
		{
			ListId:    3,
			ListTitle: "Test3",
			ListIntro: "Test3",
			UserId:    userID,
		},
	}
	daoMock.EXPECT().FindUserLists(gomock.Any(), userID, gomock.Any(), gomock.Any()).Times(1).Return(dummyList, int64(len(dummyList)), nil)
	daoMock.EXPECT().FindListMovies(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(len(dummyList)).Return([]models.ListMovieInfoWithCreateTime{
		{
			Id: 1,
		},
		{
			Id: 2,
		},
	}, int64(2), nil)

	r := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetListByID(t *testing.T) {
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

	//var userID uint = 1
	var listID uint = 1
	uri := fmt.Sprintf("/api/v1/list/%d", listID)

	daoMock.EXPECT().FindOneList(gomock.Any(), listID).Times(1).Return(&models.List{
		ListId: listID,
	}, nil)

	r := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetListMovies(t *testing.T) {
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

	//var userID uint = 1
	var listID uint = 1
	uri := fmt.Sprintf("/api/v1/list/movies/%d", listID)

	daoMock.EXPECT().FindOneList(gomock.Any(), listID).Times(1).Return(&models.List{
		ListId: listID,
	}, nil)
	daoMock.EXPECT().FindListMovies(gomock.Any(), listID, gomock.Any(), gomock.Any()).Times(1).Return([]models.ListMovieInfoWithCreateTime{}, int64(0), nil)

	r := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetOneMovieFromUserList(t *testing.T) {
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
	var movieID uint = 1
	uri := fmt.Sprintf("/api/v1/list/movie/%d", movieID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneMovieFormAnyList(gomock.Any(), movieID, userID).Times(1).Return(&models.ListMovie{
		MovieInfoId: movieID,
	}, nil)

	r := httptest.NewRequest("GET", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInsertMovieToList(t *testing.T) {
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
	var movieID uint = 1
	var listID uint = 1
	uri := fmt.Sprintf("/api/v1/list/%d/movie/%d", listID, movieID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneUserList(gomock.Any(), listID, userID).Times(1).Return(&models.List{
		ListId: listID,
	}, nil)
	daoMock.EXPECT().FindOneMovie(gomock.Any(), movieID).Times(1).Return(&models.MovieInfo{
		Id: movieID,
	}, nil)
	daoMock.EXPECT().FindOneMovieFormAnyList(gomock.Any(), movieID, userID).Times(1).Return(nil, gorm.ErrRecordNotFound)
	daoMock.EXPECT().InsertMovieToList(gomock.Any(), movieID, listID, userID).Times(1).Return(nil)

	r := httptest.NewRequest("POST", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRemoveListMovies(t *testing.T) {
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
	var listID uint = 1
	uri := fmt.Sprintf("/api/v1/list/movies/%d", listID)

	type removeListMovies struct {
		MovieIds []uint `json:"movie_ids"`
	}

	reqData := removeListMovies{
		MovieIds: []uint{1, 2, 3, 4},
	}

	reqBytes, _ := json.Marshal(reqData)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneList(gomock.Any(), listID).Times(1).Return(&models.List{
		ListId: listID,
	}, nil)
	daoMock.EXPECT().RemoveMoviesFromList(gomock.Any(), gomock.Any(), listID, userID).Times(1).Return(nil)

	r := httptest.NewRequest("DELETE", uri, bytes.NewBuffer(reqBytes))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestRemoveMovieFromList(t *testing.T) {
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
	var movieID uint = 1
	var listID uint = 1
	uri := fmt.Sprintf("/api/v1/list/%d/movie/%d", listID, movieID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)

	daoMock.EXPECT().RemoveMovieFromList(gomock.Any(), movieID, listID, userID).Times(1).Return(nil)

	r := httptest.NewRequest("DELETE", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateCustomList(t *testing.T) {
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
	var listID uint = 1
	uri := "/api/v1/lists"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type updateList struct {
		ID    uint   `json:"list_id"`
		Title string `json:"title"`
		Intro string `json:"intro"`
	}

	reqData := updateList{
		ID:    listID,
		Title: "test",
		Intro: "test",
	}

	reqByte, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneList(gomock.Any(), listID).Times(1).Return(&models.List{
		ListId: listID,
	}, nil)
	daoMock.EXPECT().UpdateList(gomock.Any(), gomock.Any()).Times(1).Return(nil)

	r := httptest.NewRequest("PATCH", uri, bytes.NewBuffer(reqByte))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()
	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

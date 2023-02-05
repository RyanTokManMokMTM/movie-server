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

func TestAcceptFriendRequest(t *testing.T) {
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
	var requestID uint = 1
	uri := "/api/v1/friend/request/accept"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type acceptRequest struct {
		RequestID uint `json:"request_id"`
	}

	reqData := acceptRequest{
		RequestID: requestID,
	}

	reqByte, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneFriendNotificationByID(gomock.Any(), requestID).Times(1).Return(&models.FriendNotification{
		ID: requestID,
	}, nil)
	daoMock.EXPECT().AcceptFriendNotification(gomock.Any(), gomock.Any()).Times(1).Return(nil)

	r := httptest.NewRequest("POST", uri, bytes.NewBuffer(reqByte))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddFriend(t *testing.T) {
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
	var friendID uint = 2
	uri := "/api/v1/friend"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type addFriend struct {
		UserID uint `json:"user_id"`
	}

	reqData := addFriend{
		UserID: friendID,
	}

	reqByte, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindUserByID(gomock.Any(), friendID).Times(1).Return(&models.User{
		ID: friendID,
	}, nil)

	daoMock.EXPECT().IsFriend(gomock.Any(), userID, friendID).Times(1).Return(false, nil)
	daoMock.
		EXPECT().
		FindOneFriendNotification(gomock.Any(), userID, friendID).
		Times(1).
		Return(nil, gorm.ErrRecordNotFound)

	daoMock.
		EXPECT().
		InsertOneFriendNotification(gomock.Any(), userID, gomock.Any()).Times(1).
		Return(uint(1), nil)

	r := httptest.NewRequest("POST", uri, bytes.NewBuffer(reqByte))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCancelFriendRequest(t *testing.T) {
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
	var requestID uint = 1
	uri := "/api/v1/friend/request/cancel"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type cancelFriendRequest struct {
		RequestID uint `json:"request_id"`
	}

	reqData := cancelFriendRequest{
		RequestID: requestID,
	}

	reqByte, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneFriendNotificationByID(gomock.Any(), requestID).Times(1).Return(&models.FriendNotification{}, nil)
	daoMock.EXPECT().CancelFriendNotification(gomock.Any(), requestID).Times(1).Return(nil)

	r := httptest.NewRequest("PATCH", uri, bytes.NewBuffer(reqByte))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeclineFriendRequest(t *testing.T) {
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
	var requestID uint = 1
	uri := "/api/v1/friend/request/decline"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type declineFriend struct {
		RequestID uint `json:"request_id"`
	}

	reqData := declineFriend{
		RequestID: requestID,
	}

	reqBytes, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneFriendNotificationByID(gomock.Any(), requestID).Times(1).Return(&models.FriendNotification{}, nil)
	daoMock.EXPECT().DeclineFriendNotification(gomock.Any(), requestID).Times(1).Return(nil)

	r := httptest.NewRequest("PATCH", uri, bytes.NewBuffer(reqBytes))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetFriendRequest(t *testing.T) {
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
	uri := "/api/v1/friend/requests"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	dummyRequest := []*models.FriendNotification{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	}

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)

	daoMock.EXPECT().GetFriendRequest(gomock.Any(), userID, gomock.Any(), gomock.Any()).Times(1).Return(dummyRequest, int64(len(dummyRequest)), nil)

	r := httptest.NewRequest("GET", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestIsFriend(t *testing.T) {
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
	var friendID uint = 2
	uri := fmt.Sprintf("/api/v1/friend/%d", friendID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)

	daoMock.EXPECT().FindUserByID(gomock.Any(), friendID).Times(1).Return(&models.User{
		ID: friendID,
	}, nil)

	daoMock.EXPECT().FindOneFriendNotification(gomock.Any(), userID, friendID).Times(1).Return(&models.FriendNotification{}, nil)

	r := httptest.NewRequest("GET", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRemoveFriend(t *testing.T) {
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
	var friendID uint = 2
	uri := "/api/v1/friend"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type deleteFriend struct {
		FriendID uint `json:"user_id"`
	}

	reqData := deleteFriend{
		FriendID: friendID,
	}

	reqBytes, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)

	daoMock.EXPECT().FindOneFriend(gomock.Any(), userID, friendID).Times(1).Return(&models.User{
		ID: friendID,
	}, nil)

	daoMock.EXPECT().RemoveFriend(gomock.Any(), userID, friendID).Times(1).Return(nil)

	r := httptest.NewRequest("DELETE", uri, bytes.NewBuffer(reqBytes))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

package logic

import (
	"bytes"
	"encoding/json"
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

func TestCreateCommentLikes(t *testing.T) {
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
	var commentID uint = 1
	var commentUserID uint = 2
	uri := "/api/v1/liked/comment"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type commentLikes struct {
		CommentId uint `json:"comment_id"`
	}

	reqData := commentLikes{
		CommentId: commentID,
	}

	reqBytes, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneComment(gomock.Any(), commentID).Times(1).Return(&models.Comment{
		CommentID: commentID,
		Comment:   "test",
		UserID:    commentUserID,
	}, nil)
	daoMock.EXPECT().InsertOneCommentLike(gomock.Any(), userID, commentID, gomock.Any()).Times(1).Return(nil)

	//TODO: send notification test
	daoMock.EXPECT().FindOneLikeCommentNotification(gomock.Any(), commentUserID, userID, commentID).Times(1).Return(gorm.ErrRecordNotFound)
	daoMock.EXPECT().InsertOneCommentLikeNotification(gomock.Any(), gomock.Any(), userID, commentID, commentUserID, gomock.Any()).Times(1).Return(nil)

	r := httptest.NewRequest("POST", uri, bytes.NewBuffer(reqBytes))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRemoveCommentLikes(t *testing.T) {
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
	var commentID uint = 1
	uri := "/api/v1/liked/comment"

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type removeComment struct {
		CommentId uint `json:"comment_id"`
	}

	reqData := removeComment{
		CommentId: commentID,
	}

	reqBytes, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneComment(gomock.Any(), commentID).Times(1).Return(&models.Comment{
		CommentID: commentID,
		Comment:   "test",
	}, nil)
	daoMock.EXPECT().RemoveOneCommentLike(gomock.Any(), userID, commentID, gomock.Any()).Times(1).Return(nil)

	r := httptest.NewRequest("PATCH", uri, bytes.NewBuffer(reqBytes))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

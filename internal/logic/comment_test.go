package logic

import (
	"bytes"
	"database/sql"
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

func TestCreateComment(t *testing.T) {
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
	var postUserID uint = 2

	uri := fmt.Sprintf("/api/v1/comments/%d", postID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type createComment struct {
		Comment string `json:"comment"`
	}

	var reqData = createComment{
		Comment: "test_comment",
	}

	reqByte, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOnePostInfo(gomock.Any(), postID).Times(1).Return(&models.Post{
		PostId:      postID,
		PostTitle:   "test_title",
		PostDesc:    "test_desc",
		UserId:      postUserID,
		MovieInfoId: 1,
	}, nil)
	daoMock.EXPECT().CreatePostComment(gomock.Any(), userID, postID, reqData.Comment).Times(1).Return(&models.Comment{
		Comment: reqData.Comment,
	}, nil)

	//TODO: Send notification test
	daoMock.EXPECT().
		InsertOneCommentNotification(gomock.Any(), postUserID, userID, postID, gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	r := httptest.NewRequest("POST", uri, bytes.NewBuffer(reqByte))
	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestCreateReplyComment(t *testing.T) {
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
	var commentUserID uint = 2
	var postID uint = 1
	var commentID uint = 1

	uri := fmt.Sprintf("/api/v1/comments/%d/reply/%d", postID, commentID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type replyComment struct {
		ParentCommentID uint   `json:"parent_id"`
		Comment         string `json:"comment"`
	}

	reqData := replyComment{
		ParentCommentID: 1,
		Comment:         "test_comment",
	}

	reqBody, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOnePostInfo(gomock.Any(), postID).Times(1).Return(&models.Post{
		PostId:      postID,
		PostTitle:   "test_title",
		PostDesc:    "test_desc",
		UserId:      userID,
		MovieInfoId: 1,
	}, nil)
	daoMock.EXPECT().FindOneComment(gomock.Any(), commentID).Times(1).Return(&models.Comment{
		CommentID: commentID,
		Comment:   "test_comment",
		UserID:    commentUserID,
		PostID:    postID,
	}, nil)
	daoMock.
		EXPECT().
		CreatePostReplyComment(gomock.Any(), userID, postID, commentID, gomock.Any(), commentUserID, gomock.Any()).Times(1).
		Return(&models.Comment{
			PostID:  postID,
			Comment: reqData.Comment,
		}, nil)

	//TODO: Notification
	daoMock.EXPECT().InsertOneReplyCommentNotification(gomock.Any(), commentUserID, userID, postID, gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)

	r := httptest.NewRequest("POST", uri, bytes.NewBuffer(reqBody))
	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestDeleteComment(t *testing.T) {
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
	uri := "/api/v1/comments"

	type deleteComment struct {
		CommentID uint `json:"comment_id"`
	}

	reqData := deleteComment{
		CommentID: commentID,
	}

	reqByte, _ := json.Marshal(reqData)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().DeleteComment(gomock.Any(), commentID).Times(1).Return(nil)

	r := httptest.NewRequest("DELETE", uri, bytes.NewBuffer(reqByte))
	r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)
}

func TestGetPostComment(t *testing.T) {
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
	uri := fmt.Sprintf("/api/v1/comments/%d", postID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOnePostInfo(gomock.Any(), postID).Times(1).Return(&models.Post{
		PostId:      postID,
		PostTitle:   "title_test",
		PostDesc:    "desc_test",
		UserId:      userID,
		MovieInfoId: 1,
	}, nil)

	dummyComment := []*models.Comment{
		{
			CommentID: 1,
			Comment:   "test_1",
			PostID:    1,
			UserID:    1,
		},
		{
			CommentID: 2,
			Comment:   "test_2",
			PostID:    1,
			UserID:    2,
		},
		{
			CommentID: 3,
			Comment:   "test_3",
			PostID:    1,
			UserID:    3,
			ParentID:  sql.NullInt64{Int64: 1, Valid: true},
		},
	}
	daoMock.EXPECT().FindPostComments(gomock.Any(), postID, userID, gomock.Any(), gomock.Any()).Times(1).Return(dummyComment, int64(len(dummyComment)), nil)

	r := httptest.NewRequest("GET", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetReplyComment(t *testing.T) {
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
	var parentCommentID uint = 1
	uri := fmt.Sprintf("/api/v1/comments/reply/%d", parentCommentID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneComment(gomock.Any(), parentCommentID).Times(1).Return(&models.Comment{
		CommentID: parentCommentID,
	}, nil)

	dummyComment := []*models.Comment{
		{
			CommentID: 1,
			Comment:   "test_1",
			PostID:    1,
			UserID:    1,
		},
		{
			CommentID: 2,
			Comment:   "test_2",
			PostID:    1,
			UserID:    2,
		},
		{
			CommentID: 3,
			Comment:   "test_3",
			PostID:    1,
			UserID:    3,
			ParentID:  sql.NullInt64{Int64: 1, Valid: true},
		},
	}
	daoMock.EXPECT().FindReplyComments(gomock.Any(), parentCommentID, userID, gomock.Any(), gomock.Any()).Times(1).Return(dummyComment, int64(len(dummyComment)), nil)

	r := httptest.NewRequest("GET", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateComment(t *testing.T) {
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
	uri := fmt.Sprintf("/api/v1/comments/%d", commentID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	type updateComment struct {
		Comment string `json:"comment"`
	}

	reqData := updateComment{
		Comment: "test_update",
	}

	reqBody, _ := json.Marshal(reqData)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneComment(gomock.Any(), commentID).Times(1).Return(&models.Comment{
		CommentID: commentID,
	}, nil)
	daoMock.EXPECT().UpdateComment(gomock.Any(), gomock.Any()).Times(1).Return(nil)

	r := httptest.NewRequest("PATCH", uri, bytes.NewBuffer(reqBody))
	r.Header.Add("Content-Type", "application/json; charset=UTF-8")
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

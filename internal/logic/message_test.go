package logic

import (
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

func TestGetRoomMessage(t *testing.T) {
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
	var roomMemberID uint = 2
	var roomID uint = 1
	var lastMsgID uint = 1
	uri := fmt.Sprintf("/api/v1/message/%d/%d", roomID, lastMsgID)

	token, err := tokenGen(c.Auth.AccessExpire, userID, c.Auth.AccessSecret, time.Now())
	assert.Nil(t, err)

	daoMock.EXPECT().FindUserByID(gomock.Any(), userID).Times(1).Return(&models.User{
		ID: userID,
	}, nil)
	daoMock.EXPECT().FindOneRoomMember(gomock.Any(), roomID, userID).Times(1).Return(&models.User{
		ID: roomMemberID,
	}, nil)

	dummyMessage := []*models.Message{
		{
			ID:      1,
			Content: "test",
			RoomID:  roomID,
			Sender:  userID,
		},
		{
			ID:      2,
			Content: "test2",
			RoomID:  roomID,
			Sender:  userID,
		},
		{
			ID:      3,
			Content: "test3",
			RoomID:  roomID,
			Sender:  roomMemberID,
		},
	}
	daoMock.EXPECT().
		GetRoomMessage(gomock.Any(), roomID, lastMsgID, gomock.Any(), gomock.Any()).Times(1).
		Return(dummyMessage, int64(100), nil)

	r := httptest.NewRequest("GET", uri, nil)
	r.Header.Add("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	ser := server.SetUpEngine(c, daoMock)
	ser.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/mock/gomock"
	"github.com/ryantokmanmokmtm/movie-server/common/crytox"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/jwtx"
	"github.com/ryantokmanmokmtm/movie-server/common/util"
	"github.com/ryantokmanmokmtm/movie-server/internal/config"
	mockdb "github.com/ryantokmanmokmtm/movie-server/internal/dao/mock"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"github.com/ryantokmanmokmtm/movie-server/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestSignUp(t *testing.T) {
	ctl := gomock.NewController(t)
	daoMock := mockdb.NewMockStore(ctl)
	defer ctl.Finish()

	var confData = []byte(`{
		"Path" : "./resources",
		"Auth":{
				"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
				"AccessExpire": 86400
		},
		"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
		"MaxBytes": 1073741824
	}`)

	var c config.Config
	err := json.Unmarshal(confData, &c)
	assert.Nil(t, err)

	uri := "/api/v1/user/signup"
	password := "admin12345"
	user := &models.User{
		Name:     "jacksontmm",
		Email:    "admin@admin.com",
		Password: crytox.PasswordEncrypt(password, c.Salt),
		Avatar:   "/defaultAvatar.jpeg",
		Cover:    "/defaultCover.jpeg",
	}

	testCases := []struct {
		Name     string
		PostData []byte
		MockTest func(store *mockdb.MockStore)
		Response func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "Success",
			PostData: []byte(`{
				"name" : "jacksontmm",
				"email" : "admin@admin.com",
				"password" : "admin12345"
			}`),
			MockTest: func(store *mockdb.MockStore) {
				store.EXPECT().FindUserByEmail(gomock.Any(), user.Email).Times(1).Return(nil, gorm.ErrRecordNotFound)
				store.EXPECT().CreateUser(gomock.Any(), user).Times(1).Return(user, nil)
			},
			Response: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			Name: "User Existed",
			PostData: []byte(`{
				"name" : "jacksontmm",
				"email" : "admin@admin.com",
				"password" : "admin12345"
			}`),
			MockTest: func(store *mockdb.MockStore) {
				store.EXPECT().FindUserByEmail(gomock.Any(), user.Email).Times(1).Return(&models.User{ID: 1}, nil)
			},
			Response: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusConflict, recorder.Code)
			},
		},
		{
			Name: "Email error",
			PostData: []byte(`{
				"name" : "jacksontmm",
				"email" : "adminadmin.com",
				"password" : "admin12345"
			}`),
			MockTest: func(store *mockdb.MockStore) {
				//Do nothing cuz return before handler
				return
			},
			Response: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "password too short",
			PostData: []byte(`{
				"name" : "jacksontmm",
				"email" : "admin@admin.com",
				"password" : "1"
			}`),
			MockTest: func(store *mockdb.MockStore) {
				return
			},
			Response: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "password too long",
			PostData: []byte(`{
				"name" : "jacksontmm",
				"email" : "admin@admin.com",
				"password" : "testingsignuppasswordover32characterlong"
			}`),
			MockTest: func(store *mockdb.MockStore) {
				return
			},
			Response: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "missing field - name",
			PostData: []byte(`{
				"email" : "admin@admin.com",
				"password" : "admin123465"
			}`),
			MockTest: func(store *mockdb.MockStore) {
				return
			},
			Response: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "missing field - email",
			PostData: []byte(`{
				"name" : "jacksontmm",
				"password" : "admin12345"
			}`),
			MockTest: func(store *mockdb.MockStore) {
				return
			},
			Response: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "missing field - password",
			PostData: []byte(`{
				"name" : "jacksontmm",
				"email" : "admin@admin.com",
			}`),
			MockTest: func(store *mockdb.MockStore) {
				return
			},
			Response: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			test.MockTest(daoMock)

			ser := server.SetUpEngine(c, daoMock)

			r, err := http.NewRequest("POST", uri, bytes.NewBuffer(test.PostData))
			r.Header.Set("Content-Type", "application/json; charset=UTF-8")
			assert.Nil(t, err)

			w := httptest.NewRecorder()
			ser.ServeHTTP(w, r)

			test.Response(t, w)
		})
	}

}

func TestSignIn(t *testing.T) {
	ctl := gomock.NewController(t)
	daoMock := mockdb.NewMockStore(ctl)
	defer ctl.Finish()

	var configInfo = []byte(`{
	"Path" : "./resources",
	"Auth":{
			"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
			"AccessExpire": 86400
	},
	"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
	"MaxBytes": 1073741824
	}`)

	uri := "/api/v1/user/login"

	var c config.Config
	err := json.Unmarshal(configInfo, &c)
	assert.Nil(t, err)

	password := "admin12345"
	email := "admin@admin.com"
	user := &models.User{
		ID:       uint(rand.Intn(10000)),
		Name:     gofakeit.Name(),
		Email:    email,
		Password: crytox.PasswordEncrypt(password, c.Salt),
		Avatar:   "/defaultAvatar.jpeg",
		Cover:    "/defaultCover.jpeg",
	}
	testCases := []struct {
		Name      string
		Data      []byte
		MockTest  func(mock *mockdb.MockStore)
		CheckResp func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "Success",
			Data: []byte(`{
			"email":"admin@admin.com",
			"password":"admin12345"
			}`),
			MockTest: func(mock *mockdb.MockStore) {
				mock.EXPECT().FindUserByEmail(gomock.Any(), user.Email).Times(1).Return(user, nil)
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			Name: "password incorrect",
			Data: []byte(`{
			"email":"admin@admin.com",
			"password":"admin123456"
			}`),
			MockTest: func(mock *mockdb.MockStore) {
				mock.EXPECT().FindUserByEmail(gomock.Any(), user.Email).Times(1).Return(user, nil)
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "User Not Found",
			Data: []byte(`{
			"email":"admin@admin.com",
			"password":"admin12345"
			}`),
			MockTest: func(mock *mockdb.MockStore) {
				mock.EXPECT().FindUserByEmail(gomock.Any(), user.Email).Times(1).Return(nil, gorm.ErrRecordNotFound)
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			Name: "Missing Email Field",
			Data: []byte(`{
			"password":"admin12345"
			}`),
			MockTest: func(mock *mockdb.MockStore) {
				return
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "Missing Password Field",
			Data: []byte(`{
			"email":"admin@admin.com"
			}`),
			MockTest: func(mock *mockdb.MockStore) {
				return
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			Name: "password Too long",
			Data: []byte(`{
			"email":"admin@admin.com",
			"password":"asdasdhjkasdkajskdjaskldjlasdjlkasjdklasjkldasjkldasjkld"
			}`),
			MockTest: func(mock *mockdb.MockStore) {
				return
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "password Too short",
			Data: []byte(`{
			"email":"admin@admin.com",
			"password":"asda"
			}`),
			MockTest: func(mock *mockdb.MockStore) {
				return
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "password Too short",
			Data: []byte(`{
			"email":"admin@admin.com",
			"password":"asda"
			}`),
			MockTest: func(mock *mockdb.MockStore) {
				return
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "email invalid",
			Data: []byte(`{
			"email":"admin.com",
			"password":"admin12345"
			}`),
			MockTest: func(mock *mockdb.MockStore) {
				return
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			Name: "email too long",
			Data: []byte(`{
			"email":"admin@testingemailtoolongover32characterlong.com",
			"password":"admin12345"
			}`),
			MockTest: func(mock *mockdb.MockStore) {
				return
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			test.MockTest(daoMock)
			r := httptest.NewRequest("POST", uri, bytes.NewBuffer(test.Data))
			r.Header.Set("Content-Type", "application/json; charset=UTF-8")

			w := httptest.NewRecorder()
			ser := server.SetUpEngine(c, daoMock)
			ser.ServeHTTP(w, r)

			test.CheckResp(t, w)
		})
	}

}

func TestUpdateUserProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	daoMock := mockdb.NewMockStore(ctl)
	defer ctl.Finish()

	uri := "/api/v1/user/profile"
	user := util.RandomUser()
	testCases := []struct {
		Name      string
		Data      []byte
		Token     func(exp int64, key string) (string, error)
		MockTest  func(mock *mockdb.MockStore)
		CheckResp func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "Success",
			Data: []byte(`{
				"name":"jackson"
			}`),
			Token: func(exp int64, key string) (string, error) {
				return tokenGen(exp, user.ID, key, time.Now())
			},
			MockTest: func(mock *mockdb.MockStore) {
				daoMock.EXPECT().FindUserByID(gomock.Any(), user.ID).Times(1).Return(&user, nil)
				daoMock.EXPECT().UpdateUser(gomock.Any(), user.ID, gomock.Any()).Times(1).Return(nil)

			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			Name: "Token Expired",
			Data: []byte(`{
				"name":"jackson"
			}`),
			Token: func(exp int64, key string) (string, error) {
				return tokenGen(int64(48*time.Hour), user.ID, key, time.Now().Add(-100*time.Hour))
			},
			MockTest: func(mock *mockdb.MockStore) {
				return
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			Name: "Field is missing",
			Data: []byte(`{}`),
			Token: func(exp int64, key string) (string, error) {
				return tokenGen(exp, user.ID, key, time.Now())
			},
			MockTest: func(mock *mockdb.MockStore) {
				return
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	//temp config data
	var confData = []byte(`{
			"Path" : "./resources",
			"Auth":{
					"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
					"AccessExpire": 86400
			},
			"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
			"MaxBytes": 1073741824
	}`)

	var c config.Config
	err := json.Unmarshal(confData, &c)
	assert.Nil(t, err)

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			test.MockTest(daoMock)
			token, err := test.Token(c.Auth.AccessExpire, c.Auth.AccessSecret)
			assert.Nil(t, err)

			r := httptest.NewRequest("PATCH", uri, bytes.NewBuffer(test.Data))
			r.Header.Add("Content-Type", "application/json; charset=UTF-8")
			r.Header.Add("Authorization", "Bearer "+token)
			w := httptest.NewRecorder()

			ser := server.SetUpEngine(c, daoMock)
			ser.ServeHTTP(w, r)
			test.CheckResp(t, w)
		})
	}

}

func TestGetUserProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	daoMock := mockdb.NewMockStore(ctl)

	var confData = []byte(`{
				"Path" : "./resources",
				"Auth":{
						"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
						"AccessExpire": 60
				},
				"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
				"MaxBytes": 1073741824,
				"log":{
					"Encoding":"plain"
				}
			}`)

	var c config.Config
	err := json.Unmarshal(confData, &c)
	assert.Nil(t, err)

	testCases := []struct {
		Name      string
		UserID    uint
		Token     func(id uint) (string, error)
		MockTest  func(mock *mockdb.MockStore, id uint)
		CheckResp func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name:   "Success",
			UserID: 1,
			Token: func(id uint) (string, error) {
				return tokenGen(c.Auth.AccessExpire, id, c.Auth.AccessSecret, time.Now())
			},
			MockTest: func(mock *mockdb.MockStore, id uint) {
				mock.EXPECT().FindUserByID(gomock.Any(), id).Times(1).Return(&models.User{
					ID: id,
				}, nil)
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			Name:   "Token expired",
			UserID: 2,
			Token: func(id uint) (string, error) {
				//49 hours ago and expired at 1 hour ago //-49+48=1
				return tokenGen(int64(48*time.Hour), id, c.Auth.AccessSecret, time.Now().Add(-49*time.Hour))
			},
			MockTest: func(mock *mockdb.MockStore, id uint) {
				return
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//t.Error(recorder.Code)
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	uri := "/api/v1/user/profile"

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			token, err := test.Token(test.UserID)
			assert.Nil(t, err)

			//t.Log(token)
			test.MockTest(daoMock, test.UserID)
			emptyBuffer := bytes.NewReader([]byte{})
			r := httptest.NewRequest("GET", uri, emptyBuffer) //for read data  and write to response
			r.Header.Add("Authorization", "Bearer "+token)

			w := httptest.NewRecorder()

			ser := server.SetUpEngine(c, daoMock)
			ser.ServeHTTP(w, r)

			test.CheckResp(t, w)
		})
	}

}

func TestGetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	//user := randomUser()
	daoMock := mockdb.NewMockStore(ctl)

	testCases := []struct {
		Name      string
		ID        uint
		MockTest  func(store *mockdb.MockStore, id uint)
		CheckResp func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "Success",
			ID:   1,
			MockTest: func(store *mockdb.MockStore, id uint) {
				daoMock.EXPECT().FindUserByID(gomock.Any(), id).Times(1).Return(&models.User{ID: id}, nil)
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			Name: "User Not Found",
			ID:   2,
			MockTest: func(store *mockdb.MockStore, id uint) {
				daoMock.EXPECT().FindUserByID(gomock.Any(), id).Times(1).Return(nil, gorm.ErrRecordNotFound)
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}
	//temp config data

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			var confData = []byte(`{
			"Path" : "./resources",
			"Auth":{
					"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
					"AccessExpire": 86400
			},
			"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
			"MaxBytes": 1073741824
			}`)

			var c config.Config
			err := json.Unmarshal(confData, &c)
			assert.Nil(t, err)

			uri := fmt.Sprintf("/api/v1/user/info/%d", test.ID)
			test.MockTest(daoMock, test.ID)

			r := httptest.NewRequest("GET", uri, nil)
			w := httptest.NewRecorder()

			ser := server.SetUpEngine(c, daoMock)
			ser.ServeHTTP(w, r)

			test.CheckResp(t, w)
		})
	}

}

func TestUpdateUserCover(t *testing.T) {
	ctl := gomock.NewController(t)
	daoMock := mockdb.NewMockStore(ctl)
	defer ctl.Finish()

	var confData = []byte(`{
	"Path" : "../../resources",
	"Auth":{
			"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
			"AccessExpire": 86400
	},
	"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
	"MaxBytes": 1073741824,
	"log":{
		"Encoding":"plain"
	}
	}`)
	user := util.RandomUser()
	var c config.Config
	err := json.Unmarshal(confData, &c)
	assert.Nil(t, err)

	uri := "/api/v1/user/cover"
	tempDataPath := "../../resources/test/testing.txt"

	testCases := []struct {
		Name      string
		Token     func(id uint) (string, error)
		MockTest  func(mock *mockdb.MockStore)
		CheckResp func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "Success",
			Token: func(id uint) (string, error) {
				return tokenGen(c.Auth.AccessExpire, id, c.Auth.AccessSecret, time.Now())
			},
			MockTest: func(mock *mockdb.MockStore) {
				daoMock.EXPECT().FindUserByID(gomock.Any(), user.ID).Times(1).Return(&user, nil)
				daoMock.EXPECT().UpdateUser(gomock.Any(), user.ID, &models.User{Cover: "/testing.txt"}).Times(1).Return(nil)
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			Name: "Token Expired",
			Token: func(id uint) (string, error) {
				//49 hours ago and expired at 1 hour ago //-49+48=1
				return tokenGen(int64(48*time.Hour), id, c.Auth.AccessSecret, time.Now().Add(-49*time.Hour))
			},
			MockTest: func(mock *mockdb.MockStore) {
				return
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//t.Error(recorder.Code)
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			Name: "User Not Found",
			Token: func(id uint) (string, error) {
				return tokenGen(c.Auth.AccessExpire, id, c.Auth.AccessSecret, time.Now())
			},
			MockTest: func(mock *mockdb.MockStore) {
				mock.EXPECT().FindUserByID(gomock.Any(), user.ID).Times(1).Return(nil, gorm.ErrRecordNotFound)
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			test.MockTest(daoMock)
			token, err := test.Token(user.ID)
			assert.Nil(t, err)

			file, err := os.Open(tempDataPath)
			assert.Nil(t, err)

			defer func(file *os.File) {
				err := file.Close()
				assert.Nil(t, err)
			}(file)

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			part, err := writer.CreateFormFile("uploadCover", filepath.Base(tempDataPath))
			assert.Nil(t, err)
			//t.Log(part)

			_, err = io.Copy(part, file)
			assert.Nil(t, err)

			err = writer.Close()
			assert.Nil(t, err)

			r := httptest.NewRequest("PATCH", uri, body)
			r.Header.Add("Authorization", "Bearer "+token)
			r.Header.Add("Content-Type", writer.FormDataContentType())
			w := httptest.NewRecorder()

			ser := server.SetUpEngine(c, daoMock)
			ser.ServeHTTP(w, r)

			test.CheckResp(t, w)
		})
	}

}

func TestUpdateUserAvatar(t *testing.T) {
	ctl := gomock.NewController(t)
	daoMock := mockdb.NewMockStore(ctl)
	defer ctl.Finish()

	var confData = []byte(`{
	"Path" : "../../resources",
	"Auth":{
			"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
			"AccessExpire": 86400
	},
	"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
	"MaxBytes": 1073741824
	}`)
	user := util.RandomUser()
	var c config.Config
	err := json.Unmarshal(confData, &c)
	assert.Nil(t, err)

	uri := "/api/v1/user/avatar"
	tempDataPath := "../../resources/test/testing.txt"

	//t.Log(body.String())
	testCases := []struct {
		Name      string
		Token     func(id uint) (string, error)
		MockTest  func(mock *mockdb.MockStore)
		CheckResp func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			Name: "Success",
			Token: func(id uint) (string, error) {
				return tokenGen(c.Auth.AccessExpire, id, c.Auth.AccessSecret, time.Now())
			},
			MockTest: func(mock *mockdb.MockStore) {
				daoMock.EXPECT().FindUserByID(gomock.Any(), user.ID).Times(1).Return(&user, nil)
				daoMock.EXPECT().UpdateUser(gomock.Any(), user.ID, &models.User{Avatar: "/testing.txt"}).Times(1).Return(nil)
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			Name: "Token Expired",
			Token: func(id uint) (string, error) {
				return tokenGen(int64(48*time.Hour), id, c.Auth.AccessSecret, time.Now().Add(-49*time.Hour))
			},
			MockTest: func(mock *mockdb.MockStore) {
				return

			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			Name: "User Not Found",
			Token: func(id uint) (string, error) {
				return tokenGen(c.Auth.AccessExpire, id, c.Auth.AccessSecret, time.Now())
			},
			MockTest: func(mock *mockdb.MockStore) {
				mock.EXPECT().FindUserByID(gomock.Any(), user.ID).Times(1).Return(nil, gorm.ErrRecordNotFound)
			},
			CheckResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			test.MockTest(daoMock)
			token, err := test.Token(user.ID)
			assert.Nil(t, err)

			file, err := os.Open(tempDataPath)
			assert.Nil(t, err)

			defer func(file *os.File) {
				err := file.Close()
				assert.Nil(t, err)
			}(file)

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			part, err := writer.CreateFormFile("uploadAvatar", filepath.Base(tempDataPath))
			assert.Nil(t, err)
			//t.Log(part)

			_, err = io.Copy(part, file)
			assert.Nil(t, err)

			err = writer.Close()
			assert.Nil(t, err)

			r := httptest.NewRequest("PATCH", uri, body)
			r.Header.Add("Authorization", "Bearer "+token)
			r.Header.Add("Content-Type", writer.FormDataContentType())
			recorder := httptest.NewRecorder()

			ser := server.SetUpEngine(c, daoMock)
			ser.ServeHTTP(recorder, r)

			test.CheckResp(t, recorder)
		})
	}

}

func tokenGen(exp int64, id uint, key string, iat time.Time) (string, error) {
	issues := iat.Unix()
	expired := iat.Add(time.Duration(exp) * time.Second).Unix()

	payload := map[string]interface{}{
		ctxtool.CTXJWTUserID: id,
	}
	return jwtx.GetToken(issues, expired, key, payload)
}

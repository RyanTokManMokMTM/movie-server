package main

import (
	"github.com/golang/mock/gomock"
	mockdb "github.com/ryantokmanmokmtm/movie-server/internal/dao/mock"
	"testing"
)

func Test_microservice(t *testing.T) {
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

	t.Error("testing...")
	return
	//var c config.Config
	//err := json.Unmarshal(configInfo, &c)
	//assert.Nil(t, err)
	//
	//uri := "/api/v1/ping"
	//r := httptest.NewRequest("GET", uri, nil)
	//w := httptest.NewRecorder()
	//
	//server := server.SetUpEngine(c, daoMock)
	//server.ServeHTTP(w, r)
	//assert.Equal(t, http.StatusOK, w.Code)

}

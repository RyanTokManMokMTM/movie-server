package jwtx

import (
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Jwtx(t *testing.T) {
	testCases := []struct {
		Name      string
		AccessKey string
		issuesAt  int64
		expiredAt int64
		payload   map[string]interface{}
	}{
		{
			Name:      "success",
			AccessKey: "XWROIEZHJFtkaUxQRlhjRyZPQy4yNFUhemg3fj1JSURDaFNPaEwleWtrXXpVPQ==",
			issuesAt:  time.Now().Unix(),
			expiredAt: time.Now().Add(time.Second * 60).Unix(),
			payload: map[string]interface{}{
				ctxtool.CTXJWTUserID: 1,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			_, err := GetToken(test.issuesAt, test.expiredAt, test.AccessKey, test.payload)
			assert.Nil(t, err)
		})
	}
}

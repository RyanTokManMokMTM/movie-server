package ctxtool

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetUserIDFromCTX(t *testing.T) {
	testCases := []struct {
		Name     string
		Ctx      context.Context
		Expected uint
	}{
		{
			Name:     "Success",
			Ctx:      context.WithValue(context.Background(), CTXJWTUserID, json.Number("1")),
			Expected: uint(1),
		},
		{
			Name:     "Failed",
			Ctx:      context.Background(),
			Expected: uint(0),
		},
		{
			Name:     "User_ID is not an integer",
			Ctx:      context.WithValue(context.Background(), CTXJWTUserID, json.Number("test")),
			Expected: uint(0),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			u := GetUserIDFromCTX(test.Ctx)
			assert.Equal(t, test.Expected, u)
		})
	}
}

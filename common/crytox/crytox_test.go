package crytox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PasswordEncrypt(t *testing.T) {
	testCases := []struct {
		Name     string
		Password string
		Salt     string
		Expected string
	}{
		{
			Name:     "Test case 1",
			Password: "jacksontmm",
			Salt:     "W4tiDEeWlwxlRPYYRRMhJ65piS1ochvMymwfVdumittPoSxhkHNnVLe6m12C4v15",
			Expected: "f3f46cb53f67fbea768f224781f88c8067c8bbe01203f3395a1c914658dc8109",
		},
		{
			Name:     "Test case 2",
			Password: "admin123465",
			Salt:     "ZUtkd1EhNmdEWGQrQTYpOXRXIENuQGpNTHRsZHRiWDZsWitFKUcxVDJfKSwveG",
			Expected: "934edd3620033ae1bd8b4901d9bcfc2c39c11808f4f7fd26781820bd979c47fd",
		},
		{
			Name:     "Test case 3",
			Password: "tommypogd1234",
			Salt:     "PiYgVUA4fEJPNUdQZ0ZPOC92IStGbFdIN2Anbm1hY1JbP2UiSy52Y2RKIFM3Ki8=",
			Expected: "cf39e5d592744daef7b9a9a3a7c27a02bf3984f7fbcb1dd43230e11358b7046f",
		},
		{
			Name:     "Test case 4",
			Password: "tim66684jimX",
			Salt:     "WGsySjlHPlwtTypPWCNjMDlza2VWYj0tRF10J0x+ZUVeSzVsTVM1TiFuNiBYWS46",
			Expected: "0124db7850d0557ec6c15a5e43670896c316256ad4985412671a5dcdc7fc8e22",
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			pw := PasswordEncrypt(test.Password, test.Salt)
			assert.Equal(t, test.Expected, pw)
		})
	}
}

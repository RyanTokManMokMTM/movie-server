package util

import "testing"

func Test_random(t *testing.T) {
	_ = RandomUser()
}

func Test_RandomUserEncryptPassword(t *testing.T) {
	_ = RandomUserEncryptPassword("testing", "YnVzYWxvbmVpbWFnZWZvcnR5cnVubmluZ3doYWxlY2VydGFpbmx5c2l4aGlkZWhlYXI=")
}

package main

import "testing"

//just a simple test
func Test_main(t *testing.T) {
	t.Run("just a simple testing", func(t *testing.T) {
		t.Log("simple testing..... - always success...")
	})
}

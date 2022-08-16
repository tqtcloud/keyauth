package utils

import "testing"

func TestMakeBearer(t *testing.T) {
	v := MakeBearer(24)
	t.Log(v)
}

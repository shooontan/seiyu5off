package util

import (
	"testing"
)

func TestStr2int(t *testing.T) {
	result := Str2int("10")
	if result != 10 {
		t.Fatal("faild test")
	}

	result = Str2int("01")
	if result != 1 {
		t.Fatal("faild test")
	}
}

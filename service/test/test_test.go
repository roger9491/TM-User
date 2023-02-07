package test

import "testing"

func Test_add_1_2(t *testing.T) {
	if add(1, 2) != 3 {
		t.Error("wrong result")
	}
}

func Test_add_33_55(t *testing.T) {
	if add(33, 55) != 88 {
		t.Error("wrong result")
	}
}

func Test_add_110_99(t *testing.T) {
	if add(110, 99) != 209 {
		t.Error("wrong result")
	}
}

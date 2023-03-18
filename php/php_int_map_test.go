package php

import "testing"

func TestIntInArray(t *testing.T) {
	a := IntInArray(1, []int{2, 6, 7, 11, 3})
	b := IntInArray(1, []int{2, 6, 7, 1, 3})
	if a == true || b == false {
		t.Error("IntInArray error", a, b)
	}
}

func TestUIntInArray(t *testing.T) {
	a := UIntInArray(1, []uint64{2, 6, 7, 11, 3})
	b := UIntInArray(1, []uint64{2, 6, 7, 1, 3})
	if a == true || b == false {
		t.Error("TestUIntInArray error", a, b)
	}
}

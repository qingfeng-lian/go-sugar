package php

func IntInArray(a int, m []int) bool {
	for _, v := range m {
		if v == a {
			return true
		}
	}
	return false
}

func UIntInArray(a uint64, m []uint64) bool {
	for _, v := range m {
		if v == a {
			return true
		}
	}
	return false
}

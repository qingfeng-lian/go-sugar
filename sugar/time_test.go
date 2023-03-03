package sugar

import "testing"

func TestGetTimeUnix(t *testing.T) {
	day := "2023-03-03 12:12:25"
	timeU, err := GetTimeUnix(day, "")
	t.Log(timeU)
	t.Log(err)
}

func TestGetDayFirstSecond(t *testing.T) {
	day := "2023-03-03 12:12:25"
	timeU, err := GetDayFirstSecond(day, "")
	t.Log(timeU)
	t.Log(err)
}

func TestGetDayLastSecond(t *testing.T) {
	day := "2023-03-03 12:12:25"
	timeU, err := GetDayLastSecond(day, "")
	t.Log(timeU)
	t.Log(err)
}

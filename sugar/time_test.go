package sugar

import (
	"fmt"
	"testing"
)

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

func TestTimeFormat(t *testing.T) {
	timeUnix := 1677903566
	day, e := TimeFormat(int64(timeUnix), "2006-01-02 15:04:05", "")
	fmt.Println(e)
	fmt.Println(day)
}

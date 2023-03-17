package sugar

import (
	"strings"
	"time"
)

// GetDayFirstSecond 获取某天的 0时0分0秒
func GetDayFirstSecond(date, timeZone string) (int64, error) {
	t, err := GetTimeUnix(date, timeZone)
	if err != nil {
		return 0, err
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return 0, err
	}
	dayFirst := time.Unix(t, 0).In(loc).Format("2006-01-02 00:00:00")
	t, err = GetTimeUnix(dayFirst, timeZone)
	return t, nil
}

// GetDayLastSecond 获取某天的 23时59分59秒
func GetDayLastSecond(date, timeZone string) (int64, error) {
	t, err := GetTimeUnix(date, timeZone)
	if err != nil {
		return 0, err
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return 0, err
	}
	lastT := t + 86399
	dayFirst := time.Unix(lastT, 0).In(loc).Format("2006-01-02 00:00:00")
	t, err = GetTimeUnix(dayFirst, timeZone)
	return t, nil
}

// GetTimeUnix 获取时间戳， 如果没有传时区，默认是 "上海"
func GetTimeUnix(date string, timeZone string) (int64, error) {
	if timeZone == "" {
		timeZone = "Asia/Shanghai"
	}
	// 时间转成时间戳
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return 0, err
	}
	// 确定格式
	timeFormat := "2006-01-02 15:04:05"
	if strings.Index(date, " ") == -1 {
		timeFormat = "2006-01-02"
	}
	unixTime, err := time.ParseInLocation(timeFormat, date, loc)
	// unixTime, err := time.ParseInLocation("2006-01-02", date, loc)
	if err != nil {
		return 0, err
	}
	return unixTime.Unix(), nil
}

func TimeCurrUnix(timeZone string) (int64, error) {
	if timeZone == "" {
		timeZone = "Asia/Shanghai"
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return 0, err
	}
	currTime := time.Now().In(loc).Unix()
	return currTime, nil
}

// TimeFormat 时间戳转换成日期 23时59分59秒
func TimeFormat(t int64, format string, timeZone string) (string, error) {
	if timeZone == "" {
		timeZone = "Asia/Shanghai"
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return "", err
	}
	day := time.Unix(t, 0).In(loc).Format(format)
	return day, nil
}

// CurrTimeFormat 当前时间戳转换成日期 23时59分59秒
func CurrTimeFormat(format string, timeZone string) (string, error) {
	if timeZone == "" {
		timeZone = "Asia/Shanghai"
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return "", err
	}
	day := time.Now().In(loc).Format(format)
	return day, nil
}

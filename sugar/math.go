package sugar

import (
	"math"
	"time"
)

// MinInt 比较两个数，返回最小值
func MinInt(a, b int64) int64 {
	return int64(math.Min(float64(a), float64(b)))
}

// MaxInt 比较两个数，返回最大值
func MaxInt(a, b int64) int64 {
	return int64(math.Max(float64(a), float64(b)))
}

// MaxStringNum 比较两个数组字符串 ，
// 参数格式 :  YYYYmmdd
// MaxStringNum(20220112，20220212)
func MaxStringNum(a, b string, timeZone string) (string, error) {
	aT, err := time.Parse("20060102", a)
	if err != nil {
		return "", err
	}
	bT, err := time.Parse("20060102", b)
	if err != nil {
		return "", err
	}
	if aT.Before(bT) {
		// aT<bT
		return b, nil
	}
	return a, nil
}

package sugar

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf16"
)

// 表情解码
func EmojiDecode(s string) string {
	re := regexp.MustCompile("\\\\u[ed][0-9a-f]{3}\\\\u[ed][0-9a-f]{3}")
	result := re.FindAllString(s, -1)
	if len(result) > 0 {
		for k := range result {
			tmp := result[k]
			result[k] = strings.Replace(result[k], "\\u", "0x", -1)
			cov1, _ := strconv.ParseUint(result[k][:6], 0, 32)
			cov2, _ := strconv.ParseUint(result[k][6:], 0, 32)
			covResult := fmt.Sprintf("%c", utf16.DecodeRune(rune(cov1), rune(cov2)))
			s = strings.Replace(s, tmp, covResult, -1)
		}
	}
	return s
}

// EmojiEncode Emoji表情转码
func EmojiEncode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			r1, r2 := utf16.EncodeRune(rs[i])
			t1 := "\\u" + fmt.Sprintf("%x", r1)
			t2 := "\\u" + fmt.Sprintf("%x", r2)
			ret += t1 + t2
		} else {
			ret += string(rs[i])
		}
	}
	return ret
}

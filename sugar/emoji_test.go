package sugar

import "testing"

func TestEmojiDecode(t *testing.T) {
	str := "文案3，你好哈\\ud83d\\udc4b"
	lastStr := "文案3，你好哈👋"
	newStr := EmojiDecode(str)
	if newStr != lastStr {
		t.Fatalf("EmojiDecode fail")
	}
}

func TestEmojiEncode(t *testing.T) {
	str := "文案3，你好哈👋"
	lastStr := "文案3，你好哈\\ud83d\\udc4b"

	newStr := EmojiEncode(str)
	if newStr != lastStr {
		t.Fatalf("TestEmojiEncode fail")
	}
}

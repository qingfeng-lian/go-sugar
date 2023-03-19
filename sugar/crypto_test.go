package sugar

import "testing"

func TestMd5(t *testing.T) {
	qf := "qingfeng"
	needMd5 := "27e9244b456c18df737b57fdbc4558a8"
	realCreate := Md5(qf)
	if realCreate != needMd5 {
		t.Errorf("create md5 error: create:%s, need:%s", realCreate, needMd5)
		return
	}
	t.Log("create md5", realCreate)
}

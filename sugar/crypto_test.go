package sugar

import (
	"fmt"
	"testing"
)

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

func TestRC4Encrypt(t *testing.T) {
	qf := "qingfeng"
	key := "qf-test"
	needMd5 := []byte{164, 168, 13, 224, 204, 81, 187, 96}

	realCreate := RC4Encrypt(key, qf)
	fmt.Println(realCreate)
	if string(realCreate) != string(needMd5) {
		t.Errorf("create rc4 error: create:%s, need:%s", realCreate, needMd5)
		return
	}
	t.Log("create rc4", realCreate)
}

func TestRC4EncryptBaseEncode(t *testing.T) {
	qf := "qingfeng"
	key := "qf-test"
	needMd5 := "pKgN4MxRu2A="
	realCreate := RC4EncryptBaseEncode(key, qf)
	if realCreate != needMd5 {
		t.Errorf("create rc4 error: create:%s, need:%s", realCreate, needMd5)
		return
	}
	t.Log("create rc4", realCreate)
}

func TestRC4EncryptBaseDecode(t *testing.T) {
	qf := "pKgN4MxRu2A="
	key := "qf-test"
	needMd5 := "qingfeng"
	realCreate := RC4EncryptBaseDecode(key, qf)
	if realCreate != needMd5 {
		t.Errorf("deconde rc4 error: create:%s, need:%s", realCreate, needMd5)
		return
	}
	t.Log("deconde rc4", realCreate)
}

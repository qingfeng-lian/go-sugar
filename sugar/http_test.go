package sugar

import (
	"fmt"
	"testing"
)

func TestHttpPost(t *testing.T) {
	url := "http://op-pre.xiaoyezi.com/ref/student/identity"
	type resParam struct {
		StudentUuid string `json:"student_uuid"`
	}
	// res := resParam{
	// 	StudentUuid: "426477035032959249",
	// }
	res := map[string]string{
		"student_uuid": "426477035032959249",
	}
	req := map[string]interface{}{}
	// err := HttpPostForm(url, res, req)
	err := HttpPost(url, res, req, map[string]string{})
	t.Log("dddddddddd", err, req)
}

func TestHttpGet(t *testing.T) {
	url := "http://dss-pre.xiaoyezi.com/internal/student"
	type reqData struct {
		Code int `json:"code"`
		Data struct {
			Id   uint64 `json:"id"`
			Name string `json:"name"`
		} `json:"data"`
	}
	res := map[string]string{
		"id": "10211",
	}
	req := reqData{}
	err := HttpGet(url, res, &req, map[string]string{})
	t.Log(err)
	t.Log(fmt.Sprintf("req:%+v", req))
}

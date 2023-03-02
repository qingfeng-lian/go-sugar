package sugar

import "testing"

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
	err := HttpPost(url, res, req)
	t.Log("dddddddddd", err, req)
}

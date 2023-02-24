package sugar

import "testing"

func TestPathExists(t *testing.T) {
	path := "/ddddd"
	isExists := PathExists(path)
	t.Log(isExists)
}

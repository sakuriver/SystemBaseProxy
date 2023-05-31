package network

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQueryParameters(t *testing.T) {
	reqBody := bytes.NewBufferString("request body")
	req := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)

	// 生成後は*http.Requestオブジェクトと同じように扱える
	q := req.URL.Query()
	q.Add("check", "aaabbbccc")
	req.URL.RawQuery = q.Encode()
	resultQuery := GetQueryParameters(req)

	if resultQuery.Get("check") == "" {
		t.Errorf("require parameter error")
	}

}

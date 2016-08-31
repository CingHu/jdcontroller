package app

import (
	"testing"
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

func RequestHttp(method, urlStr string, bodyMap map[string]interface{}) (int, []byte) {
	client := &http.Client{}
	//	urlStr = "http://192.168.15.51:9090/v1" + urlStr
	urlStr = "http://127.0.0.1:9099/v1"+urlStr
	bodyEncode, _ := encodeJson(bodyMap)
	fmt.Printf(urlStr)
	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(bodyEncode))
	if err != nil {
		return 408, nil
	}

	req.Header.Set("User-Agent", "XENIUMD-AGENT")
	if method == "POST" {
		req.Header.Set("Content-Type", "plain/text")
	}

	resp, err := client.Do(req)
	if err != nil {
		status := -1
		if resp != nil {
			status = resp.StatusCode
		}
		return status, nil
	}

	dataBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return resp.StatusCode, nil
	}
	return resp.StatusCode, dataBody
}

func encodeJson(data interface {}) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Test_getTunnelFlow(t *testing.T) {
	url := "/ofctl/tunnelflow"

	jsonMap := make(map[string] interface{})
	jsonMap["dpid"] = "00006a27855efb4a"
	httpCode, body := RequestHttp("GET", url, jsonMap)

	if httpCode != 200 {
		t.Log("httpCode:", httpCode)
		t.Log("body:", string(body))
		t.Fatal("ERROR")
	}

	if body == nil {
		t.Fatal("body is nil")
	}

	t.Log(string(body), "\n")
}


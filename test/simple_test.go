package gotest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAnswer(t *testing.T) {
	msg := "深圳天气"
	//get请求
	//http.Get的参数必须是带http://协议头的完整url,不然请求结果为空
	resp, _ := http.Get("http://api.qingyunke.com/api.php?key=free&appid=0&msg=" + msg)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	response := string(body)
	var robotMsg robotMsg
	if err := json.Unmarshal([]byte(response), &robotMsg); err == nil {
		t.Log(robotMsg.Content)
	} else {
		t.Error("Test_Answer:" + err.Error())
	}

}

type robotMsg struct {
	Result  int    `json:"result"`
	Content string `json:"content"`
}

package gotest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	test := "搜题 琳"
	split := strings.Split(test, " ")

	resp, _ := http.Post("https://ninja.yua.im/ninja/qa",
		"application/x-www-form-urlencoded",
		strings.NewReader("search="+split[1]))

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	response := string(body)

	//msg := "未找到结果"

	var exam Exam
	if err := json.Unmarshal([]byte(response), &exam); err == nil {
		if exam.IsSuc && exam.Data.Total > 0 {
			all := make([]string, exam.Data.Total+1)
			all[0] = "小改改为你找到以下结果:"
			for index, v := range exam.Data.Rows {
				title := strconv.Itoa(index+1) + ". " + v.Title
				s1 := make([]string, len(v.Answers)+1)
				s1[0] = title
				for k, v := range v.Answers {
					if v.IsCorrect {
						//正确
						s1[k+1] = v.Content + " √"
					} else {
						s1[k+1] = v.Content
					}
				}
				//一条题目及回答
				all[index+1] = strings.Join(s1, "\n")
			}
			t.Log(strings.Join(all, "\n"))
		}
	} else {
		t.Error("TestSearch:" + err.Error())
	}
}

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

type Exam struct {
	Data struct {
		Rows []struct {
			_id     string `json:"_id"`
			Agree   int    `json:"agree"`
			Answers []struct {
				Content   string `json:"content"`
				IsCorrect bool   `json:"is_correct"`
			} `json:"answers"`
			CreateDt  string `json:"create_dt"`
			CreatedBy string `json:"created_by"`
			Disagree  int    `json:"disagree"`
			Tip       string `json:"tip"`
			Title     string `json:"title"`
			UpdateDt  string `json:"update_dt"`
			UpdatedBy string `json:"updated_by"`
		} `json:"rows"`
		Total int `json:"total"`
	} `json:"data"`
	IsSuc bool `json:"is_suc"`
}

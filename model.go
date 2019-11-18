package main

import "github.com/jinzhu/gorm"

//机器人智能回复消息
type RobotMsg struct {
	Result  int    `json:"result"`
	Content string `json:"content"`
}

//兑换码
type Key struct {
	gorm.Model
	Key string `gorm:"size:128"`
}

//搜题
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

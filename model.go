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

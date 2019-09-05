package main

import "github.com/Tnze/CoolQ-Golang-SDK/cqp"

//go:generate cqcfg -c .
// cqp: 名称: 忍3管家
// cqp: 版本: 1.0.1:1
// cqp: 作者: mtdhllf
// cqp: 简介: 巴拉巴拉~
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "me.cqp.mtdhllf.ninja.robot"
	cqp.PrivateMsg = onPrivateMsg
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	cqp.SendPrivateMsg(fromQQ, msg) //复读机
	return 0
}

package main

import (
	"fmt"
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/robfig/cron"
	"strings"
	"time"
)

//go:generate cqcfg -c .
// cqp: 名称: 忍3管家
// cqp: 版本: 1.0.1:1
// cqp: 作者: mtdhllf
// cqp: 简介: 巴拉巴拉~
func main() { /*此处应当留空*/ }

var db, err = gorm.Open("postgres", "postgres", "host=94.191.102.44 port=5432 user=postgres dbname=robot password=postgres sslmode=disable")

var c = cron.New()

func init() {
	cqp.AppID = "me.cqp.mtdhllf.ninja.robot"
	cqp.PrivateMsg = onPrivateMsg
	cqp.GroupMsg = onGroupMsg

	//数据库
	if err != nil {
		panic(err.Error())
	}
	// 自动迁移模式
	db.AutoMigrate(&Key{})

	//定时器1(7-23时,每小时执行一次)
	spec1 := "0 0 7-23/1 * * ?"
	err = c.AddFunc(spec1, testJob)

	if err != nil {
		panic(err.Error())
	}
	c.Start()
}

/**私聊入口*/
func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	code := onKeyPrivateMsg(subType, msgID, fromQQ, msg, font)
	//cqp.SendPrivateMsg(fromQQ, msg) //复读机
	return code
}

/**群聊入口*/
func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	code := onKeyGroupMsg(subType, msgID, fromGroup, fromQQ, fromAnonymous, msg, font)
	return code
}

//<editor-fold defaultstate="collapsed" desc="结构体">
//兑换码
type Key struct {
	gorm.Model
	Key string `gorm:"size:128"`
}

//</editor-fold>

//<editor-fold defaultstate="collapsed" desc="私聊">
func onKeyPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	code := int32(0)
	if strings.Contains(msg, "兑换码") {
		switch {
		//查询
		case strings.Contains(msg, "本周") || strings.Contains(msg, "这周") || strings.Contains(msg, "查询"):
			code = 1
			//时间
			now := time.Now()
			offset := int(time.Monday - now.Weekday())
			if offset > 0 {
				offset = -6
			}
			weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
			weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset+7)
			fmt.Println(weekStart, weekEnd)
			//结果
			var results []Key
			//查询
			db.Order("created_at").Find(&results, "created_at BETWEEN ? AND ?", weekStart, weekEnd)
			//生成结果
			if len(results) > 0 {
				var keys []string
				keys = append(keys, "本周兑换码:")
				for _, v := range results {
					keys = append(keys, v.Key)
				}
				cqp.SendPrivateMsg(fromQQ, strings.Join(keys, "\n"))
			} else {
				cqp.SendPrivateMsg(fromQQ, "本周兑换码:\n暂无兑换码,你可以私聊我<兑换码 帮助>来获取操作指令哦")
			}
		//新增
		case strings.HasPrefix(msg, "兑换码新增 "):
			code = 1
			var key string
			n, err := fmt.Sscanf(msg, "兑换码新增 %s", &key)
			if n != 1 || err != nil {
				cqp.AddLog(cqp.Debug, "兑换码", "兑换码指令不正确: "+msg)
				return 0
			}
			if len(key) > 0 {
				var old Key
				db.Find(&old, "key = ?", key)
				if old.Key != key {
					db.Create(&Key{Key: key})
					cqp.SendPrivateMsg(fromQQ, "新增兑换码: "+key)
				} else {
					cqp.SendPrivateMsg(fromQQ, "该兑换码已存在!")
				}
			}
		//删除
		case strings.HasPrefix(msg, "兑换码删除 "):
			code = 1
			var key string
			n, err := fmt.Sscanf(msg, "兑换码删除 %s", &key)
			if n != 1 || err != nil {
				cqp.AddLog(cqp.Debug, "兑换码", "兑换码指令不正确: "+msg)
				return 0
			}
			if len(key) > 0 {
				var old Key
				db.Find(&old, "key = ?", key)
				if old.Key == key {
					db.Delete(&Key{Key: key}, "key = ?", key)
					cqp.SendPrivateMsg(fromQQ, "删除兑换码: "+key)
				} else {
					cqp.SendPrivateMsg(fromQQ, "该兑换码不存在!")
				}
			}
		}

		//帮助
		if strings.Contains(msg, "帮助") {
			code = 1
			sendHelp(true, fromQQ)
		}
	}
	return code
}

//</editor-fold>

//<editor-fold defaultstate="collapsed" desc="群聊">
func onKeyGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	code := int32(0)
	if strings.Contains(msg, "兑换码") {
		switch {
		//查询
		case strings.Contains(msg, "本周") || strings.Contains(msg, "这周") || strings.Contains(msg, "查询"):
			code = 1
			//时间
			now := time.Now()
			offset := int(time.Monday - now.Weekday())
			if offset > 0 {
				offset = -6
			}
			weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
			weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset+7)
			fmt.Println(weekStart, weekEnd)
			//结果
			var results []Key
			//查询
			db.Order("created_at").Find(&results, "created_at BETWEEN ? AND ?", weekStart, weekEnd)
			//生成结果
			if len(results) > 0 {
				var keys []string
				keys = append(keys, "本周兑换码:")
				for _, v := range results {
					keys = append(keys, v.Key)
				}
				cqp.SendGroupMsg(fromGroup, strings.Join(keys, "\n"))
			} else {
				cqp.SendGroupMsg(fromGroup, "本周兑换码:\n暂无兑换码,你可以私聊我<兑换码 帮助>来获取操作指令哦")
			}
		//新增
		case strings.HasPrefix(msg, "兑换码新增 "):
			code = 1
			var key string
			n, err := fmt.Sscanf(msg, "兑换码新增 %s", &key)
			if n != 1 || err != nil {
				cqp.AddLog(cqp.Debug, "兑换码", "兑换码指令不正确: "+msg)
				return 0
			}
			if len(key) > 0 {
				var old Key
				db.Find(&old, "key = ?", key)
				if old.Key != key {
					db.Create(&Key{Key: key})
					cqp.SendGroupMsg(fromGroup, "新增兑换码: "+key)
				} else {
					cqp.SendGroupMsg(fromGroup, "该兑换码已存在!")
				}
			}
		//删除
		case strings.HasPrefix(msg, "兑换码删除 "):
			code = 1
			var key string
			n, err := fmt.Sscanf(msg, "兑换码删除 %s", &key)
			if n != 1 || err != nil {
				cqp.AddLog(cqp.Debug, "兑换码", "兑换码指令不正确: "+msg)
				return 0
			}
			if len(key) > 0 {
				var old Key
				db.Find(&old, "key = ?", key)
				if old.Key == key {
					db.Delete(&Key{Key: key}, "key = ?", key)
					cqp.SendGroupMsg(fromGroup, "删除兑换码: "+key)
				} else {
					cqp.SendGroupMsg(fromGroup, "该兑换码不存在!")
				}
			}
		}

		//帮助
		if strings.Contains(msg, "帮助") {
			code = 1
			sendHelp(false, fromGroup)
		}
	}
	return code
}

//</editor-fold>

/**帮助菜单*/
func sendHelp(single bool, from int64) {
	help := "兑换码功能:\n" +
		"1.查询指令☞<兑换码本周/本周兑换码/兑换码查询>\n" +
		"2.新增指令☞<兑换码新增 [key]>\n" +
		"3.删除指令☞<兑换码删除 [key]>\n"
	if single {
		cqp.SendPrivateMsg(from, help)
	} else {
		cqp.SendGroupMsg(from, help)
	}
}

/**定时任务*/
func testJob() {
	msg := "【滚来滚去】"
	cqp.SendGroupMsg(816440954, msg)
}

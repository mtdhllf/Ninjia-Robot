package key

import (
	"strings"
)

//兑换码功能-群聊处理
func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	if strings.HasPrefix(msg, "兑换码") {

	}
	return 0
}

//兑换码功能-私聊处理
func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	if strings.HasPrefix(msg, "兑换码") {

	}
	return 0
}

package redis_user

import (
	"fmt"

	"github.com/AnnonaOrg/annona_client/internal/constvar"
	"github.com/AnnonaOrg/annona_client/internal/dbredis"
	"github.com/AnnonaOrg/annona_client/model/user_info"
)

// 获取所有关键词列表
func GetAllKeyword() ([]string, error) {
	return GetAllByKey(constvar.REDIS_SET_KEY_ALL_KEYWORD)
}

// 获取所有用户列表
func GetAllUserInfoHashList() ([]string, error) {
	return GetAllByKey(constvar.REDIS_SET_KEY_ALL_USER)
}

// 获取所有屏蔽词列表
func GetAllBlockword() ([]string, error) {
	return GetAllByKey(constvar.REDIS_SET_KEY_ALL_BLOCKKEYWORD)
}

// 获取所有列表
func GetAllByKey(keyStr string) ([]string, error) {
	var err error
	list := make([]string, 0)

	list, err = dbredis.GetSetMembers(
		keyStr,
	)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list, nil
	}
	return list, fmt.Errorf("GET(%s): Null", keyStr)
}

// 检查关键词是否在用户关键词列表 true 在，不在 false
func IsKeywordOfUserCheck(userInfoHash string, keywordStr string) bool {
	key := constvar.REDIS_SET_KEY_USER_PREFIX_KEYWORD + userInfoHash
	if isOk, err := dbredis.IsMemberOfSet(key, keywordStr); err != nil {
		return false
	} else {
		return isOk
	}
}

// 根据用户infohash 获取用户信息
func GetUserInfoByUserInfoHash(userInfoHash string) (user_info.UserInfo, error) {
	var userInfo user_info.UserInfo
	key := constvar.REDIS_KEY_PREFIX_USERINFO + userInfoHash
	err := dbredis.GetKeyValue(key, &userInfo)
	return userInfo, err
}

// 根据机器人ID获取机器人token
func GetBotTokenByBotId(botID string) (string, error) {
	var retText string
	key := constvar.REDIS_KEY_PREFIX_BOT_TOKEN + botID
	err := dbredis.GetKeyValue(key, &retText)
	return retText, err
}

// 检查屏蔽来源会话id是否在所有列表中 true 在，不在 false
func IsBlockformchatidOfAllCheck(checkStr string) bool {
	key := constvar.REDIS_SET_KEY_ALL_BLOCKFORMCHATID
	if isOk, err := dbredis.IsMemberOfSet(key, checkStr); err != nil {
		return false
	} else {
		return isOk
	}
}

// 检查屏蔽来源Senderid是否在所有列表中 true 在，不在 false
func IsBlockformsenderidOfAllCheck(checkStr string) bool {
	key := constvar.REDIS_SET_KEY_ALL_BLOCKFORMSENDERID
	if isOk, err := dbredis.IsMemberOfSet(key, checkStr); err != nil {
		return false
	} else {
		return isOk
	}
}

// 检查屏蔽来源会话id是否在用户列表中 true 在，不在 false
func IsBlockformchatidOfUserCheck(userInfoHash string, checkStr string) bool {
	key := constvar.REDIS_SET_KEY_USER_PREFIX_BLOCKFORMCHATID + userInfoHash
	if isOk, err := dbredis.IsMemberOfSet(key, checkStr); err != nil {
		return false
	} else {
		return isOk
	}
}

// 检查屏蔽关键词是否在用户列表中 true 在，不在 false
func IsBlockwordOfUserCheck(userInfoHash string, checkStr string) bool {
	key := constvar.REDIS_SET_KEY_USER_PREFIX_BLOCKKEYWORD + userInfoHash
	if isOk, err := dbredis.IsMemberOfSet(key, checkStr); err != nil {
		return false
	} else {
		return isOk
	}
}

// 检查屏蔽来源Senderid是否在用户列表中 true 在，不在 false
func IsBlockformsenderidOfUserCheck(userInfoHash string, checkStr string) bool {
	key := constvar.REDIS_SET_KEY_USER_PREFIX_BLOCKFORMSENDERID + userInfoHash
	if isOk, err := dbredis.IsMemberOfSet(key, checkStr); err != nil {
		return false
	} else {
		return isOk
	}
}

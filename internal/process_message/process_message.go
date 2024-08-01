package process_message

import (
	"fmt"
	"strings"
	"time"

	"github.com/AnnonaOrg/annona_client/internal/service"

	"github.com/AnnonaOrg/annona_client/internal/redis_user"
	log "github.com/sirupsen/logrus"
)

// processMessageKeywords 处理消息关键词
// chatID 回话ID
// senderID 发送者ID
// messageID 消息ID
// messageDate 消息日期
// messageContentText 消息内容(处理过的)
// originalText 原始消息内容
// messageLink 消息链接
// messageLinkIsPublic 消息链接是否为公开链接
func ProcessMessageKeywords(chatID, senderID int64, senderUsername string, messageID int64, messageDate, messageContentText, originalText string, messageLink string, messageLinkIsPublic bool) {
	chatIDStr := fmt.Sprintf("%d", chatID)
	senderIDStr := fmt.Sprintf("%d", senderID)
	messageIDStr := fmt.Sprintf("%d_%d", chatID, messageID)
	allUserMap := make(map[string]string, 0)
	allBlockUserMap := make(map[string]string, 0)
	isInBlockFromChatID := redis_user.IsBlockformchatidOfAllCheck(chatIDStr)
	isInBlockFromSenderID := redis_user.IsBlockformsenderidOfAllCheck(senderIDStr)

	// 将原始消息文本转换为小写，方便进行关键词匹配（忽略大小写）
	normalizedText := strings.ToLower(originalText)

	allUserList, err := redis_user.GetAllUserInfoHashList()
	if err != nil {
		log.Errorf("GetAllUserInfoHashList()Fail: %v", err)
		return
	} else {
		// log.Debugf("GetAllUserInfoHashList()Success: %+v", allUserList)
	}
	// 检测屏蔽群组信息
	if isInBlockFromChatID {
		for _, userHash := range allUserList {
			// 跳过已经检出的userHash
			if _, isOk := allBlockUserMap[userHash]; isOk {
				continue
			}
			//  记录检出的userHash
			if isOk := redis_user.IsBlockformchatidOfUserCheck(userHash, chatIDStr); isOk {
				allBlockUserMap[userHash] = chatIDStr
				// log.Infof("屏蔽来源会话ID '%s' 在用户(%s)中被检测到！\n", chatID, userHash)
				log.Debugf("IsBlockformchatidOfUserCheck(%s,%s): true", userHash, chatIDStr)
			}
		}
	}
	if len(allBlockUserMap) == len(allUserList) {
		log.Debugf("len(allBlockUserMap) == len(allUserList)")
		return
	}
	// 检测屏蔽用户信息
	if isInBlockFromSenderID {
		for _, userHash := range allUserList {
			// 跳过已经检出的userHash
			if _, isOk := allBlockUserMap[userHash]; isOk {
				continue
			}
			//  记录检出的userHash
			if isOk := redis_user.IsBlockformsenderidOfUserCheck(userHash, senderIDStr); isOk {
				allBlockUserMap[userHash] = senderIDStr
				// log.Infof("屏蔽来源会话ID '%s' 在用户(%s)中被检测到！\n", senderID, userHash)
				log.Debugf("IsBlockformsenderidOfUserCheck(%s,%s): true", userHash, senderIDStr)
			}
		}
	}
	if len(allBlockUserMap) == len(allUserList) {
		log.Debugf("len(allBlockUserMap) == len(allUserList)")
		return
	}

	allBlockwordList, err := redis_user.GetAllBlockword()
	if err != nil {
		// log.Errorf("GetAllBlockword()Fail: %v", err)
	} else {
		log.Debugf("GetAllBlockword()Success: %+v", allBlockwordList)
	}
	// 检测屏蔽关键词
	for _, keyword := range allBlockwordList {
		if strings.Contains(normalizedText, keyword) {
			// log.Infof("屏蔽关键词 '%s' 在消息中被检测到！\n", keyword)
			log.Debugf("Blockword(%s): %s", keyword, originalText)
			// 检出符合条件的用户 userHash
			for _, userHash := range allUserList {
				// 跳过已经检出的userHash
				if _, isOk := allBlockUserMap[userHash]; isOk {
					continue
				}
				// 记录检出的userHash
				if isOk := redis_user.IsBlockwordOfUserCheck(userHash, keyword); isOk {
					allBlockUserMap[userHash] = keyword
					// log.Infof("屏蔽关键词 '%s' 在用户(%s)中被检测到！\n", keyword, userHash)
					log.Debugf("IsBlockwordOfUserCheck(%s,%s): true", userHash, keyword)
				}
			}
		}
		// 当检出的用户数和总用户相等的时候 就跳过后面的关键词检测
		if len(allBlockUserMap) == len(allUserList) {
			break
		}
	}

	if len(allBlockUserMap) == len(allUserList) {
		log.Debugf("len(allBlockUserMap) == len(allUserList)")
		return
	}

	allKeywordList, err := redis_user.GetAllKeyword()
	if err != nil {
		log.Errorf("GetAllKeyword()Fail: %v", err)
		return
	} else {
		// log.Debugf("GetAllKeyword()Success: %+v", allKeywordList)
	}
	// 检测关键词
	for _, keyword := range allKeywordList {
		if strings.Contains(normalizedText, keyword) {
			log.Debugf("Keyword(%s): %s", keyword, originalText)
			// 检出符合条件的用户 userHash
			for _, userHash := range allUserList {
				// 跳过已经Block 的 userHash
				if _, isOk := allBlockUserMap[userHash]; isOk {
					continue
				}
				// 跳过已经检出的userHash
				if _, isOk := allUserMap[userHash]; isOk {
					continue
				}
				// 记录检出的userHash
				if isOk := redis_user.IsKeywordOfUserCheck(userHash, keyword); isOk {
					allUserMap[userHash] = keyword
					log.Debugf("IsKeywordOfUserCheck(%s,%s): true", userHash, keyword)
				}
			}

		}
		// 当检出的用户数和总用户相等的时候 就跳过后面的关键词检测
		if len(allUserMap) == len(allUserList) {
			break
		}
	}

	if len(allUserMap) == 0 {
		log.Debugf("allUserMap: %+v", allUserMap)
		return
	}

	// 根据检出的用户信息map 推送信息
	var keyworldList []string
	for k, v := range allUserMap {
		vc := v
		keyworldList = append(keyworldList, vc)

		if user, err := redis_user.GetUserInfoByUserInfoHash(k); err != nil {
			log.Errorf("GetUserInfoByUserInfoHash(%s)Fail: %v", k, err)
			continue
		} else {
			log.Debugf("GetUserInfoByUserInfoHash(%s)Success: %+v", k, user)
			timeDuration := time.Unix(user.Exp, 0).Sub(time.Now())
			if timeDuration < time.Second {
				log.Debugf("Ignore the user %s (%d) Exp: %d timeDuration: %d", user.TelegramUsername, user.TelegramChatId, user.Exp, timeDuration)
				continue
			}
			if timeDuration < time.Hour {
				messageContentText = "服务即将到期(可购充值卡续期)\n" + messageContentText
			}

			toChatID := user.TelegramChatId
			if user.TelegramNoticeChatId != 0 {
				toChatID = user.TelegramNoticeChatId
			}
			startBotId := fmt.Sprintf("%d", user.TelegramStartBotId)
			botToken := ""
			if botTokenTmp, err := redis_user.GetBotTokenByBotId(startBotId); err != nil {
				log.Errorf("GetBotTokenByBotId(%s)Fail: %v", startBotId, err)
			} else {
				botToken = botTokenTmp
				log.Debugf("GetBotTokenByBotId(%s)Success: %v", startBotId, botToken)
			}
			messageContentText = "关键词: #" + vc + "\n" + messageContentText
			log.Debugf("will send messageContentText: %s To userInfo: %+v", messageContentText, user)
			if retText, err := sendMessage(
				chatIDStr, senderIDStr, toChatID, botToken,
				messageIDStr, messageDate, messageContentText, messageLink, messageLinkIsPublic,
				vc,
			); err != nil {
				log.Errorf("msg Send( %s )Fail: %v", messageContentText, err)
			} else {
				log.Debugf("msg Send( %s )Success: %s", messageContentText, retText)
			}
		}
	}

	go service.CreateKeyworldHistoryEx(
		chatID, senderID, senderUsername,
		messageID, originalText, messageLink,
		strings.Join(keyworldList, ","),
	)
}

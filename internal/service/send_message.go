package service

import (
	"fmt"
	"strconv"

	"github.com/AnnonaOrg/annona_client/internal/request"
	"github.com/AnnonaOrg/osenv"
	"github.com/AnnonaOrg/pkg/send2server"
	log "github.com/sirupsen/logrus"
)

// processMessageKeywords 处理消息关键词
// botToken 机器人Bot
// toChatID 发送到ID
// messageID 消息ID
// messageContentText 消息内容(处理过的)
func SendMessage(
	formMessageID int64,
	chatIDStr, senderIDStr string, toChatID int64, botToken,
	messageID, messageDate, messageContentText string, messageLink string, messageLinkIsPublic bool,
	keyworld string,
) (string, error) {
	senderID, _ := strconv.ParseInt(senderIDStr, 10, 64)
	chatID, _ := strconv.ParseInt(chatIDStr, 10, 64)

	chatUsername := GetUsername(chatID)
	chatTitle := GetChatTitle(chatID)
	senderUsername := GetUsername(senderID)
	senderTitle := GetUserFirstLastName(senderID)

	var richMsg request.FeedRichMsgRequest

	richMsg.BotInfo.BotToken = botToken
	richMsg.ChatInfo.ToChatID = toChatID
	richMsg.MsgID = fmt.Sprintf("%d", toChatID) + messageID
	richMsg.MsgTime = messageDate
	richMsg.Msgtype = "rich"
	richMsg.Text.Content = messageContentText

	richMsg.FormInfo.FormMessageID = formMessageID
	richMsg.FormInfo.FormChatID = chatIDStr
	richMsg.FormInfo.FormChatUsername = chatUsername
	richMsg.FormInfo.FormChatTitle = chatTitle

	richMsg.FormInfo.FormSenderID = senderIDStr
	richMsg.FormInfo.FormSenderUsername = senderUsername
	richMsg.FormInfo.FormSenderTitle = senderTitle

	richMsg.FormInfo.FormKeyworld = keyworld

	richMsg.NoButton = false
	richMsg.Link = messageLink
	richMsg.LinkIsPublic = messageLinkIsPublic

	msgContentSuffix := ""
	if len(richMsg.FormInfo.FormChatTitle) > 0 {
		msgContentSuffix = "来源:" + richMsg.FormInfo.FormChatTitle + "\n"
		if len(richMsg.FormInfo.FormSenderTitle) > 0 {
			msgContentSuffix = "发送人:" + richMsg.FormInfo.FormSenderTitle + "\n" + msgContentSuffix
		}
	}
	// retText = fmt.Sprintf("#ID%d", senderID)
	msgContentSuffix = msgContentSuffix + "#ID" + senderIDStr
	if len(senderUsername) > 0 {
		msgContentSuffix = msgContentSuffix + " @" + senderUsername
	}
	richMsg.Text.ContentEx = messageContentText + "\n" + msgContentSuffix

	serverRouter := osenv.GetNoticeOfFeedRichMsgPushUrl()
	serverChannel := fmt.Sprintf("%d", toChatID)
	serverToken := fmt.Sprintf("%d", toChatID)
	serverPath := osenv.GetNoticeOfFeedRichMsgPushUrlPath()

	retText, err := send2server.SendMsgToServer(
		richMsg,
		serverRouter,
		serverChannel,
		serverToken,
		serverPath,
	)
	log.Debugf("sendMessage(%+v),serverRouter(%s),serverChannel(%s),serverToken(%s),serverPath(%s): %s , %v",
		richMsg,
		serverRouter, serverChannel, serverToken, serverPath,
		retText, err,
	)
	return retText, err

}

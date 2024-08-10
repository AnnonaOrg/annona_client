package service

import (
	"fmt"
	// "strconv"

	"github.com/AnnonaOrg/annona_client/internal/request"
	"github.com/AnnonaOrg/osenv"
	"github.com/AnnonaOrg/pkg/send2server"
	log "github.com/sirupsen/logrus"
)

// processMessageKeywords 处理消息关键词
// formMessageID 消息ID
// botToken 机器人Bot
// toChatID 发送到ID
// messageIDStr fmt.Sprintf("%d_%d", chatID, messageID)
// messageContentText 消息内容(处理过的)
func SendMessage(
	formMessageID int64,
	chatID, senderID int64, toChatID int64, botToken,
	messageIDStr, messageDate, messageContentText string, messageLink string, messageLinkIsPublic bool,
	keyworld string,
) (string, error) {
	// senderID, _ := strconv.ParseInt(senderIDStr, 10, 64)
	// chatID, _ := strconv.ParseInt(chatIDStr, 10, 64)
	chatIDStr := fmt.Sprintf("%d", chatID)
	senderIDStr := fmt.Sprintf("%d", senderID)

	chatUsername := GetUsername(chatID)
	chatTitle := GetChatTitle(chatID)
	senderUsername := GetUsername(senderID)
	senderTitle := GetUserFirstLastName(senderID)

	var richMsg request.FeedRichMsgRequest

	richMsg.BotInfo.BotToken = botToken
	richMsg.ChatInfo.ToChatID = toChatID
	richMsg.MsgID = fmt.Sprintf("%d_", toChatID) + messageIDStr
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
	msgContentSuffixHtml := ""

	if len(richMsg.FormInfo.FormSenderTitle) > 0 {
		textTmp := ""
		if len(senderUsername) > 0 {
			textTmp = " @" + senderUsername
		}
		textTmp = "发送人:" + richMsg.FormInfo.FormSenderTitle + textTmp
		msgContentSuffix = textTmp
	}

	if len(richMsg.FormInfo.FormChatTitle) > 0 {

		if len(chatUsername) > 0 {
			msgContentSuffixHtml = msgContentSuffix + "\n" +
				"来源:" + "<a href=\"http://t.me/" + chatUsername + "\">" + richMsg.FormInfo.FormChatTitle + "</a>"
		} else if len(messageLink) > 0 {
			msgContentSuffixHtml = msgContentSuffix + "\n" +
				"来源:" + "<a href=\"" + messageLink + "\">" + richMsg.FormInfo.FormChatTitle + "</a>"
		} else {
			msgContentSuffixHtml = msgContentSuffix + "\n" +
				"来源:" + richMsg.FormInfo.FormChatTitle
		}
		msgContentSuffix = msgContentSuffix + "\n" +
			"来源:" + richMsg.FormInfo.FormChatTitle
	}
	msgContentSuffix = msgContentSuffix + "\n" + "#ID" + senderIDStr
	msgContentSuffixHtml = msgContentSuffixHtml + "\n" + "#ID" + senderIDStr
	richMsg.Text.ContentEx = messageContentText + "\n" + msgContentSuffix
	richMsg.Text.ContentHtml = messageContentText + "\n" + msgContentSuffixHtml

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

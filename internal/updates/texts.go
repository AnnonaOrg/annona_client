package updates

import (
	"fmt"

	"github.com/AnnonaOrg/annona_client/internal/process_message"

	"github.com/AnnonaOrg/annona_client/internal/api"
	"github.com/AnnonaOrg/annona_client/utils"
	log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
)

// handleText handles incoming text messages.
func handleText(message *client.Message, senderID int64) {
	messageContent := api.GetMessageFormattedText(message.Content) // message.Content.(*client.MessageText)
	messageContentText := messageContent.Text
	// log.Info("messageContent", messageContent, "messageContent.Text", messageContentText.Text)
	var (
		messageLink         string
		messageLinkIsPublic bool
	)
	if message.CanGetMediaTimestampLinks {
		if messageLinkTmp, err := api.GetMessageLink(message.ChatId, message.Id, 0, false, message.IsTopicMessage); err != nil {
			log.Errorf("api.GetMessageLink(%d,%d)%w", message.ChatId, message.Id, err)
		} else {
			messageLink = messageLinkTmp.Link
			messageLinkIsPublic = messageLinkTmp.IsPublic
		}
	}

	newMsg := func(chatID, messageID, senderID int64,
		messageContentText string,
		messageLink string, messageLinkIsPublic bool,
		messageData int32,
	) string {
		// retText :=
		// 	"chatID: " + fmt.Sprintf("%d", chatID) + "\n" +
		// 		"messageID: " + fmt.Sprintf("%d", messageID) + "\n" +
		// 		"senderID: " + fmt.Sprintf("tg://user?id=%d", senderID) + "\n" +
		// 		"messageLink: " + messageLink + " " + fmt.Sprintf("%t", messageLinkIsPublic) + "\n" +
		// 		"messageData: " + utils.FormatTimestamp2String(int64(messageData)) + "\n" +
		// 		"messageContentText: " + messageContentText
		retText :=
			// "群组ID: " + fmt.Sprintf("%d", chatID) + "\n" +
			// "messageID: " + fmt.Sprintf("%d", messageID) + "\n" +
			fmt.Sprintf("#ID%d", senderID) + "\n" +
				fmt.Sprintf("用户ID: tg://user?id=%d", senderID) + "\n" +
				// "messageLink: " + messageLink + " " + fmt.Sprintf("%t", messageLinkIsPublic) + "\n" +
				"消息日期: " + utils.FormatTimestamp2String(int64(messageData)) + "\n" +
				"消息内容: \n" + messageContentText

		return retText
	}
	log.Debugf("newMsg: %s", newMsg(message.ChatId, message.Id, senderID,
		messageContentText,
		messageLink, messageLinkIsPublic,
		message.Date,
	))

	go process_message.ProcessMessageKeywords(
		fmt.Sprintf("%d", message.ChatId),
		fmt.Sprintf("%d", senderID),
		fmt.Sprintf("%d_%d", message.ChatId, message.Id),
		utils.FormatTimestamp2String(int64(message.Date)),
		newMsg(message.ChatId, message.Id, senderID,
			messageContentText,
			messageLink, messageLinkIsPublic,
			message.Date,
		),
		messageContentText,
		messageLink,
		messageLinkIsPublic,
	)
}

// message_str = f'**关键词: {keywords}** __群组信息__ \n\n用户ID: [tg://user?id={event.from_id.user_id}](tg://user?id={event.from_id.user_id}) \n群组ID: {event.chat_id} \n群名称: {event.chat.title}\n消息位置: [点击查看]({channel_msg_url})\n消息时间: {china_time_str}'

// func ProcessMessageKeywords(chatID, senderID, messageID, messageDate, messageContentText, originalText string, messageLink string, messageLinkIsPublic bool)

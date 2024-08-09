package updates

import (
	"fmt"

	"github.com/AnnonaOrg/annona_client/internal/api"
	"github.com/AnnonaOrg/annona_client/internal/process_message"
	"github.com/AnnonaOrg/annona_client/internal/service"
	"github.com/AnnonaOrg/annona_client/utils"
	"github.com/AnnonaOrg/osenv"
	log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
)

// handleText handles incoming text messages.
func handleText(message *client.Message, senderID int64, senderUsername string) {
	messageContent := api.GetMessageFormattedText(message.Content) // message.Content.(*client.MessageText)
	messageContentText := messageContent.Text
	// log.Info("messageContent", messageContent, "messageContent.Text", messageContentText.Text)
	var (
		messageLink         string
		messageLinkIsPublic bool
	)
	// if message.CanGetMediaTimestampLinks {
	// 	if !service.IsPublicChat(message.ChatId) {
	// 		if messageLinkTmp, err := api.GetMessageLink(message.ChatId, message.Id, 0, false, false); err != nil {
	// 			log.Errorf("handleText.(api.GetMessageLink(%d,%d,inMessageThread:false)): %v", message.ChatId, message.Id, err)
	// 			if messageLinkTmp, err := api.GetMessageLink(message.ChatId, message.Id, 0, false, true); err != nil {
	// 				log.Errorf("handleText.(api.GetMessageLink(%d,%d,inMessageThread:false)): %v", message.ChatId, message.Id, err)
	// 				service.SetNotPublicChat(message.ChatId, err.Error())
	// 			} else {
	// 				messageLink = messageLinkTmp.Link
	// 				messageLinkIsPublic = messageLinkTmp.IsPublic
	// 			}
	// 		} else {
	// 			messageLink = messageLinkTmp.Link
	// 			messageLinkIsPublic = messageLinkTmp.IsPublic
	// 		}

	// 	}
	// }
	if !service.IsPublicChat(message.ChatId) {
		if messageLinkTmp, err := api.GetMessageLink(message.ChatId, message.Id, 0, false, false); err != nil {
			log.Errorf("handleText.(api.GetMessageLink(%d,%d,inMessageThread:false)): %v", message.ChatId, message.Id, err)
			if messageLinkTmp, err := api.GetMessageLink(message.ChatId, message.Id, 0, false, true); err != nil {
				log.Errorf("handleText.(api.GetMessageLink(%d,%d,inMessageThread:false)): %v", message.ChatId, message.Id, err)
				service.SetNotPublicChat(message.ChatId, err.Error())
			} else {
				messageLink = messageLinkTmp.Link
				messageLinkIsPublic = messageLinkTmp.IsPublic
			}
		} else {
			messageLink = messageLinkTmp.Link
			messageLinkIsPublic = messageLinkTmp.IsPublic
		}
	}

	newMsg := func(chatID, messageID,
		senderID int64, senderUsername string,
		messageContentText string,
		messageLink string, messageLinkIsPublic bool,
		messageData int32,
	) string {
		retText := ""
		// retText = fmt.Sprintf("#ID%d", senderID)
		// if len(senderUsername) > 0 {
		// 	retText = retText + " @" + senderUsername
		// }
		if osenv.IsTDlibSimpleMessage() {
			messageContentText = utils.GetStringRuneN(messageContentText, 20)
		}
		// retText = //retText + //"\n" +
		// fmt.Sprintf("用户ID: tg://user?id=%d", senderID) + "\n" +
		// "messageLink: " + messageLink + " " + fmt.Sprintf("%t", messageLinkIsPublic) + "\n" +
		retText = "消息日期: " + utils.FormatTimestamp2String(int64(messageData)) + "\n" +
			"消息内容: \n" + messageContentText

		return retText
	}

	go process_message.ProcessMessageKeywords(
		message.ChatId,
		senderID, senderUsername,
		message.Id, //fmt.Sprintf("%d_%d", message.ChatId, message.Id),
		utils.FormatTimestamp2String(int64(message.Date)),
		newMsg(message.ChatId, message.Id,
			senderID, senderUsername,
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

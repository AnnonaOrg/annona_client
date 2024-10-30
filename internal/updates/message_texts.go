package updates

import (
	"strings"

	"github.com/AnnonaOrg/annona_client/internal/api"
	"github.com/AnnonaOrg/annona_client/internal/service"
	"github.com/AnnonaOrg/annona_client/utils"
	"github.com/AnnonaOrg/osenv"
	"github.com/zelenin/go-tdlib/client"
)

// handleText handles incoming text messages.
func handleText(message *client.Message, senderID int64, senderUsername string) {
	messageContent := api.GetMessageFormattedText(message.Content) // message.Content.(*client.MessageText)
	messageContentText := messageContent.Text

	if strings.EqualFold(messageContentText, "/ping") {
		if message.ChatId < 0 {
			log.Debugf("message: %+v", message)
			return
		}
		// if message.SenderId.MessageSenderType() == client.TypeMessageSenderUser {
		// 	sender := message.SenderId.(*client.MessageSenderUser)
		// 	senderID := sender.UserId
		// 	if _, err := api.SendMessageText("pong", senderID); err != nil {
		// 		log.Errorf("SendMessageText(pong,%d): %v", senderID, err)
		// 	}
		// }
		if _, err := api.SendMessageText("pong", senderID); err != nil {
			log.Errorf("SendMessageText(pong,%d): %v", senderID, err)
		}
		return
	}

	messageDateStr := utils.FormatTimestamp2String(int64(message.Date))
	messageContentTextEx := messageContentText
	if osenv.IsTDlibSimpleMessage() {
		messageContentTextEx = utils.GetStringRuneN(messageContentText, 20)
	}
	messageContentTextEx = "日期: " + messageDateStr + "\n" +
		"内容: " + messageContentTextEx

	go service.ProcessMessageKeywords(
		message.ChatId,
		senderID, senderUsername,
		message.Id,
		messageDateStr,
		messageContentTextEx,
		messageContentText,
		// messageLink,
		// messageLinkIsPublic,
		message.IsTopicMessage, //messageIsTopicMessage,
		int64(message.Date),
	)
}

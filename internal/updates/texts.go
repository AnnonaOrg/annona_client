package updates

import (
	"github.com/AnnonaOrg/annona_client/internal/api"
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
	var (
		messageLink         string
		messageLinkIsPublic bool
	)

	if messageLinkTmp, err := api.GetMessageLink(message.ChatId, message.Id, 0, false, message.IsTopicMessage); err != nil {
		log.Errorf("handleText.(api.GetMessageLink(%d,%d,inMessageThread:%t),MessageThreadId:%d): %v",
			message.ChatId, message.Id, message.IsTopicMessage, message.MessageThreadId,
			err)
	} else {
		messageLink = messageLinkTmp.Link
		messageLinkIsPublic = messageLinkTmp.IsPublic
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
		messageLink,
		messageLinkIsPublic,
		int64(message.Date),
	)
}

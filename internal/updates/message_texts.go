package updates

import (
	"strings"

	"github.com/AnnonaOrg/annona_client/internal/api"
	"github.com/AnnonaOrg/annona_client/internal/log"
	"github.com/AnnonaOrg/annona_client/internal/repository"
	"github.com/AnnonaOrg/annona_client/internal/service"
	"github.com/AnnonaOrg/annona_client/utils"
	"github.com/AnnonaOrg/osenv"
	"github.com/zelenin/go-tdlib/client"
)

// handleText handles incoming text messages.
func handleText(message *client.Message) {
	//跳过频道消息
	if message.IsChannelPost {
		return
	}
	chatID := message.ChatId
	if isTrue, err := api.IsCanGetMessageLink(chatID); !isTrue || err != nil {
		log.Errorf("IsCanGetMessageLink err: %v", err)
		return
	}

	messageContent := api.GetMessageFormattedText(message.Content) // message.Content.(*client.MessageText)
	messageContentText := messageContent.Text
	if strings.EqualFold(messageContentText, "/ping") {
		if message.ChatId < 0 {
			log.Debugf("message: %+v", message)
			return
		}
		if message.SenderId.MessageSenderType() == client.TypeMessageSenderUser {
			sender := message.SenderId.(*client.MessageSenderUser)
			senderID := sender.UserId
			if _, err := api.SendMessageText("pong", senderID); err != nil {
				log.Errorf("SendMessageText(pong,%d): %v", senderID, err)
			}
		}
		// if _, err := api.SendMessageText("pong", senderID); err != nil {
		// 	log.Errorf("SendMessageText(pong,%d): %v", senderID, err)
		// }
		return
	}

	if service.IsEnableBlockLongText() {
		if count := service.GetBlockLongTextMaxCount(); len([]rune(messageContentText)) > count {
			log.Debugf("忽略长文本: %s", messageContentText)
			return
		}
	}

	senderID, err := api.GetSenderID(message) //api.GetSenderUserID(message)
	if err != nil {
		log.Errorf("GetSenderID: %v", err)
		return
	}
	// 忽略群消息
	if senderID <= 0 {
		return
	}
	//跳过机器人自己发出的消息
	if senderID == repository.Me.Id {
		return
	}
	//跳过机器人消息
	if isBot := service.IsBotID(senderID); isBot {
		return
	}

	//跳过匿名消息
	if message.ChatId == senderID && message.ChatId > 0 {
		return
	}
	senderUsername := service.GetUsername(senderID)

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

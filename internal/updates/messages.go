package updates

import (
	"github.com/AnnonaOrg/annona_client/internal/api"
	"github.com/AnnonaOrg/annona_client/internal/repository"

	log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
)

// handleNewMessage handles incoming messages.
func handleNewMessage(message *client.Message) {
	// log.Debugf("Message: %#v", message)
	//跳过频道消息
	if message.IsChannelPost {
		log.Debugf("IsChannelPost: %#v", message)
		return
	}

	senderID, err := api.GetSenderID(message) //api.GetSenderUserID(message)
	if err != nil {
		log.Errorf("GetSenderID: %+v", err)
		return
	}
	if senderID <= 0 {
		log.Debugf("Ignore Message senderID(a channel?): %d", senderID)
		return
	}

	//跳过机器人自己发出的消息
	// Tdlib delivers updates from self
	if senderID == repository.Me.Id {
		log.Debug("Message from self")
		return
	}
	//跳过机器人消息
	if isBot, err := api.IsViaBotByUserID(senderID, message.ViaBotUserId); err != nil {
		log.Debugf("IsViaBot(err): %v", err)
	} else if isBot {
		log.Debugf("Bot Message: %#v", message)
		// log.Infof("IsViaBot: %d", message.ViaBotUserId)
		return
	}
	//跳过匿名消息
	if message.ChatId == senderID && message.ChatId > 0 {
		log.Debugf("Message is anonymous")
		return
	}
	senderUsername, _ := api.GetUsernameByID(senderID)
	// if err != nil {
	// 	log.Errorf("GetUsernameByID(%d): %+v", senderID, err)
	// }
	// if !message.IsChannelPost && !message.IsTopicMessage && message.ChatId == senderID && message.ChatId > 0 {
	// 	log.Debugf("Message is private")
	// }

	switch message.Content.MessageContentType() {
	case client.TypeMessageText:
		handleText(message, senderID, senderUsername)
		// case client.TypeMessageAnimation:
		// 	handleMedia(message, client.TypeAnimation, false)
		// case client.TypeMessagePhoto:
		// 	handleMedia(message, client.TypePhoto, false)
		// case client.TypeMessageVideo:
		// 	handleMedia(message, client.TypeVideo, false)
	}

}

// handleUpdatedMessage handles updated messages.
func handleUpdatedMessage(umc *client.UpdateMessageContent) {

	// Since we are just given the updated content,
	// we need to get the full message
	message, err := api.GetMessage(umc.ChatId, umc.MessageId)
	if err != nil {
		log.Errorf("Unable to get message data: %+v", err)
		return
	}
	// log.Debugf("Message: %#v", message.Content)
	senderID, err := api.GetSenderID(message) //api.GetSenderUserID(message)
	if err != nil {
		log.Error("GetSenderID: %+v", err)
		return
	}

	// Tdlib delivers updates from self
	if senderID == repository.Me.Id {
		log.Debug("Message from self")
		return
	}

	senderUsername, _ := api.GetUsernameByID(senderID)
	// if err != nil {
	// 	log.Errorf("GetUsernameByID(%d): %+v", senderID, err)
	// }

	// Updates to textual messages can be handled normally, without any specific worry
	if umc.NewContent.MessageContentType() == client.TypeMessageText {

		if !message.IsChannelPost {
			handleText(message, senderID, senderUsername)
		}

		return
	}

}

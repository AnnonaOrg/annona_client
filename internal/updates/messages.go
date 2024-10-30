package updates

import (
	"github.com/AnnonaOrg/annona_client/internal/api"
	log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
)

// handleNewMessage handles incoming messages.
func handleNewMessage(message *client.Message) {
	switch message.Content.MessageContentType() {
	case client.TypeMessageText:
		handleText(message)
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

	// Updates to textual messages can be handled normally, without any specific worry
	if umc.NewContent.MessageContentType() == client.TypeMessageText {
		if !message.IsChannelPost {
			handleText(message)
		}

		return
	}

}

package api

import "github.com/zelenin/go-tdlib/client"

func SendMessageText(text string, chatID int64) (*client.Message, error) {
	formattedText, err := client.ParseTextEntities(&client.ParseTextEntitiesRequest{
		Text:      text,
		ParseMode: &client.TextParseModeMarkdown{Version: 2},
	})

	msg, err := tdlibClient.SendMessage(&client.SendMessageRequest{
		ChatId:          chatID,
		MessageThreadId: 0,
		ReplyTo:         nil,
		Options:         nil,
		ReplyMarkup:     nil,
		InputMessageContent: &client.InputMessageText{
			Text: formattedText,
			LinkPreviewOptions: &client.LinkPreviewOptions{
				IsDisabled: true,
			},
			ClearDraft: false,
		},
	})
	return msg, err
}

// GetMessage returns the client.Message with the input messageID
// in the input chatID.
func GetMessage(chatID, messageID int64) (*client.Message, error) {

	message, err := tdlibClient.GetMessage(&client.GetMessageRequest{
		ChatId:    chatID,
		MessageId: messageID,
	})

	return message, err

}

//func ReportMessage(chatID, messageID int64, reason client.ReportReason, text string) (*client.Ok, error) {
//	// var messageIds []int64
//	// messageIds = append(messageIds, messageID)
//	return tdlibClient.ReportChat(&client.ReportChatRequest{
//		ChatId:     chatID,
//		MessageIds: []int64{messageID},
//		Reason:     reason, //client.ReportReasonSpam{},
//		Text:       text,
//	})
//}

// GetMessageFormattedText returns the client.FormattedText structure for
// supported message types, nil otherwise.
func GetMessageFormattedText(mc client.MessageContent) *client.FormattedText {

	switch mc.MessageContentType() {
	case client.TypeMessageText:
		return mc.(*client.MessageText).Text
	case client.TypeMessagePhoto:
		return mc.(*client.MessagePhoto).Caption
	case client.TypeMessageAnimation:
		return mc.(*client.MessageAnimation).Caption
	case client.TypeMessageVideo:
		return mc.(*client.MessageVideo).Caption
	default:
		return nil
	}

}

// 获取消息链接
func GetMessageLink(chatID, messageID int64, mediaTimestamp int32, forAlbum bool, inMessageThread bool) (*client.MessageLink, error) {
	messageLink, err := tdlibClient.GetMessageLink(&client.GetMessageLinkRequest{
		ChatId:          chatID,
		MessageId:       messageID,
		MediaTimestamp:  mediaTimestamp,
		ForAlbum:        forAlbum,
		InMessageThread: inMessageThread,
	})

	return messageLink, err
}

// const (
// 	ContentModeText      = "text"
// 	ContentModeAnimation = "animation"
// 	ContentModeAudio     = "audio"
// 	ContentModeDocument  = "document"
// 	ContentModePhoto     = "photo"
// 	ContentModeVideo     = "video"
// 	ContentModeVoiceNote = "voiceNote"
// )

// func GetMessageFormattedTextEx(messageContent client.MessageContent) (*client.FormattedText, ContentMode) {
// 	var (
// 		formattedText *client.FormattedText
// 		contentMode   ContentMode
// 	)
// 	// TODO: как использовать switch для разблюдовки по приведению типа?
// 	if content, ok := messageContent.(*client.MessageText); ok {
// 		formattedText = content.Text
// 		contentMode = ContentModeText
// 	} else if content, ok := messageContent.(*client.MessagePhoto); ok {
// 		formattedText = content.Caption
// 		contentMode = ContentModePhoto
// 	} else if content, ok := messageContent.(*client.MessageAnimation); ok {
// 		formattedText = content.Caption
// 		contentMode = ContentModeAnimation
// 	} else if content, ok := messageContent.(*client.MessageAudio); ok {
// 		formattedText = content.Caption
// 		contentMode = ContentModeAudio
// 	} else if content, ok := messageContent.(*client.MessageDocument); ok {
// 		formattedText = content.Caption
// 		contentMode = ContentModeDocument
// 	} else if content, ok := messageContent.(*client.MessageVideo); ok {
// 		formattedText = content.Caption
// 		contentMode = ContentModeVideo
// 	} else if content, ok := messageContent.(*client.MessageVoiceNote); ok {
// 		formattedText = content.Caption
// 		contentMode = ContentModeVoiceNote
// 	} else {
// 		// client.MessageExpiredPhoto
// 		// client.MessageSticker
// 		// client.MessageExpiredVideo
// 		// client.MessageVideoNote
// 		// client.MessageLocation
// 		// client.MessageVenue
// 		// client.MessageContact
// 		// client.MessageDice
// 		// client.MessageGame
// 		// client.MessagePoll
// 		// client.MessageInvoice
// 		formattedText = &client.FormattedText{Entities: make([]*client.TextEntity, 0)}
// 		contentMode = ""
// 	}
// 	return formattedText, contentMode
// }

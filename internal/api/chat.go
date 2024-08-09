package api

import "github.com/zelenin/go-tdlib/client"

// GetChat retrieves a chat by its chatID.
// It may be required before the bot is able to send messages
// to a certain chat, if it hasn't received updates from it.
func GetChat(chatID int64) (*client.Chat, error) {
	return tdlibClient.GetChat(&client.GetChatRequest{ChatId: chatID})
}
func GetChatTitle(chatID int64) (string, error) {
	chat, err := GetChat(chatID)
	if err != nil {
		return "", err
	}
	chatTitle := ""
	if chat != nil && len(chat.Title) > 0 {
		chatTitle = chat.Title
	}
	return chatTitle, nil
}

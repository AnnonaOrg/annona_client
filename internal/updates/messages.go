package updates

import (
	"github.com/zelenin/go-tdlib/client"
)

// handleNewMessage handles incoming messages.
func handleNewMessage(message *client.Message) {
	switch message.Content.MessageContentType() {
	case client.TypeMessageText:
		handleText(message)
	}
}

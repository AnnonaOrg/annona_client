package updates

import (
	log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
)

// HandleUpdates handles incoming updates, dispatching them
// to the appropriate sub-handlers.
func HandleUpdates(listener *client.Listener) {

	for update := range listener.Updates {
		log.Tracef(
			"update.GetType(): %s, update.GetClass(): %s, Value: %#v",
			update.GetType(), update.GetClass(), update,
		)
		if update.GetClass() == client.ClassUpdate {
			switch update.GetType() {
			case client.TypeUpdateNewMessage:
				// log.Debugf("update.GetType():%s", client.TypeUpdateNewMessage)
				handleNewMessage(update.(*client.UpdateNewMessage).Message)
				// case client.TypeUpdateMessageContent:
				// 	log.Debugf("update.GetType():%s", client.TypeUpdateMessageContent)
				// 	// log.Debugln(update.(*client.UpdateMessageContent).NewContent.MessageContentType())
				// 	handleUpdatedMessage(update.(*client.UpdateMessageContent))
				// 	//handleUpdatedMessage(update.(*client.UpdateMessageContent).NewContent)
				// case client.TypeUpdateDeleteMessages:
				// 	log.Info("update.GetType()", client.TypeUpdateDeleteMessages)
				// 	handleNewDeletion(update.(*client.UpdateDeleteMessages))
				// default:
				// 	log.Tracef("update.GetType(): %s, Value: %#v", update.GetType(), update)
			}

		}

	}

}

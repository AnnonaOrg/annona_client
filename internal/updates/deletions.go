package updates

import (
	log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
)

// handleNewDeletion handles deletion notifications and marks
// deleted channel posts as deleted in the database.
func handleNewDeletion(messages *client.UpdateDeleteMessages) {
	log.Debugln("permanent deletions: ", messages.MessageIds)
}

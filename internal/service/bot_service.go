package service

import (
	"github.com/AnnonaOrg/annona_client/internal/db_data"
)

func AddBotIDToSet(botID int64) error {
	return db_data.AddBotIDToSet(botID)
}
func IsBotID(userID int64) bool {
	return db_data.IsBotID(userID)
}

package service

// import (
// 	"github.com/AnnonaOrg/annona_client/internal/db_data"
// 	"github.com/AnnonaOrg/annona_client/internal/log"
// )

// func IsPublicChat(chatID int64) bool {
// 	_, isOk, _ := db_data.GetNotPublicChat(chatID)
// 	return isOk
// }

// func SetNotPublicChat(chatID int64, value string) {
// 	// NotPublicChatMap.Store(chatID, value)
// 	if err := db_data.SetNotPublicChat(chatID, value); err != nil {
// 		log.Errorf("SetNotPublicChat(%d,%s): %v", chatID, value, err)
// 	}
// }

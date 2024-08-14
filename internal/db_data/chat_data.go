package db_data

// import (
// 	"fmt"
// 	"time"
// 	// "github.com/AnnonaOrg/annona_client/internal/log"
// )

// const (
// 	CHAT_NOT_PUBLIC_prefix = "CHAT_NOT_PUBLIC_prefix_"
// )

// func SetNotPublicChat(chatID int64, value string) error {
// 	return AddKeyValueWithExpiration(
// 		fmt.Sprintf("%s%d", CHAT_NOT_PUBLIC_prefix, chatID),
// 		value,
// 		time.Hour*1,
// 	)
// }

// // 存在记录 返回 值，true，
// func GetNotPublicChat(chatID int64) (string, bool, error) {
// 	value := ""

// 	if err := GetKeyValue(
// 		fmt.Sprintf("%s%d", CHAT_NOT_PUBLIC_prefix, chatID),
// 		&value,
// 	); err != nil {
// 		if err != NilErr {
// 			// log.Errorf("GetNotPublicChat(%d): %v", chatID, err)
// 			return "", false, err
// 		} else {
// 			return "", false, nil
// 		}
// 	}

// 	return value, true, nil
// }

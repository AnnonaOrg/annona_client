package service

import (
	"strings"

	"github.com/AnnonaOrg/annona_client/internal/api"
	"github.com/AnnonaOrg/annona_client/internal/db_data"
	"github.com/AnnonaOrg/annona_client/internal/log"
)

// const (
// 	USERNAME_null string = "NULL"
// )

func AddBotIDToSet(botID int64) error {
	return db_data.AddBotIDToSet(botID)
}
func IsBotID(userID int64) bool {
	isBot := db_data.IsBotID(userID)
	if isBot {
		return true
	}
	list := GetUsernames(userID)
	for _, v := range list {
		vc := strings.ToLower(v)
		if strings.HasSuffix(vc, "bot") {
			if err := AddBotIDToSet(userID); err != nil {
				log.Errorf("IsBotID.AddBotIDToSet(%d): %v", userID, err)
			}
			return true
		}
	}
	return false
}

func SetUsername(userID int64, usernames string) error {
	return db_data.SetUsername(userID, usernames)
}
func GetUsernames(userID int64) []string {
	// return db_data.GetUsername(userID)
	usernames := db_data.GetUsername(userID)
	if usernames == "" {
		if usernameList, err := api.GetUsernamesByID(userID); err != nil {
			usernames = "NULL"
			log.Errorf("GetUsername.GetUsernamesByID(%d): %v", userID, err)
			if err := SetUsername(userID, usernames); err != nil {
				log.Errorf("GetUsername.SetUsername(%d,%s): %v", userID, usernames, err)
			}
		} else {
			if len(usernameList) > 0 {
				usernames = strings.Join(usernameList, ",")
				if err := SetUsername(userID, usernames); err != nil {
					log.Errorf("GetUsername.SetUsername(%d,%s): %v", userID, usernames, err)
				}
				return usernameList
			}
		}
	}
	if usernames == "NULL" {
		return nil
	}
	var list []string
	if strings.Contains(usernames, ",") {
		list = strings.Split(usernames, ",")
	} else if len(usernames) > 0 {
		list = append(list, usernames)
	}
	return list
}

func GetUsername(userID int64) string {
	list := GetUsernames(userID)
	if len(list) > 0 {
		return list[0]
	}
	return ""
}

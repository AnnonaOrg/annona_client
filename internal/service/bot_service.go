package service

import (
	"strings"

	"github.com/AnnonaOrg/annona_client/internal/constvar"

	"github.com/AnnonaOrg/annona_client/internal/api"
	"github.com/AnnonaOrg/annona_client/internal/db_data"
	"github.com/AnnonaOrg/annona_client/internal/log"
)

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
	usernames := db_data.GetUsername(userID)
	if usernames == "" {
		var usernameList []string
		var err error
		if userID > 0 {
			usernameList, err = api.GetUsernamesByID(userID)
			if err != nil {
				log.Errorf("api.GetUsernamesByID(%d): %v", userID, err)
			}
		} else {
			usernameList, err = api.GetSupergroupUsernamesByID(userID)
			if err != nil {
				log.Errorf("api.GetSupergroupUsernamesByID(%d): %v", userID, err)
			}
		}

		if err != nil {
			usernames = constvar.REDIS_VALUE_NULL
			if err := SetUsername(userID, usernames); err != nil {
				log.Errorf("GetUsername.SetUsername(%d,%s): %v", userID, usernames, err)
			}
		}
		if len(usernameList) > 0 {
			usernames = strings.Join(usernameList, ",")
			if err := SetUsername(userID, usernames); err != nil {
				log.Errorf("GetUsername.SetUsername(%d,%s): %v", userID, usernames, err)
			}
			return usernameList
		}
	}
	if usernames == constvar.REDIS_VALUE_NULL {
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
	for _, v := range list {
		if len(v) > 0 {
			return v
		}
	}
	return ""
}

func SetUserFirstLastName(userID int64, firstLastName string) error {
	return db_data.SetUserFirstLastName(userID, firstLastName)
}
func GetUserFirstLastName(userID int64) string {
	firstLastName := db_data.GetUserFirstLastName(userID)
	if len(firstLastName) == 0 {
		if firstName, lastName, err := api.GetUserFirstLastName(userID); err != nil {
			// firstLastName = "NULL"
			log.Errorf("api.GetUserFirstLastName(%d): %v", userID, err)
		} else {
			firstLastName = firstName + " " + lastName
		}
		if err := SetUserFirstLastName(userID, firstLastName); err != nil {
			log.Errorf("SetUserFirstLastName(%d,%s): %v", userID, firstLastName, err)
		}
	}

	if firstLastName == constvar.REDIS_VALUE_NULL {
		return ""
	}
	return firstLastName
}

func GetChatTitle(chatID int64) string {
	chatTitle := db_data.GetUserFirstLastName(chatID)
	if len(chatTitle) == 0 {
		if titleTmp, err := api.GetChatTitle(chatID); err != nil {
			log.Errorf("GetChatTitle(%d): %v", chatID, err)
		} else {
			chatTitle = titleTmp
		}
		if err := SetUserFirstLastName(chatID, chatTitle); err != nil {
			log.Errorf("SetUserFirstLastName(%d,%s): %v", chatID, chatTitle, err)
		}
	}
	if chatTitle == constvar.REDIS_VALUE_NULL {
		return ""
	}
	return chatTitle
}

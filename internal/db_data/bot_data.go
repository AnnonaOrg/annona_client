package db_data

import (
	"fmt"
	"time"

	"github.com/AnnonaOrg/annona_client/internal/log"
)

const (
	BOT_ID_SET_prefix = "BOT_ID_SET_prefix_"
	USER_NAME_prefix  = "USER_NAME_prefix_"

	USER_FirstLastName_prefix = "USER_FirstLastName_prefix"
)

func AddBotIDToSet(botID int64) error {
	return AddToSet(
		BOT_ID_SET_prefix,
		botID,
	)
}
func IsBotID(userID int64) bool {
	isBot, err := IsMemberOfSet(BOT_ID_SET_prefix, userID)
	if err != nil {
		return false
	}
	return isBot
}

func SetUsername(userID int64, username string) error {
	if len(username) == 0 {
		return fmt.Errorf("the username is NULL")
	}
	if username == "NULL" {
		return AddKeyValueWithExpiration(
			fmt.Sprintf("%s%d", USER_NAME_prefix, userID),
			username,
			time.Hour*24,
		)
	}
	return AddKeyValue(
		fmt.Sprintf("%s%d", USER_NAME_prefix, userID),
		username,
	)
}
func GetUsername(userID int64) string {
	username := ""
	if err := GetKeyValue(
		fmt.Sprintf("%s%d", USER_NAME_prefix, userID),
		&username,
	); err != nil {
		if err != NilErr {
			log.Errorf("GetUsername(%d): %v", userID, err)
		}
	}

	return username
}

func SetUserFirstLastName(userID int64, firstLastName string) error {
	if len(firstLastName) == 0 {
		return fmt.Errorf("the firstLastName is NULL")
	}

	return AddKeyValue(
		fmt.Sprintf("%s%d", USER_FirstLastName_prefix, userID),
		firstLastName,
	)
}
func GetUserFirstLastName(userID int64) string {
	firstLastName := ""
	if err := GetKeyValue(
		fmt.Sprintf("%s%d", USER_FirstLastName_prefix, userID),
		&firstLastName,
	); err != nil {
		if err != NilErr {
			log.Errorf("GetUserFirstLastName(%d): %v", userID, err)
		}
	}

	return firstLastName
}

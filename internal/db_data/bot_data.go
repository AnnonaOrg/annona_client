package db_data

import (
	"fmt"
	"strings"

	"github.com/umfaka/umfaka_core/internal/log"
)

const (
	BOT_ID_SET_PREFIX
)

func AddBotIDToSet(botID int64) error {
	return AddToSet(
		BOT_ID_SET_PREFIX,
		botID,
	)
}
func IsBotID(userID int64) bool {
	isBot, err := IsMemberOfSet(BOT_ID_SET_PREFIX, botID)
	if err != nil {
		return false
	}
	return isBot
}

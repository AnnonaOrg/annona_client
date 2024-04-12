package api

import (
	"fmt"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
)

var botSyncMap sync.Map

// GetUserByID returns a client.User given a userID.
func GetUserByID(userID int64) (*client.User, error) {
	user, err := tdlibClient.GetUser(&client.GetUserRequest{UserId: userID})
	return user, err
}
func GetUsernamesByID(userID int64) ([]string, error) {
	var usernames []string

	user, err := tdlibClient.GetUser(&client.GetUserRequest{UserId: userID})
	if err != nil {
		return nil, fmt.Errorf("tdlibClient.GetUser(%d) %v", userID, err)
	}
	uMap := make(map[string]string, 0)
	if user != nil && user.Usernames != nil {
		for _, v := range user.Usernames.ActiveUsernames {
			if _, ok := uMap[v]; !ok {
				uMap[v] = v
				usernames = append(usernames, v)
			}
		}
		for _, v := range user.Usernames.DisabledUsernames {
			if _, ok := uMap[v]; !ok {
				uMap[v] = v
				usernames = append(usernames, v)
			}
		}
		if len(user.Usernames.EditableUsername) > 0 {
			v := user.Usernames.EditableUsername
			if _, ok := uMap[v]; !ok {
				uMap[v] = v
				usernames = append(usernames, v)
			}
		}
	}
	if len(usernames) > 0 {
		return usernames, nil
	}
	return nil, fmt.Errorf("the userID(%d) have no username", userID)
}

func GetUsernameByID(userID int64) (string, error) {
	usernames, err := GetUsernamesByID(userID)
	if err != nil {
		return "", err
	}
	for _, v := range usernames {
		if len(v) > 0 {
			return v, nil
		}
	}
	return "", fmt.Errorf("the userID(%d) have no username: %v", userID, usernames)
}

func CheckBotUsernameByUserID(userID int64) (string, bool, error) {
	usernames, err := GetUsernamesByID(userID)
	if err != nil {
		return "", false, fmt.Errorf("GetUsernamesByID(%d): %v", userID, err)
	}
	for _, v := range usernames {
		vc := strings.ToLower(v)
		if strings.HasSuffix(vc, "bot") {
			return v, true, nil
		}
	}
	return "", false, nil
}

func IsViaBotByUserID(userID int64, viaBotUserId int64) (bool, error) {
	if viaBotUserId != 0 {
		log.Debugf("ViaBotUserId(%d) Is a Bot ", viaBotUserId)
		return true, nil
	}

	userKey := fmt.Sprintf("%d", userID)
	if botUsername, ok := botSyncMap.Load(userKey); ok {
		log.Debugf("userID: %s Is a Bot userName: %s", userKey, botUsername)
		return true, nil
	}

	if botUsername, isBot, err := CheckBotUsernameByUserID(userID); err != nil {
		log.Debugf("CheckBotUsernameByUserID(%d) err : %v", userID, err)
		return false, err
	} else if isBot {
		botSyncMap.Store(userKey, botUsername)
		log.Debugf("%s Is a Bot : %s ", userKey, botUsername)
		return true, nil
	}
	return false, nil
}

func GetSenderID(message *client.Message) (int64, error) {
	switch message.SenderId.MessageSenderType() {
	case client.TypeMessageSenderUser:
		sender := message.SenderId.(*client.MessageSenderUser)
		return sender.UserId, nil
	case client.TypeMessageSenderChat:
		sender := message.SenderId.(*client.MessageSenderChat)
		return sender.ChatId, nil
	default:
		return 0, fmt.Errorf("the message was not sent by a chat or user")
	}
}

// GetSenderUserID returns the sender user id, if the message
// was not sent on behalf of a chat
func GetSenderUserID(message *client.Message) (int64, error) {

	if message.SenderId.MessageSenderType() != client.TypeMessageSenderUser {
		return 0, fmt.Errorf("the message was not sent by a user")
	}
	sender := message.SenderId.(*client.MessageSenderUser)
	return sender.UserId, nil
}
func GetSenderChatID(message *client.Message) (int64, error) {

	if message.SenderId.MessageSenderType() != client.TypeMessageSenderChat {
		return 0, fmt.Errorf("the message was not sent by a chat")
	}

	sender := message.SenderId.(*client.MessageSenderChat)
	return sender.ChatId, nil
}

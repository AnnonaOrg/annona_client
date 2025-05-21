package api

import (
	"fmt"
	"github.com/zelenin/go-tdlib/client"
)

// GetChat retrieves a chat by its chatID.
// It may be required before the bot is able to send messages
// to a certain chat, if it hasn't received updates from it.
func GetChat(chatID int64) (*client.Chat, error) {
	return tdlibClient.GetChat(&client.GetChatRequest{ChatId: chatID})
}
func GetSupergroup(supergroupID int64) (*client.Supergroup, error) {
	return tdlibClient.GetSupergroup(&client.GetSupergroupRequest{SupergroupId: supergroupID})
}

func JoinChatByInviteLink(inviteLink string) (*client.Chat, error) {
	return tdlibClient.JoinChatByInviteLink(&client.JoinChatByInviteLinkRequest{
		InviteLink: inviteLink,
	})
}

// JoinChat 加入一个会话(群组，频道)，私聊消息和加密会话消息不可用
func JoinChat(chatID int64) error {
	_, err := tdlibClient.JoinChat(&client.JoinChatRequest{ChatId: chatID})
	return err
}

// LeaveChat 退出会话(群组，频道),私聊消息和加密会话消息不可用
func LeaveChat(chatID int64) error {
	_, err := tdlibClient.LeaveChat(&client.LeaveChatRequest{ChatId: chatID})
	return err
}

func GetChatTitle(chatID int64) (string, error) {
	chat, err := GetChat(chatID)
	if err != nil {
		return "", fmt.Errorf("GetChat(%d) err: %v", chatID, err)
	}
	chatTitle := ""
	if chat != nil && len(chat.Title) > 0 {
		chatTitle = chat.Title
	}
	return chatTitle, nil
}

// IsCanGetMessageLink 判断是否为超级群，如果没有用户名，识别为不支持生成消息链接
func IsCanGetMessageLink(chatID int64) (bool, error) {
	chat, err := GetChat(chatID)
	if err != nil || chat == nil {
		return false, fmt.Errorf("GetChat(%d) err: %v", chatID, err)
	}

	// 判断类型和属性
	switch t := chat.Type.(type) {
	case *client.ChatTypeSupergroup:
		// 获取更详细的 supergroup 信息
		supergroup, err := GetSupergroup(t.SupergroupId)
		if err != nil || supergroup == nil {
			return false, fmt.Errorf("GetSupergroup(%d:%s) err: %v", t.SupergroupId, chat.Title, err)
		}

		// 判断是否是频道
		if t.IsChannel {
			if username := supergroup.Usernames; username != nil {
				if len(username.ActiveUsernames) > 0 || len(username.EditableUsername) > 0 {
					return true, nil
				} else {
					return false, fmt.Errorf("这是私有频道(%+v)，不能生成消息链接(%d:%s)", username, t.SupergroupId, chat.Title)
				}
			} else {
				return false, fmt.Errorf("这是私有频道，不能生成消息链接(%d:%s)", t.SupergroupId, chat.Title)
			}
		} else {
			if username := supergroup.Usernames; username != nil {
				if len(username.ActiveUsernames) > 0 || len(username.EditableUsername) > 0 {
					return true, nil
				} else {
					return false, fmt.Errorf("这是私有超级群(%+v)，不能生成消息链接(%d:%s)", username, t.SupergroupId, chat.Title)
				}
			} else {
				return false, fmt.Errorf("这是私有超级群，不能生成消息链接(%d:%s)", t.SupergroupId, chat.Title)
			}
		}

	default:
		return false, fmt.Errorf("不支持生成消息链接的 chat(%+v) 类型", t)
	}

}

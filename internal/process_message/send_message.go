package process_message

import (
	"fmt"

	"github.com/AnnonaOrg/osenv"
	"github.com/AnnonaOrg/pkg/send2server"
	log "github.com/sirupsen/logrus"
)

type FeedRichMsgModel struct {
	Msgtype      string                `json:"msgtype"  form:"msgtype"` //rich text image video
	MsgID        string                `json:"msgID"  form:"msgID"`
	MsgTime      string                `json:"msgTime"  form:"msgTime"`
	Text         FeedRichMsgTextModel  `json:"text"  form:"text"`
	Image        FeedRichMsgImageModel `json:"image"  form:"image"`
	Video        FeedRichMsgVideoModel `json:"video"  form:"video"`
	Link         string                `json:"link"  form:"link"`
	LinkIsPublic bool                  `json:"linkIsPublic"  form:"linkIsPublic"`

	BotInfo  FeedRichMsgBotInfoModel  `json:"botInfo"  form:"botInfo"`
	ChatInfo FeedRichMsgChatInfoModel `json:"chatInfo"  form:"chatInfo"`
	FormInfo FeedRichMsgFormInfoModel `json:"formInfo"  form:"formInfo"`
	NoButton bool                     `json:"noButton" form:"noButton"`
}
type FeedRichMsgTextModel struct {
	Content         string `json:"content"  form:"content"`
	ContentEx       string `json:"contentEx"  form:"contentEx"`
	ContentExPic    string `json:"contentExPic"  form:"contentExPic"`
	ContentMarkdown string `json:"contentMarkdown"  form:"contentMarkdown"`
}
type FeedRichMsgImageModel struct {
	PicURL   string `json:"picURL"  form:"picURL"`
	FilePath string `json:"filePath"  form:"filePath"`
	// (Optional)
	Caption string `json:"caption,omitempty"`
}
type FeedRichMsgVideoModel struct {
	FileURL  string `json:"fileURL"  form:"fileURL"`
	FilePath string `json:"filePath"  form:"filePath"`
	// (Optional)
	Caption string `json:"caption,omitempty"`
}
type FeedRichMsgBotInfoModel struct {
	BotToken string `json:"botToken" form:"botToken"`
}
type FeedRichMsgChatInfoModel struct {
	ToChatID int64 `json:"toChatID" form:"toChatID"`
}

type FeedRichMsgFormInfoModel struct {
	FormChatID   string `json:"formChatID" form:"formChatID"`
	FormSenderID string `json:"formSenderID" form:"formSenderID"`

	FormKeyworld string `json:"formKeyworld" form:"formKeyworld"`
}

func (msg *FeedRichMsgModel) ToString() (res string) {
	res = fmt.Sprintf(
		"msgID:%s\n"+"msgType:%s\n"+"msgTime:%s\n"+"toChatID:%d",
		msg.MsgID, msg.Msgtype, msg.MsgTime, msg.ChatInfo.ToChatID,
	)
	if len(msg.Text.Content) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Text.Content)
	}
	if len(msg.Image.PicURL) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Image.PicURL)
	}
	if len(msg.Video.FileURL) > 0 {
		res = fmt.Sprintf("%s\n%s", res, msg.Video.FileURL)
	}

	return
}

// processMessageKeywords 处理消息关键词
// botToken 机器人Bot
// toChatID 发送到ID
// messageID 消息ID
// messageContentText 消息内容(处理过的)
func sendMessage(
	chatID, senderID string, toChatID int64, botToken,
	messageID, messageDate, messageContentText string, messageLink string, messageLinkIsPublic bool,
	keyworld string,
) (string, error) {
	var richMsg FeedRichMsgModel

	richMsg.BotInfo.BotToken = botToken
	richMsg.ChatInfo.ToChatID = toChatID
	richMsg.MsgID = fmt.Sprintf("%d", toChatID) + messageID
	richMsg.MsgTime = messageDate
	richMsg.Msgtype = "rich"
	richMsg.Text.Content = messageContentText
	richMsg.FormInfo.FormChatID = chatID
	richMsg.FormInfo.FormSenderID = senderID
	richMsg.FormInfo.FormKeyworld = keyworld
	richMsg.NoButton = false
	richMsg.Link = messageLink
	richMsg.LinkIsPublic = messageLinkIsPublic

	serverRouter := osenv.GetNoticeOfFeedRichMsgPushUrl()
	serverChannel := fmt.Sprintf("%d", toChatID)
	serverToken := fmt.Sprintf("%d", toChatID)
	serverPath := osenv.GetNoticeOfFeedRichMsgPushUrlPath()

	retText, err := send2server.SendMsgToServer(
		richMsg,
		serverRouter,
		serverChannel,
		serverToken,
		serverPath,
	)
	log.Debugf("sendMessage(%+v),serverRouter(%s),serverChannel(%s),serverToken(%s),serverPath(%s): %s , %v",
		richMsg,
		serverRouter, serverChannel, serverToken, serverPath,
		retText, err,
	)
	return retText, err

	// return send2server.SendMsgToServer(
	// 	richMsg,
	// 	serverRouter,
	// 	serverChannel,
	// 	serverToken,
	// 	serverPath,
	// )
}

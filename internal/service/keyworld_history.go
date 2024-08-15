package service

import (
	"encoding/json"
	"fmt"

	"github.com/AnnonaOrg/annona_client/internal/log"
	"github.com/AnnonaOrg/annona_client/internal/request"
	"github.com/AnnonaOrg/annona_client/internal/response"
	"github.com/AnnonaOrg/annona_client/utils"
	"github.com/AnnonaOrg/osenv"
)

func CreateKeyworldHistory(req *request.KeyworldHistoryInfo) error {
	apiDomain := osenv.GetCoreApiUrl()
	apiToken := osenv.GetCoreApiToken()
	apiPath := "/apis/v1/keyword_history/item/add"

	retBody, err := utils.DoPostJsonToOpenAPI(apiDomain, apiPath, apiToken, req)
	if err != nil {
		log.Errorf("DoPostJsonToOpenAPI(%s,%s,%s,%+v): %v", apiDomain, apiPath, apiToken, req, err)
		return fmt.Errorf("服务请求出错: %v", err)
	}
	var apiResponse response.Response

	err = json.Unmarshal(retBody, &apiResponse)
	if err != nil {
		log.Errorf("Unmarshal(%s): %v", string(retBody), err)
		return fmt.Errorf("解析信息: %s ,失败: %v", string(retBody), err)
	}
	if apiResponse.Code != 0 {
		log.Errorf("响应状态异常: %+v", apiResponse)
		return fmt.Errorf("err msg: %s", apiResponse.Message)
	}
	return nil
}
func CreateKeyworldHistoryEx(
	chatID, senderID int64, senderUsername string,
	messageID int64, messageContentText string, messageLink string,
	keyworld string,
	messageDateStr string,
	messageDate int64,
) error {
	req := &request.KeyworldHistoryInfo{}
	req.ChatId = chatID
	req.SenderId = senderID
	req.SenderUsername = senderUsername
	req.MessageId = messageID
	req.MessageContentText = messageContentText
	req.MessageLink = messageLink
	req.KeyWorld = keyworld
	req.MessageDateEx = messageDateStr
	req.MessageDate = messageDate
	return CreateKeyworldHistory(req)
}

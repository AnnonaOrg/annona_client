package updates

import (
	log "github.com/sirupsen/logrus"
	"github.com/zelenin/go-tdlib/client"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// HandleUpdates handles incoming updates, dispatching them
// to the appropriate sub-handlers.
func HandleUpdates(listener *client.Listener) {

	for update := range listener.Updates {
		//log.Tracef(
		//	"update.GetType(): %s, update.GetClass(): %s, Value: %#v",
		//	update.GetType(), update.GetClass(), update,
		//)
		//if update.GetType() == "error" {
		//	errObj := update.(*client.Error)
		//	if strings.HasPrefix(errObj.Message, "FLOOD_WAIT_") {
		//		log.Errorf("捕获到 FLOOD_WAIT: %v", errObj.Message)
		//
		//		// 提取等待秒数
		//		re := regexp.MustCompile(`FLOOD_WAIT_(\d+)`)
		//		match := re.FindStringSubmatch(errObj.Message)
		//		if len(match) == 2 {
		//			waitSec, _ := strconv.Atoi(match[1])
		//			log.Errorf("建议等待 %d 秒再重试", waitSec)
		//		}
		//	}
		//}
		//if update.GetClass() == client.ClassUpdate {
		//	switch update.GetType() {
		//	case client.TypeUpdateNewMessage:
		//		// log.Debugf("update.GetType():%s", client.TypeUpdateNewMessage)
		//		handleNewMessage(update.(*client.UpdateNewMessage).Message)
		//		// case client.TypeUpdateMessageContent:
		//		// 	log.Debugf("update.GetType():%s", client.TypeUpdateMessageContent)
		//		// 	// log.Debugln(update.(*client.UpdateMessageContent).NewContent.MessageContentType())
		//		// 	handleUpdatedMessage(update.(*client.UpdateMessageContent))
		//		// 	//handleUpdatedMessage(update.(*client.UpdateMessageContent).NewContent)
		//		// case client.TypeUpdateDeleteMessages:
		//		// 	log.Info("update.GetType()", client.TypeUpdateDeleteMessages)
		//		// 	handleNewDeletion(update.(*client.UpdateDeleteMessages))
		//		// default:
		//		// 	log.Tracef("update.GetType(): %s, Value: %#v", update.GetType(), update)
		//	}
		//
		//}

		switch update.GetClass() {
		case client.ClassError:
			handleError(update.(*client.Error))
		case client.ClassUpdate:
			switch update.GetType() {
			case client.TypeUpdateNewMessage:
				handleNewMessage(update.(*client.UpdateNewMessage).Message)
			default:
				log.Tracef("未处理的 update 类型: %s", update.GetType())
			}
		}

	}

}
func handleError(errObj *client.Error) {
	if strings.HasPrefix(errObj.Message, "FLOOD_WAIT_") {
		log.Errorf("捕获到 FLOOD_WAIT 错误: %v", errObj.Message)

		waitSec, ok := extractFloodWaitSeconds(errObj.Message)
		if ok {
			log.Warnf("正在等待 %d 秒...", waitSec)
			time.Sleep(time.Duration(waitSec) * time.Second)
			log.Infof("等待结束，继续执行")
		} else {
			log.Warn("未能提取 FLOOD_WAIT 秒数")
		}
	} else {
		log.Errorf("TDLib 错误 (%d): %s", errObj.Code, errObj.Message)
	}
}

// 提取 FLOOD_WAIT_xx 的通用函数：
func extractFloodWaitSeconds(msg string) (int, bool) {
	re := regexp.MustCompile(`FLOOD_WAIT_(\d+)`)
	match := re.FindStringSubmatch(msg)
	if len(match) == 2 {
		if sec, err := strconv.Atoi(match[1]); err == nil {
			return sec, true
		}
	}
	return 0, false
}

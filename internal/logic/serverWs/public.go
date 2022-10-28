package serverWs

import (
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func SendNotificationToUser(from, to uint, data string) error {
	logx.Info("send notification...")
	message := &Message{
		OpCode:       OpText,
		Type:         SYSTEM,
		MessageID:    "",
		GroupID:      0,
		ToUser:       to,
		UserID:       from,
		Content:      data,
		SendTime:     time.Now().Unix(),
		GroupMembers: nil,
	}
	globalHub.broadcast <- message
	return nil
}
func SendNotificationToUserWithUserInfo(to uint, fromUserInfo *models.User, data string) error {
	logx.Info("send add friend notification...")
	message := &Message{
		OpCode:       OpText,
		Type:         SYSTEM,
		MessageID:    "",
		GroupID:      0,
		ToUser:       to,
		UserID:       fromUserInfo.ID,
		Content:      data,
		SendTime:     time.Now().Unix(),
		GroupMembers: nil,
		UserDetail: SenderData{
			UserID:     fromUserInfo.ID,
			UserName:   fromUserInfo.Name,
			UserAvatar: fromUserInfo.Avatar,
		},
	}
	globalHub.broadcast <- message
	return nil
}

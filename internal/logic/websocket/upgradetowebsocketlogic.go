package websocket

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"net/http"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

var webSocketUpgrade = websocket.Upgrader{}

type UpgradeToWebSocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpgradeToWebSocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpgradeToWebSocketLogic {
	return &UpgradeToWebSocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpgradeToWebSocketLogic) UpgradeToWebSocket(req *types.UpgradeToWebSocketReq, w http.ResponseWriter, r *http.Request) (resp *types.UpgradeToWebSocketResp, err error) {
	// todo: add your logic here and delete this line
	//Upgrade http to websocket
	conn, err := webSocketUpgrade.Upgrade(w, r, nil)
	if err != nil {
		return nil, errx.NewErrCode(errx.WEBSOCKET_CONNECTION_ERROR)
	}

	logx.Info("HTTP Upgrade :%+v", conn)

	defer conn.Close()
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			logx.Info("read error :%v", err)
			return nil, errx.NewCommonMessage(errx.WEBSOCKET_READ_MESSAG_ERROR, err.Error())
		}

		replyMessage := fmt.Sprintf("i got your message :%s", msg)
		err = conn.WriteMessage(messageType, []byte(replyMessage))
		if err != nil {
			return nil, errx.NewCommonMessage(errx.WEBSOCKET_WRITE_MESSAGE_ERROR, err.Error())
		}
	}
}

package websocket

import (
	"context"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/common/errx"
	"net/http"

	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/ryantokmanmokmtm/movie-server/internal/types"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

var UserConn = make(map[uint]*websocket.Conn)

type webSocketMsgRec struct {
	Msg    string `json:"msg"`
	ToUser uint   `json:"user_id"`
}

type webSocketMsgResp struct {
	Msg      string `json:"msg"`
	FromUser uint   `json:"user_id"`
}

type UpgradeToWebSocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	req    *http.Request
	resp   http.ResponseWriter
}

func NewUpgradeToWebSocketLogic(ctx context.Context, svcCtx *svc.ServiceContext, req *http.Request, resp http.ResponseWriter) *UpgradeToWebSocketLogic {
	return &UpgradeToWebSocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		req:    req,
		resp:   resp,
	}
}

func (l *UpgradeToWebSocketLogic) UpgradeToWebSocket(req *types.UpgradeToWebSocketReq) (resp *types.UpgradeToWebSocketResp, err error) {
	// todo: add your logic here and delete this line
	//Getting UserInfo from JWT/Request
	//check user in db
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindUserByID(l.ctx, userId)
	if err != nil {
		return nil, errx.NewErrCode(errx.USER_NOT_EXIST)
	}

	//upgrade http request to websocket
	//store websocket connection with its userId
	upgrade := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrade.Upgrade(l.resp, l.req, nil)
	if err != nil {
		return nil, errx.NewErrCode(errx.WEBSOCKET_CONNECTION_ERROR)
	}
	defer conn.Close()
	logx.Info("user %d is connected.\n")
	UserConn[userId] = conn
	for {
		var msg webSocketMsgRec
		err := conn.ReadJSON(&msg)
		if err != nil {
			return nil, errx.NewErrCode(errx.WEBSOCKET_READ_MESSAG_ERROR)
		}
		//insert message into db/ history

		if _, ok := UserConn[msg.ToUser]; ok {
			logx.Info("User %d send msg %s to user %d\n", userId, msg.Msg, msg.ToUser)
			err = conn.WriteJSON(&webSocketMsgResp{
				Msg:      "Received your message",
				FromUser: userId,
			})

			if err != nil {
				return nil, errx.NewErrCode(errx.WEBSOCKET_WRITE_MESSAGE_ERROR)
			}
		}
	}
}

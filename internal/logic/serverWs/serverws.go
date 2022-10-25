package serverWs

import (
	"github.com/gorilla/websocket"
	"github.com/ryantokmanmokmtm/movie-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

var globalHub *ChannelMap

type OpCode byte

const (
	OpContinuation OpCode = 0x0
	OpText         OpCode = 0x1
	OpBinary       OpCode = 0x2
	OpClose        OpCode = 0x8
	OpPing         OpCode = 0x9
	OpPong         OpCode = 0xa
)

const (
	SYSTEM = iota
	MESSAGE
)

type SenderData struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
}

type MessageReq struct {
	OpCode    OpCode `json:"opcode"`
	GroupID   uint   `json:"group_id"`
	MessageID string `json:"message_id"`
	Message   string `json:"message"`
	SentTime  int64  `json:"sent_time"`
}

type Message struct {
	OpCode       OpCode     `json:"opcode"`
	Type         uint       `json:"message_type"` //system , message , ping ,pong
	MessageID    string     `json:"message_id"`   //same as sender
	GroupID      uint       `json:"group_id"`     //for chat
	GroupMembers []uint     `json:"-"`
	ToUser       uint       `json:"to_user"` //for notification
	UserID       uint       `json:"user_id"`
	UserDetail   SenderData `json:"sender_info"`
	Content      string     `json:"content"`
	SendTime     int64      `json:"send_time"` //same as sender
}

const (
	ReadWait  = 60
	WriteWait = 20
	ReadLimit = 1024
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewServerWS(svcCtx *svc.ServiceContext) func(w http.ResponseWriter, r *http.Request) {
	if globalHub == nil {
		logx.Info("initialing hub...")
		globalHub = NewChannelMap()
		go globalHub.Run()
	}

	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("Req came in...")
		ServerWS(svcCtx, globalHub, w, r)
	}
}

func ServerWS(ctx *svc.ServiceContext, hub *ChannelMap, w http.ResponseWriter, r *http.Request) {
	logx.Info("Trying to connect...")
	userId := ctxtool.GetUserIDFromCTX(r.Context())
	if userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := ctx.DAO.FindUserByID(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logx.Infof("Client %d(name:%s) is connected via websocket", userId, u.Name)
	client := NewClientConn(userId, conn, hub, ctx)
	hub.register <- client
	//hub.broadcast <- &Message{
	//	Type:         SYSTEM,
	//	GroupID:      0,
	//	ToUser:       0,
	//	UserID:       u.ID,
	//	Content:      fmt.Sprintf("[SYSTEM] %s is now online.", u.Name),
	//	SendTime:     time.Now().Unix(),
	//	GroupMembers: nil,
	//}

	go client.ReadLoop()
	go client.WriteLoop()

}

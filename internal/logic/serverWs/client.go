package serverWs

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/ryantokmanmokmtm/movie-server/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

type ClientConn struct {
	sync.Mutex
	once sync.Once
	hub  *ChannelMap

	UserID   uint
	conn     *websocket.Conn
	svcCtx   *svc.ServiceContext
	send     chan []byte
	isClosed chan struct{}
}

func NewClientConn(userID uint, conn *websocket.Conn, hub *ChannelMap, svcCtx *svc.ServiceContext) *ClientConn {
	return &ClientConn{
		hub:      hub,
		UserID:   userID,
		conn:     conn,
		send:     make(chan []byte),
		svcCtx:   svcCtx,
		isClosed: make(chan struct{}),
	}
}

func (c *ClientConn) ReadLoop() {
	defer func() {
		c.hub.unRegister <- c // remove client from map
		c.conn.Close()        //close connection
	}()

	c.conn.SetReadDeadline(time.Now().Add(time.Second * ReadWait))
	c.conn.SetReadLimit(ReadLimit)
	//c.conn.SetPongHandler(func(string) error {
	//	c.conn.SetReadDeadline(time.Now().Add(time.Second * ReadWait))
	//
	//	return nil
	//})

	for {
		//get data from connection
		//c.conn.SetReadDeadline(time.Now().Add(time.Second * ReadWait))
		_, msg, err := c.conn.ReadMessage()
		ctx := context.Background()
		if err != nil {
			logx.Error(err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logx.Error(err)
			}
			break
		}

		req := &MessageReq{}
		err = json.Unmarshal(msg, req)
		if err != nil {
			logx.Error(err)
			//there may send back an error message
			c.conn.SetReadDeadline(time.Now().Add(time.Second * ReadWait))
			continue
		}

		if req.OpCode == OpPong {
			//send the pong message
			logx.Info("received a pong message from client")
			c.conn.SetReadDeadline(time.Now().Add(time.Second * ReadWait))
			continue
		}

		if req.OpCode == OpPing {
			logx.Info("received a ping message from client")
			c.hub.broadcast <- &Message{
				OpCode: OpPong,
			}
			continue
		}

		u, err := c.svcCtx.DAO.FindUserByID(ctx, c.UserID)
		if err != nil {
			logx.Error(err)
			continue
		}
		//TODO: Get Room ID From JSON
		if err := c.svcCtx.DAO.ExistInTheRoom(ctx, c.UserID, req.GroupID); err != nil {
			logx.Error(err)
			continue
		}
		//TODO: Check User
		//TODO: Store Message
		if err := c.svcCtx.DAO.InsertOneMessage(ctx, req.GroupID, c.UserID, req.Message, req.MessageID, req.SentTime); err != nil {
			logx.Error(err)
			continue
		}

		if err := c.svcCtx.DAO.UpdateIsRead(ctx, req.GroupID, false); err != nil {
			logx.Error(err)
			continue
		}

		//TODO: Update the last_sender
		if err := c.svcCtx.DAO.UpdateLastSender(ctx, req.GroupID, c.UserID); err != nil {
			logx.Error(err)
			continue
		}

		//TODO: send the message to all user to all room user who is online
		allUser, err := c.svcCtx.DAO.GetRoomUsers(ctx, req.GroupID)
		if err != nil {
			logx.Error(err)
			continue
		}

		//TODO: Update user active room state - if user room is not active
		//Set both user to active
		//for now just one user there

		//update all user's room state
		for _, urs := range allUser { //set all user to active for now...including sender itself
			if err := c.svcCtx.DAO.UpdateRoomActiveState(ctx, req.GroupID, urs, true); err != nil {
				logx.Error(err)
				continue
			}
		}

		message := &Message{
			OpCode:    req.OpCode,
			Type:      MESSAGE,
			GroupID:   req.GroupID,
			MessageID: req.MessageID,
			ToUser:    0,        //send the message to this user . currently this message is for group chat...
			UserID:    c.UserID, //message is from this user
			UserDetail: SenderData{ //TOOD: do we need this data???
				UserID:     u.ID,
				UserName:   u.Name,
				UserAvatar: u.Avatar,
			},
			Content:      req.Message,
			SendTime:     req.SentTime,
			GroupMembers: allUser,
		}

		c.hub.broadcast <- message
	}
}

func (c *ClientConn) WriteLoop() {
	t := time.NewTicker(time.Second * (ReadWait * 9 / 10))
	defer func() {
		c.hub.unRegister <- c // remove client from map
		c.conn.Close()        //close connection
	}()

	for {
		select {
		case data, ok := <-c.send:
			/*
				TODO:
				Response:
				1. Type of data - system or message
				2. UserSent
				3. Content ： message

			*/

			//set  write deadline and send
			c.conn.SetWriteDeadline(time.Now().Add(time.Second * WriteWait))
			if !ok {
				logx.Error("send channel is closed")
				c.conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logx.Error(err)
				return
			}

			_, _ = w.Write(data)
			n := len(c.send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(data)
			}

			if err := w.Close(); err != nil {
				logx.Error(err)
				return
			}

		case <-t.C:
			//TODO: if connection is left -> break

			logx.Info("send a ping message")
			c.conn.SetWriteDeadline(time.Now().Add(time.Second * WriteWait))
			//send a ping message
			pingMessage := Message{
				OpCode: OpPing,
				Type:   SYSTEM,
				ToUser: c.UserID, //send a ping to this user
			}

			data, _ := json.Marshal(pingMessage)
			c.conn.WriteMessage(websocket.TextMessage, data)
		case <-c.isClosed:
			logx.Info("user is disconnected")
			return
		}

	}
}

func (c *ClientConn) Close() {
	c.once.Do(func() {
		logx.Info("Connected is closed")
		c.isClosed <- struct{}{}
	})
}

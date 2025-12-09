package ws

import (
	"chatapp/backend/dto"
	"chatapp/backend/services"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

type IMessageClient interface {
	CreateMessage(service services.IMessageService)
	GetRoomMessages(service services.IMessageService)
}

type MessageClient struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func newMessageClient(hub *Hub, conn *websocket.Conn) *MessageClient {
	return &MessageClient{hub: hub, conn: conn, send: make(chan []byte, 256)}
}

func (c *MessageClient) CreateMessage(service services.IMessageService) {
	//読み取り処理
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var input dto.MessageInput
		if err := json.Unmarshal(message, &input); err != nil {
			log.Printf("Invalid message format: %v", err)
			c.conn.WriteMessage(websocket.TextMessage, []byte(`{"error":"Invalid message format"}`))
			continue
		}

		newMessage, err := service.CreateMessage(4, input.Content, input.RoomID)
		if err != nil {
			errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})
			c.conn.WriteMessage(websocket.TextMessage, errMsg)
			continue
		}

		// メッセージを broadcast
		data, _ := json.Marshal(newMessage)
		c.hub.broadcast <- data
	}
}

func (c *MessageClient) GetRoomMessages(service services.IMessageService) {
	ticker := time.NewTicker(pingPeriod)
	//hubから取得する処理
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

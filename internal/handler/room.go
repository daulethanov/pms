package handler

import (
	"net/http"
	"github.com/gorilla/websocket"
)

type RoomHandlerInterface interface{
	RoomMessage(w http.ResponseWriter, r *http.Request)
	// EnteredByInviteLink(w http.ResponseWriter, r *http.Request)
}

type RoomHandler struct{

}

func NewRoomHandler() RoomHandlerInterface{
	return &RoomHandler{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
 

func (a *RoomHandler) RoomMessage(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }
    defer conn.Close()
	
	for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            break
        }

        if err := conn.WriteMessage(messageType, message); err != nil {
            break
        }
    }
}


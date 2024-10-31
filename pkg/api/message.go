package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var sessions = make(map[string]chan string)

func MessageService(c *gin.Context) {

	key := uuid.NewString()

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		c.Abort()
		return
	}
	defer ws.Close()

	sendCh := make(chan string, 10)
	defer close(sendCh)

	sessions[key] = sendCh

	go eventProc(sendCh, ws)

	for {
		mtype, data, err := ws.ReadMessage()

		if err != nil {
			log.Println("receving error!: ", err)
			break
		}

		if mtype == 1 { // TEXT FRAME
			Broadcast(string(data))
		}
	}

	delete(sessions, key)
}

func eventProc(respCh chan string, ws *websocket.Conn) {
	for p := range respCh {
		ws.WriteMessage(1, []byte(p))
	}
}

func Broadcast(message string) {
	for _, s := range sessions {
		s <- message
	}
}

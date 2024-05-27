package handlers

import (
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/fsmardani/go-for-example/database"

)

func WebSocketHandler(c *websocket.Conn) {
	defer c.Close()
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		// Save message to MongoDB
		err = database.SaveMessage(string(msg))
		if err != nil {
			log.Println("MongoDB insert error:", err)
			continue
		}

		// Echo message back to the client
		if err = c.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
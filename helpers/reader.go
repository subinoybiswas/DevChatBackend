package helpers

import (
	"log"

	"github.com/gorilla/websocket"
)

func Reader(conn *websocket.Conn) {
	defer func() {
		delete(Clients, conn)
		conn.Close()
	}()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Log the received message for clarity
		log.Printf("Received: %s", string(p))

		// Broadcast the received message to all Clients except the sender
		for client := range Clients {
			if client != conn {
				if err := client.WriteMessage(websocket.TextMessage, p); err != nil {
					log.Printf("Error writing message: %v", err)
					client.Close()
					delete(Clients, client)
				}
			}
		}
	}
}
package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/subinoybiswas/devchat/helpers"
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	// Upgrade this connection to a WebSocket
	ws, err := helpers.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Register new client
	helpers.Clients[ws] = true

	log.Println("Client Connected")
	err = ws.WriteMessage(websocket.TextMessage, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	helpers.Reader(ws)
}

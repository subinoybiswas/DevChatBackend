package routes

import (
	"net/http"

	"github.com/subinoybiswas/devchat/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/ws", controllers.WsEndpoint)
}
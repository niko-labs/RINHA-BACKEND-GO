package routes

import (
	"log"
	"net/http"
)

const ROTA_PING = "/ping"

func Ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping")
	w.Write([]byte("Pong"))
	return
}

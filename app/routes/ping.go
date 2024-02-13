package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ROTA_PING = "/ping"

func Ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping")
	w.Write([]byte("Pong"))
	return
}

func (r RotaBase) RotaPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

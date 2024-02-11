package main

import (
	"log"
	"net/http"
	"os"
	"rinha-backend-2024-q1/database"
	"rinha-backend-2024-q1/helpers"
	"rinha-backend-2024-q1/routes"
)

func init() {
	log.Println("Carregando Ambiente")
	helpers.LoadEnv()
	database.ConectarAoPostgreSQL()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(routes.ROTA_TRANSACOES, routes.Transacoes)
	mux.HandleFunc(routes.ROTA_EXTRATO, routes.Extrato)
	mux.HandleFunc(routes.ROTA_PING, routes.Ping)

	SERVER_ADDR := os.Getenv("SERVER_ADDR")
	SERVER_PORT := os.Getenv("SERVER_PORT")
	SERVER := SERVER_ADDR + ":" + SERVER_PORT

	log.Println("Server running on: ", SERVER)
	http.ListenAndServe(SERVER, mux)

}

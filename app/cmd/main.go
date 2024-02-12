package main

import (
	"log"
	"os"
	"rinha-backend-2024-q1/database"
	"rinha-backend-2024-q1/helpers"
	"rinha-backend-2024-q1/repositories"
	"rinha-backend-2024-q1/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	helpers.PrintArt()
	helpers.LoadEnv()
}

func main() {

	db := database.ConectarBancoDados()
	defer db.Close()

	repos := repositories.Iniciar(db)
	rotas := routes.Iniciar(repos)

	r := gin.Default()

	r.GET(routes.ROTA_PING, rotas.RotaPing)
	r.GET(routes.ROTA_EXTRATO, rotas.ConsultarExtrato)
	r.POST(routes.ROTA_TRANSACOES, rotas.RealizarTransacao)

	SERVER_ADDR := os.Getenv("SERVER_ADDR")
	SERVER_PORT := os.Getenv("SERVER_PORT")
	SERVER := SERVER_ADDR + ":" + SERVER_PORT

	log.Println("Server running on: ", SERVER)
	r.Run(SERVER)

}
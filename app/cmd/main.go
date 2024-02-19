package main

import (
	"log"
	"os"
	"rinha-backend-2024-q1/database"
	"rinha-backend-2024-q1/helpers"
	"rinha-backend-2024-q1/repositories"
	"rinha-backend-2024-q1/routes"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

const (
	MAX_MEMORY = 64 * 1024 * 1024
)

func init() {
	helpers.PrintArt()
	helpers.LoadEnv()
	helpers.PrintServerMode()

}

func main() {
	debug.SetGCPercent(300)
	debug.SetMaxStack(MAX_MEMORY)

	// f, _ := os.Create("trace.out")
	// trace.Start(f)
	// defer trace.Stop()

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

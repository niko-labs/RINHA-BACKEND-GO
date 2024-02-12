package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	dir, err := os.Getwd()

	if err != nil {
		log.Println("Erro ao pegar o diretório atual")
		panic(err)
	}
	dir = dir + "/.env"

	log.Println("Buscando arquivo .env em: ", dir)

	err = godotenv.Load(dir)
	if err != nil {
		log.Println("Error loading .env file")
		// panic(err)
	}
}

func PrintArt() {
	art := `
    __     ____ ______ ______   _____  _   __ ___     __ __  ______
   / /    /  _// ____// ____/  / ___/ / | / //   |   / //_/ / ____/
  / /     / / / /_   / __/     \__ \ /  |/ // /| |  / ,<   / __/   
 / /___ _/ / / __/  / /___    ___/ // /|  // ___ | / /| | / /___   
/_____//___//_/    /_____/   /____//_/ |_//_/  |_|/_/ |_|/_____/   
`
	log.Println("\n", art)
}

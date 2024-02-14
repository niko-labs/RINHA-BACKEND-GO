package helpers

import (
	"log"
	"os"
)

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

func PrintServerMode() {
	msg := "Executando em modo: "
	if os.Getenv("GIN_MODE") == "release" {
		log.Println(msg, "release")
	} else {
		log.Println(msg, "debug")
	}
}

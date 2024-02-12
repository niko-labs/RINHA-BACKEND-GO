package database

import (
	"database/sql"

	"log"

	_ "github.com/lib/pq"

	"os"
)

var (
	dbCnx *sql.DB
)

func ConectarAoPostgreSQL() {

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := "host=" + DB_HOST + " port=" + DB_PORT + " user=" + DB_USER + " password=" + DB_PASS + " dbname=" + DB_NAME + " sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("Error on open database connection")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("Erro ao pingar o banco de dados")
		panic(err)
	}

	dbCnx = db
	log.Println("Banco de dados conectado com sucesso")
}

func PegarConexao() *sql.DB {
	if dbCnx == nil {
		panic("Conexão com o banco de dados não configurada")
	}
	return dbCnx
}

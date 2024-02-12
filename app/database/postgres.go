package database

import (
	"context"
	"rinha-backend-2024-q1/helpers"

	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"

	"os"
)

func ConectarBancoDados() *pgxpool.Pool {

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")

	DATABASE_URL := "postgres://" + DB_USER + ":" + DB_PASS + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_NAME + "?sslmode=disable"
	log.Println(DATABASE_URL)

	db, err := pgxpool.New(context.Background(), DATABASE_URL)
	helpers.VerificaErro(err)

	err = db.Ping(context.Background())
	helpers.VerificaErroComMsgLog(err, "Erro ao conectar ao banco de dados")

	log.Println("Banco de dados conectado com sucesso")
	return db
}

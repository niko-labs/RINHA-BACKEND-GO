package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type RepositorioBase struct {
	db *pgxpool.Pool
}

func Iniciar(db *pgxpool.Pool) *RepositorioBase {
	return &RepositorioBase{db: db}
}

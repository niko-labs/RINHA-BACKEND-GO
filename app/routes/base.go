package routes

import "rinha-backend-2024-q1/repositories"

type RotaBase struct {
	repo *repositories.RepositorioBase
}

func Iniciar(repo *repositories.RepositorioBase) *RotaBase {
	return &RotaBase{repo: repo}
}

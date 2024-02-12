package repositories

import (
	"context"
)

func (r *RepositorioBase) ObterExtrato(ctx context.Context, id int) (string, error) {
	resultado := r.db.QueryRow(ctx, Q_EXTRATO_CLIENTE, id)
	var extrato string
	err := resultado.Scan(&extrato)
	if err != nil {
		return "", err
	}
	return extrato, nil

}

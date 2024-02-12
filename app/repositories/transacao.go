package repositories

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *RepositorioBase) ExecutarTransacao(ctx context.Context, id int, valor int64, tipo string, descricao string) (*int64, *int64, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback(ctx)

	var limite, saldo int64

	if err = tx.QueryRow(ctx, Q_CLIENTE_INFOS, id).Scan(&limite, &saldo); err != nil {
		return nil, nil, err
	}

	var saldoAtualizado int64
	if tipo == "d" {
		saldoAtualizado = saldo - valor
	} else {
		saldoAtualizado = saldo + valor
	}

	if saldoAtualizado < limite*-1 {
		return nil, nil, errors.New("Saldo Insuficiente")
	}

	batch := &pgx.Batch{}
	batch.Queue(CD_STMT_UPDATE, saldoAtualizado, id)
	batch.Queue(T_INSERT_INFO, id, valor, tipo, descricao)

	br := tx.SendBatch(ctx, batch)
	if err = br.Close(); err != nil {
		return nil, nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, nil, err
	}

	return &limite, &saldoAtualizado, nil

}

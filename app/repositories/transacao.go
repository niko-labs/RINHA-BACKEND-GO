package repositories

import (
	"context"
	"errors"
	"log"
)

func (r *RepositorioBase) ExecutarTransacao(ctx context.Context, id int, valor int64, tipo string, descricao string) (*int64, *int64, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback(ctx)

	var limiteBanco, saldoBanco int64
	if err = tx.QueryRow(ctx, CLIENTE_INFO, id).Scan(&saldoBanco, &limiteBanco); err != nil {
		return nil, nil, err
	}

	// if err = tx.Commit(ctx); err != nil {
	// 	return nil, nil, err
	// }

	// tx, err = r.db.Begin(ctx)

	var UPDATE_SALDO string
	// 5 - 15 > 3 =
	if tipo == "d" && ((saldoBanco - valor) < -limiteBanco) {
		log.Println("Saldo insuficiente")
		return nil, nil, errors.New("Saldo insuficiente")
	}
	if tipo == "c" {
		UPDATE_SALDO = UPDATE_SALDO_C
	} else {
		UPDATE_SALDO = UPDATE_SALDO_D
	}

	if _, err := tx.Exec(ctx, T_INSERT_INFO, id, valor, tipo, descricao); err != nil {
		return nil, nil, err
	}
	if _, err = tx.Exec(ctx, UPDATE_SALDO, id, valor); err != nil {
		return nil, nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, nil, err
	}

	var saldoFinal, limiteFinal int64
	if err = r.db.QueryRow(ctx, CLIENTE_INFO, id).Scan(&saldoFinal, &limiteFinal); err != nil {
		return nil, nil, err
	}

	return &saldoFinal, &limiteFinal, nil

}

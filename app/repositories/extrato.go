package repositories

import (
	"context"
	"rinha-backend-2024-q1/types"
	"time"
)

func (r *RepositorioBase) ObterExtrato(ctx context.Context, id int) (*types.ExtratoOutput, error) {
	var limite, saldo int64

	if err := r.db.QueryRow(ctx, Q_CLIENTE_INFOS, id).Scan(&limite, &saldo); err != nil {
		return nil, err
	}

	var infoExtrato []types.InfoTransacao = []types.InfoTransacao{}
	if transacoes, err := r.db.Query(ctx, Q_EXTRATO, id); err != nil {
		return nil, err
	} else {
		for transacoes.Next() {
			var t types.InfoTransacao
			if err := transacoes.Scan(&t.Valor, &t.Tipo, &t.Descricao, &t.RealizadaEm); err != nil {
				return nil, err
			}
			infoExtrato = append(infoExtrato, t)
		}
	}

	extrato := &types.ExtratoOutput{
		Saldo: types.InfoSaldo{
			Total:       saldo,
			Limite:      limite,
			DataExtrato: time.Now(),
		},
		Transacoes: infoExtrato,
	}

	return extrato, nil
}

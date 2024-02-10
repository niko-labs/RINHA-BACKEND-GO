package types

import "time"

type TransacaoOutput struct {
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}

// Extrato

type ExtratoSaldo struct {
	Total       int64     `json:"total"`
	Limite      int64     `json:"limite"`
	DataExtrato time.Time `json:"data_extrato"`
}

type ExtratoOutput[T any] struct {
	Saldo      ExtratoSaldo `json:"saldo"`
	Transacoes []T          `json:"ultimas_transacoes"`
}

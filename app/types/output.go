package types

import "time"

type TransacaoOutput struct {
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}

// Extrato

type InfoSaldo struct {
	Total       int64     `json:"total"`
	Limite      int64     `json:"limite"`
	DataExtrato time.Time `json:"data_extrato"`
}

type InfoTransacao struct {
	Valor       int64     `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}

type ExtratoOutput struct {
	Saldo      InfoSaldo       `json:"saldo"`
	Transacoes []InfoTransacao `json:"ultimas_transacoes"`
}

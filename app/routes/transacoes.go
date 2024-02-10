package routes

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"rinha-backend-2024-q1/database"
	"rinha-backend-2024-q1/helpers"
	"rinha-backend-2024-q1/types"
)

const ROTA_TRANSACOES = "POST /clientes/{id}/transacoes"

func Transacoes(w http.ResponseWriter, r *http.Request) {

	id := helpers.PegaIdDoPathValue(r)

	body, _ := io.ReadAll(r.Body)
	dadosTransacao := &types.TransacaoInput{}
	json.Unmarshal(body, dadosTransacao)

	if dadosTransacao.Tipo != "c" && dadosTransacao.Tipo != "d" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if len(dadosTransacao.Descricao) == 0 || len(dadosTransacao.Descricao) > 10 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	db := database.PegarConexao()

	cliente, err := buscaInfoDoCliente(db, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if dadosTransacao.Tipo == types.DEBITO {
		saldo, err := debitar(db, id, cliente, dadosTransacao)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		response := types.TransacaoOutput{Limite: cliente.Limite, Saldo: *saldo}
		_json, _ := json.Marshal(response)
		w.Write(_json)
		return
	}
	if dadosTransacao.Tipo == types.CREDITO {
		saldo, err := creditar(db, id, cliente, dadosTransacao)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
		response := types.TransacaoOutput{Limite: cliente.Limite, Saldo: *saldo}
		_json, _ := json.Marshal(response)
		w.Write(_json)
		return
	}

}

func debitar(db *sql.DB, id int64, cliente *types.TbCliente, transacao *types.TransacaoInput) (saldo *int64, err error) {
	// Transforma o valor da transação em centavos
	valor := helpers.TransformarEmCentavos(transacao.Valor)

	// Atualiza o saldo do cliente
	cliente.Saldo -= valor

	// Verifica se o saldo do cliente é menor que o limite
	if cliente.Saldo < -cliente.Limite {
		return nil, errors.New("Saldo insuficiente")
	}

	// Inicia Transação
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	// Executa STMT_UPDATE e STMT_INSERT
	tx.Exec(database.CD_STMT_UPDATE, cliente.Saldo, id)
	tx.Exec(database.TD_STMT_INSERT, id, valor, transacao.Descricao)

	// Commita a transação
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Retorna o saldo do cliente
	return &cliente.Saldo, nil
}

func creditar(db *sql.DB, id int64, cliente *types.TbCliente, transacao *types.TransacaoInput) (saldo *int64, err error) {

	// Transforma o valor da transação em centavos
	valor := helpers.TransformarEmCentavos(transacao.Valor)

	// Atualiza o saldo do cliente
	cliente.Saldo += valor

	// Inicia Transação
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	// Executa STMT_UPDATE e STMT_INSERT
	tx.Exec(database.CD_STMT_UPDATE, cliente.Saldo, id)
	tx.Exec(database.TC_STMT_INSERT, id, valor, transacao.Descricao)

	// Commita a transação
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &cliente.Saldo, nil
}

func buscaInfoDoCliente(db *sql.DB, id int64) (*types.TbCliente, error) {

	rows, err := db.Query(database.Q_CLIENTE, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cliente types.TbCliente
	for rows.Next() {
		err = rows.Scan(&cliente.Id, &cliente.Limite, &cliente.Saldo)
	}
	if err != nil {
		return nil, err
	}
	// if cliente.Id == 0 {
	// 	return nil, errors.New("Cliente não encontrado")
	// }

	return &cliente, nil
}

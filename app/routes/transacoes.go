package routes

import (
	"database/sql"
	"errors"
	"io"
	"net/http"
	"rinha-backend-2024-q1/database"
	"rinha-backend-2024-q1/helpers"
	"rinha-backend-2024-q1/types"
)

const ROTA_TRANSACOES = "POST /clientes/{id}/transacoes"

var usuarios = map[int8]types.TbCliente{}

func Transacoes(w http.ResponseWriter, r *http.Request) {

	id := helpers.PegaIdDoPathValue(r)

	idValido := helpers.VerificaSeIdEstaEntreUmOuCinco(id)
	if !idValido {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, _ := io.ReadAll(r.Body)

	dadosTransacao := &types.TransacaoInput{}
	err := helpers.Json.Unmarshal(body, dadosTransacao)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if dadosTransacao.Tipo != "c" && dadosTransacao.Tipo != "d" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if len(dadosTransacao.Descricao) == 0 || len(dadosTransacao.Descricao) > 10 || dadosTransacao.Descricao == "" {
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
		_json, _ := helpers.Json.Marshal(response)
		w.Write(_json)
		return
	}
	if dadosTransacao.Tipo == types.CREDITO {
		saldo, err := creditar(db, id, cliente, dadosTransacao)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
		}
		response := types.TransacaoOutput{Limite: cliente.Limite, Saldo: *saldo}
		_json, _ := helpers.Json.Marshal(response)
		w.Write(_json)
		return
	}

}

// func verificaCorpoDaRequisicao(transacao *types.TransacaoInput) error {
// 	if transacao.Tipo != "c" && transacao.Tipo != "d" {
// 		return errors.New("Tipo de transação inválido")
// 	}

// 	if len(*transacao.Descricao) == 0 || len(*transacao.Descricao) > 10 || transacao.Descricao == nil || *transacao.Descricao == "" {
// 		return errors.New("Descrição inválida")
// 	}

// 	return nil
// }

func debitar(db *sql.DB, id int8, cliente *types.TbCliente, transacao *types.TransacaoInput) (saldo *int64, err error) {

	// Atualiza o saldo do cliente
	cliente.Saldo -= transacao.Valor

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
	tx.Exec(database.TD_STMT_INSERT, id, transacao.Valor, transacao.Descricao)

	// Commita a transação
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Retorna o saldo do cliente
	return &cliente.Saldo, nil
}

func creditar(db *sql.DB, id int8, cliente *types.TbCliente, transacao *types.TransacaoInput) (saldo *int64, err error) {

	// Atualiza o saldo do cliente
	cliente.Saldo += transacao.Valor

	// Inicia Transação
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	// Executa STMT_UPDATE e STMT_INSERT
	tx.Exec(database.CD_STMT_UPDATE, cliente.Saldo, id)
	tx.Exec(database.TC_STMT_INSERT, id, transacao.Valor, transacao.Descricao)

	// Commita a transação
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &cliente.Saldo, nil
}

func buscaInfoDoCliente(db *sql.DB, id int8) (*types.TbCliente, error) {

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

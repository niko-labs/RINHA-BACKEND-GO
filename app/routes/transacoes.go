package routes

import (
	"database/sql"
	"io"
	"net/http"
	"rinha-backend-2024-q1/database"
	"rinha-backend-2024-q1/helpers"
	"rinha-backend-2024-q1/types"
)

const ROTA_TRANSACOES = "POST /clientes/{id}/transacoes"

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

	if !dadosTransacao.Validar() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	db := database.PegarConexao()

	saldo, limite, err := executarTransacao(db, id, dadosTransacao)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	response := types.TransacaoOutput{Limite: *limite, Saldo: *saldo}
	_json, _ := helpers.Json.Marshal(response)
	w.Write(_json)
	return

}

func executarTransacao(db *sql.DB, id int8, info *types.TransacaoInput) (saldo, limite *int64, err error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, nil, err
	}

	err = tx.QueryRow(database.T_DC_CTE_UNICO, id, info.Valor, info.Tipo, info.Descricao).Scan(&saldo, &limite)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, nil, err
	}
	return saldo, limite, err
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

	return &cliente, nil
}

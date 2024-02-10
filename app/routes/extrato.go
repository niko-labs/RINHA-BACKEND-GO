package routes

import (
	"database/sql"
	"encoding/json"

	"net/http"
	"rinha-backend-2024-q1/database"
	"rinha-backend-2024-q1/helpers"
	"rinha-backend-2024-q1/types"
	"time"
)

const ROTA_EXTRATO = "GET /clientes/{id}/extrato"

func Extrato(w http.ResponseWriter, r *http.Request) {

	id := helpers.PegaIdDoPathValue(r)

	db := database.PegarConexao()

	cliente, err := vericarSeClienteExiste(db, id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	transacoes, err := buscarExtratoDoCliente(db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	exSaldo := types.ExtratoSaldo{Total: cliente.Saldo, Limite: cliente.Limite, DataExtrato: time.Now()}
	extrato := types.ExtratoOutput[types.TbTransacao]{Saldo: exSaldo, Transacoes: *transacoes}

	_json, _ := json.Marshal(extrato)
	w.Write(_json)

}

func buscarExtratoDoCliente(db *sql.DB, id int64) (*[]types.TbTransacao, error) {

	STMT := "SELECT valor, tipo, descricao, realizada_em FROM transacoes WHERE cliente_id = $1 ORDER BY realizada_em DESC LIMIT 10;"
	rows, err := db.Query(STMT, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transacoes := []types.TbTransacao{}

	for rows.Next() {
		transacao := types.TbTransacao{}
		rows.Scan(&transacao.Valor, &transacao.Tipo, &transacao.Descricao, &transacao.RealizadoEm)
		transacoes = append(transacoes, transacao)
	}

	return &transacoes, nil
}

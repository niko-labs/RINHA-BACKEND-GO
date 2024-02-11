package routes

import (
	"net/http"
	"rinha-backend-2024-q1/database"
	"rinha-backend-2024-q1/helpers"
)

const ROTA_EXTRATO = "GET /clientes/{id}/extrato"

func Extrato(w http.ResponseWriter, r *http.Request) {

	id := helpers.PegaIdDoPathValue(r)

	idValido := helpers.VerificaSeIdEstaEntreUmOuCinco(id)
	if !idValido {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	db := database.PegarConexao()

	var dbJson string
	rows, err := db.Query(database.Q_EXTRATO_CLIENTE, id)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&dbJson)
	}

	w.Write([]byte(dbJson))
	return

}

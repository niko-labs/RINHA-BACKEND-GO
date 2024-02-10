package helpers

import (
	"net/http"
	"strconv"
)

func PegaIdDoPathValue(r *http.Request) int8 {
	id := r.PathValue("id")
	idToInteger, _ := strconv.Atoi(id)
	return int8(idToInteger)
}

func TransformarEmCentavos(valor int64) int64 {
	return int64(valor * 100)
}

func VerificaSeIdEstaEntreUmOuCinco(id int8) bool {
	if id >= 1 && id <= 5 {
		return true
	}
	return false
}

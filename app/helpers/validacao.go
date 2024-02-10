package helpers

import (
	"net/http"
	"strconv"
)

func PegaIdDoPathValue(r *http.Request) int64 {
	id := r.PathValue("id")
	idToInteger, _ := strconv.Atoi(id)
	return int64(idToInteger)
}

func TransformarEmCentavos(valor int64) int64 {
	return int64(valor * 100)
}

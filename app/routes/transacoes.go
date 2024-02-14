package routes

import (
	"net/http"
	"rinha-backend-2024-q1/helpers"
	"rinha-backend-2024-q1/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

const ROTA_TRANSACOES = "/clientes/:id/transacoes"

func (r RotaBase) RealizarTransacao(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if idValido := helpers.VerificaSeIdMenorIgualCinco(id); !idValido {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	transacao := &types.TransacaoInput{}
	if err := c.ShouldBindJSON(transacao); err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	if !transacao.Validar() {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	limite, saldo, err := r.repo.ExecutarTransacao(c, id, transacao.Valor, transacao.Tipo, transacao.Descricao)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	c.JSON(http.StatusOK, types.TransacaoOutput{
		Saldo:  *saldo,
		Limite: *limite,
	})

}

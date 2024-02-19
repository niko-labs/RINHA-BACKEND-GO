package routes

import (
	"context"
	"net/http"
	"rinha-backend-2024-q1/helpers"
	"rinha-backend-2024-q1/types"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const ROTA_TRANSACOES = "/clientes/:id/transacoes"

func (r RotaBase) RealizarTransacao(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if helpers.VerificaSeIdMaiorQueCinco(id) {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	transacao := &types.TransacaoInput{}
	_ = c.ShouldBindJSON(transacao)

	if !transacao.Validar() {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	limite, saldo, err := r.repo.ExecutarTransacaoCreditoDebito(ctx, id, transacao.Valor, transacao.Tipo, transacao.Descricao)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"saldo":  *saldo,
		"limite": *limite,
	})
}

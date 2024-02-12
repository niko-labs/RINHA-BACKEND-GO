package routes

import (
	"context"
	"net/http"
	"rinha-backend-2024-q1/helpers"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const ROTA_EXTRATO = "/clientes/:id/extrato"

func (r RotaBase) ConsultarExtrato(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if idValido := helpers.VerificaSeIdMenorIgualCinco(id); !idValido {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	extrato, err := r.repo.ObterExtrato(ctx, id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	c.String(http.StatusOK, extrato)

}

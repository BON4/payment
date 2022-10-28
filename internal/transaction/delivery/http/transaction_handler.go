package http

import (
	"github.com/BON4/payment/internal/domain"
	"github.com/gin-gonic/gin"
)

type TransctionHandler struct {
	tUc domain.TxUsecase
}

func NewTransactionHandler(group *gin.RouterGroup, tUc domain.TxUsecase) *TransctionHandler {
	return &TransctionHandler{
		tUc: tUc,
	}
}

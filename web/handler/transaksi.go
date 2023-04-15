package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type transaksi_handler struct {
}

func Newtransaksi_handler() *transaksi_handler {
	return &transaksi_handler{}
}

func (h *transaksi_handler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "transaksi_index.html", nil)
}

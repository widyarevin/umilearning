package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type acaraHandler struct {
}

func NewacaraHandler() *acaraHandler {
	return &acaraHandler{}
}

func (h *acaraHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "acara_index.html", nil)
}

func (h *acaraHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "new_acara.html", nil)
}

func (h *acaraHandler) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "acara_edit.html", nil)
}

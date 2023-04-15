package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type dashboardHandler struct {
}

func NewDashboardHandler() *dashboardHandler {
	return &dashboardHandler{}
}

func (h *dashboardHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

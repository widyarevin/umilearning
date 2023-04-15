package handler

import (
	"net/http"
	"umiEvient/acara"
	"umiEvient/helper"

	"github.com/gin-gonic/gin"
)

type acaraHandler struct {
	acaraService acara.Service
}

func NewacaraHandler(acara acara.Service) *acaraHandler {
	return &acaraHandler{acara}
}

func (h *acaraHandler) GetAcaras(c *gin.Context) {
	acaras, err := h.acaraService.GetAllAcaras()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var acaraResponses []acara.AcaraResponse
	for _, k := range acaras {
		res := acara.FormatAcara(k)
		acaraResponses = append(acaraResponses, res)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": acaraResponses,
	})
}

func (h *acaraHandler) GetAcara(c *gin.Context) {
	var input acara.GetAcaraDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	acaraDetail, err := h.acaraService.GetAcaraByID(input)
	if err != nil {
		response := helper.APIRESPONSE("Failed to get detail of acara", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIRESPONSE("Acara detail Succes", http.StatusOK, "success", acara.FormatAcara(acaraDetail))
	c.JSON(http.StatusOK, response)
}

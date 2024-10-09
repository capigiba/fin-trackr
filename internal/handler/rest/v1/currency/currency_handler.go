package currency

import (
	"fintrack/internal/domain/model"
	"fintrack/internal/service/currency"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
	service currency.CurrencyService
}

func NewCurrencyHandler(srv currency.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{service: srv}
}

func (h *CurrencyHandler) Convert(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")

	req := model.ConvertRequest{
		Q: []string{from + "_" + to},
	}

	rate, err := h.service.Convert(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rate": rate.Results[from+"_"+to].Val})
}

func (h *CurrencyHandler) ConvertCompact(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")

	req := model.ConvertRequest{
		Q: []string{from + "_" + to},
	}

	rate, err := h.service.ConvertCompact(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rate": rate})
}

func (h *CurrencyHandler) ConvertHistorical(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	dateStr := c.Query("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}

	req := model.ConvertHistoricalRequest{
		Q:    []string{from + "_" + to},
		Date: date,
	}

	rate, err := h.service.ConvertHistorical(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rate": rate.Results[from+"_"+to].Val})
}

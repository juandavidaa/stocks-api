package stocks

import (
	"github.com/gin-gonic/gin"
	"github.com/juandavidaa/stocks-api/internal/dto"
	"github.com/juandavidaa/stocks-api/internal/usecases/stocks"
)

type Handlers struct {
	GetBestStocksUC stocks.GetBestStocks
}

func (h Handlers) getBestStocks(c *gin.Context) {
	var req dto.GetStocks

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	out, status, err := h.GetBestStocksUC.Execute(c, req)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, out)
}

package users

import (
	"github.com/gin-gonic/gin"
	"github.com/juandavidaa/stocks-api/internal/dto"
	"github.com/juandavidaa/stocks-api/internal/usecases/users"
)

type Handlers struct {
	CreateUC users.Create
	LoginUC  users.Login
}

func (h Handlers) create(c *gin.Context) {
	var req dto.CreateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error ": err.Error()})
		return
	}

	out, status, err := h.CreateUC.Execute(c, req)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, out)
}
func (h Handlers) login(c *gin.Context) {
	var req users.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	out, status, err := h.LoginUC.Execute(c, req)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, out)
}

package users

import (
	"github.com/gin-gonic/gin"
)

type Module struct {
	Handlers Handlers
}

func (m Module) Register(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("", m.Handlers.create)
		users.POST("/login", m.Handlers.login)
	}
}

package main

import (
	"database/sql"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juandavidaa/stocks-api/core"
	"github.com/juandavidaa/stocks-api/internal/infra/middleware"
	"github.com/juandavidaa/stocks-api/internal/infra/modules/stocks"
	"github.com/juandavidaa/stocks-api/internal/infra/modules/users"
)

func buildRouter(db *sql.DB) *gin.Engine {
	engine := gin.Default()

	//cors
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := engine.Group("api/v1")

	// users
	usersMod := users.New(db)
	usersMod.Register(v1)

	// Auth middleware
	authenticated := v1.Group("/")
	auth := middleware.Auth()
	authenticated.Use(auth)

	// stocks
	stocksMod := stocks.New(db)
	stocksMod.Register(authenticated)

	return engine
}

func main() {
	cfg := core.ConfigInstance()
	db := core.ConnectDB(cfg)
	defer db.Close()

	router := buildRouter(db)

	//router.GET("/", seed)
	router.Run(":" + cfg.AppPort)
}

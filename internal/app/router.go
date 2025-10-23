package app

import (
	"github.com/TheArchitectEngineer/mirros_api/internal/config"
	"github.com/TheArchitectEngineer/mirros_api/internal/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	h := handlers.NewHandler(db, cfg)

	// Health
	r.GET("/health", h.Health)

	// API group
	api := r.Group("/api")
	{
		api.GET("/users", h.ListUsers)
		api.POST("/users", h.CreateUser)
		// Additional routes will be added when translating the original endpoints
	}

	return r
}
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/controllers"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", controllers.SignUp)
			auth.POST("/login", controllers.Login)
			auth.POST("/logout", controllers.Logout)
			auth.GET("/me", middlewares.AuthMiddleware(), controllers.Me)
		}

		entries := v1.Group("/entries")
		entries.Use(middlewares.AuthMiddleware())
		{
			entries.GET("", controllers.GetEntries)
			entries.GET("/forecast", controllers.GetForecast)
			entries.POST("", controllers.CreateEntry)
			entries.POST("/bulk", controllers.CreateEntries)
			entries.PUT("/:id", controllers.UpdateEntry)
			entries.DELETE("/:id", controllers.DeleteEntry)
		}
	}
}

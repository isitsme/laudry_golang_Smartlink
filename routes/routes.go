package routes

import (
	"github.com/gin-gonic/gin"
	"laundry-api/controllers"
	"laundry-api/middleware"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Auth
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Public service routes
		api.GET("/services", controllers.ListServices)
		api.GET("/services/:id", controllers.GetService)
	}

	// protected
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		// service management (create/update/delete) - protected
		protected.POST("/services", controllers.CreateService)
		protected.PUT("/services/:id", controllers.UpdateService)
		protected.DELETE("/services/:id", controllers.DeleteService)

		// orders
		protected.POST("/orders", controllers.CreateOrder)
		protected.GET("/orders", controllers.ListOrders)
		protected.GET("/orders/:id", controllers.GetOrder)
		protected.PUT("/orders/:id", controllers.UpdateOrder)
		protected.DELETE("/orders/:id", controllers.DeleteOrder)
	}
}

package router

import (
	"galon-app/controller"
	"galon-app/controller/galon"
	"galon-app/controller/order"
	"galon-app/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	protected := r.Group("/admin")
	protected.Use(middlewares.JwtAuthMiddleware())

	protected.GET("/user", controller.CurrentUser)
	protected.POST("/user/galon", galon.Add)
	protected.PUT("/user/galon/:id", galon.Update)
	protected.DELETE("/user/galon/:id", galon.Delete)
	protected.GET("/user/galon", galon.GetAll)

	protected.POST("user/order", order.Make)
	protected.PUT("user/order/:id/processing", order.UpdateProcessing)
	protected.PUT("user/order/:id/on-delivery", order.UpdateDelivery)
	protected.PUT("user/order/:id/delivered", order.UpdateDelivered)
	protected.PUT("user/order/:id/completed", order.UpdateCompleted)

	r.Run(":" + os.Getenv("PORT"))
}

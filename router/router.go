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

	r.GET("/user", middlewares.JwtAuthMiddleware(), controller.CurrentUser)

	galonRoute := r.Group("/galon", middlewares.JwtAuthMiddleware())

	galonRoute.POST("/", galon.Add)
	galonRoute.PUT("/:id", galon.Update)
	galonRoute.DELETE("/:id", galon.Delete)
	galonRoute.GET("/", galon.GetAll)

	orderRoute := r.Group("/order", middlewares.JwtAuthMiddleware())

	orderRoute.POST("/", order.Make)
	orderRoute.PUT("/:id/processing", order.UpdateProcessing)
	orderRoute.PUT("/:id/on-delivery", order.UpdateDelivery)
	orderRoute.PUT("/:id/delivered", order.UpdateDelivered)
	orderRoute.PUT("/:id/completed", order.UpdateCompleted)
	orderRoute.DELETE("/:id/cancel", order.UpdateCanceled)
	orderRoute.GET("/", order.GetAll)
	orderRoute.GET("/user", order.GetByUserId)

	r.Run(":" + os.Getenv("PORT"))
}

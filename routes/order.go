package routes

import (
	"be/controllers"

	"github.com/gin-gonic/gin"
)

func Order(r *gin.Engine){
	r.GET("/orders", controllers.GetOrders)
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders/:id", controllers.DetailOrder)
	r.PUT("/orders/:id", controllers.UpdateOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)
}
package routes

import (
	"be/controllers"

	"github.com/gin-gonic/gin"
)

func Car(r *gin.Engine)  {
	r.GET("/cars", controllers.GetCars)
	r.POST("/cars", controllers.CreateCar)
	r.GET("/cars/:id", controllers.DetailCar)
	r.PUT("/cars/:id", controllers.UpdateCar)
	r.DELETE("/cars/:id", controllers.DeleteCar)
}
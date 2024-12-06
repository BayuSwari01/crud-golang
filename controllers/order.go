package controllers

import (
	"be/config"
	"be/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetOrders(ctx *gin.Context)  {
	db, _ := config.Connect()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Could not get DB: %v", err)
		}
		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	var orders []models.Order

	err := db.Find(&orders).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to get orders",
		})
		return
	}

	ctx.JSON(200, orders)
}

func CreateOrder(ctx *gin.Context)  {
	db, _ := config.Connect()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Could not get DB: %v", err)
		}
		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	var order models.Order
	ctx.BindJSON(&order)

	var car models.Car
	err := db.First(&car, order.CarId).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Car not found",
		})
		return
	}

	err = db.Create(&order).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to create order",
			"error": err,
		})
		return
	}

	ctx.JSON(200, order)
}

func DetailOrder(ctx *gin.Context)  {
	db, _ := config.Connect()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Could not get DB: %v", err)
		}
		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	var order models.Order
	var car models.Car
	id := ctx.Param("id")

	err := db.First(&order, id).Error

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to get order",
		})
		return
	}

	err = db.First(&car, order.CarId).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to get relation car",
		})
		return
	}

	//create new struct for response
	type Response struct {
		ID int `json:"id"`
		Car models.Car `json:"car"`
		Order models.Order `json:"order"`
	}

	var response Response
	response.ID = order.OrderId
	response.Car = car
	response.Order = order

	ctx.JSON(200, response)
}

func UpdateOrder(ctx *gin.Context)  {
	db, _ := config.Connect()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Could not get DB: %v", err)
		}
		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	id := ctx.Param("id")
	var order models.Order

	err := db.First(&order, id).Error
	
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to get order",
		})
		return
	}

	ctx.BindJSON(&order)

	err = db.Save(&order).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to update order",
		})
		return
	}

	ctx.JSON(200, order)
}

func DeleteOrder(ctx *gin.Context)  {
	db, _ := config.Connect()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Could not get DB: %v", err)
		}
		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	var order models.Order
	id := ctx.Param("id")

	err := db.Delete(&order, id).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to delete order",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Order deleted successfully",
	})
}


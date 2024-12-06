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

	err := db.Create(&order).Error
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
	id := ctx.Param("id")

	err := db.First(&order, id).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to get order",
		})
		return
	}

	ctx.JSON(200, order)
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

	var order models.Order
	ctx.BindJSON(&order)

	err := db.Save(&order).Error
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


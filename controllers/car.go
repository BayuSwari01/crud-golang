package controllers

import (
	"be/config"
	"be/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCars(ctx *gin.Context)  {
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

	limitParam := ctx.DefaultQuery("limit", "10")
	pageParam := ctx.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit
		
	var cars []models.Car

	err = db.Limit(limit).Offset(offset).Find(&cars).Error
    if err != nil {
        ctx.JSON(500, gin.H{
            "message": "Failed to get cars",
            "error":   err.Error(),
        })
        return
    }

	ctx.JSON(200, cars)
}

func CreateCar(ctx *gin.Context)  {
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

	var car models.Car
	ctx.BindJSON(&car)

	err := db.Create(&car).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to create car",
		})
		return
	}

	ctx.JSON(200, car)
}

func DetailCar(ctx *gin.Context)  {
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
	var car models.Car

	err := db.First(&car, id).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to get car",
		})
		return
	}

	ctx.JSON(200, car)
}

func UpdateCar(ctx *gin.Context)  {
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
	var car models.Car

	err := db.First(&car, id).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to get car",
		})
		return
	}

	ctx.BindJSON(&car)

	err = db.Save(&car).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to update car",
		})
		return
	}

	ctx.JSON(200, car)
}

func DeleteCar(ctx *gin.Context)  {
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
	var car models.Car

	err := db.First(&car, id).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to get car",
		})
		return
	}

	err = db.Delete(&car).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to delete car",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Car deleted successfully",
	})
}
package main

import (
	"be/migrations"
	"be/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main()  {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	migrations.RunMigrations()

	r := gin.Default()

	routes.Car(r)
	routes.Order(r)
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong. masuk",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
package migrations

import (
	"be/config"
	"be/models"
	"log"
)

func RunMigrations() {
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

	
	err := db.AutoMigrate(&models.Car{}, &models.Order{})

    if err != nil {
        log.Fatalf("Migration failed: %v", err)
    }

    log.Println("Migration completed successfully")
}

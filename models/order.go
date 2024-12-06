package models

import "time"

type Order struct {
	OrderId         int       `json:"order_id" gorm:"primaryKey;not null"`
	CarId           int       `json:"car_id" gorm:"not null"`
	OrderDate       time.Time `json:"order_date" gorm:"not null"`
	PickupDate      time.Time `json:"pickup_date" gorm:"not null"`
	DropoffDate     time.Time `json:"dropoff_date" gorm:"not null"`
	PickupLocation  string    `json:"pickup_location" gorm:"not null"`
	DropoffLocation string    `json:"dropoff_location" gorm:"not null"`
}

// type Car struct {
// 	ID    int    `json:"id" gorm:"primaryKey;not null"`
// 	Model string `json:"model" gorm:"not null"`
// 	Brand string `json:"brand" gorm:"not null"`
// }

// {
// 	"car_id": 1,
// 	"order_date": "2021-08-01T00:00:00Z",
// 	"pickup_date": "2021-08-02T00:00:00Z",
// 	"dropoff_date": "2021-08-03T00:00:00Z",
// 	"pickup_location": "Bandung",
// 	"dropoff_location": "Jakarta"
// }
// {
// 	"car_id": 2,
// 	"order_date": "2021-08-01T00:00:00Z",
// 	"pickup_date": "2021-08-02T00:00:00Z",
// 	"dropoff_date": "2021-08-03T00:00:00Z",
// 	"pickup_location": "Bandung",
// 	"dropoff_location": "Jakarta"
// }
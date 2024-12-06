package models

import "time"

type Order struct {
	OrderId    int `json:"order_id" gorm:"primaryKey;not null"`
	CarId	  int `json:"car_id" gorm:"not null"`
	OrderDate time.Time `json:"order_date" gorm:"not null"`
	PickupDate time.Time `json:"pickup_date" gorm:"not null"`
	DropoffDate time.Time `json:"dropoff_date" gorm:"not null"`
	PickupLocation string `json:"pickup_location" gorm:"not null"`
	DropoffLocation string `json:"dropoff_location" gor:"not null"`
}

// {
// 	"car_id": 1,
// 	"order_date": "2021-08-01",
// 	"pickup_date": "2021-08-02",
// 	"dropoff_date": "2021-08-03",
// 	"pickup_location": "Bandung",
// 	"dropoff_location": "Jakarta"
// }
// {
// 	"car_id": 2,
// 	"order_date": "2021-08-01",
// 	"pickup_date": "2021-08-02",
// 	"dropoff_date": "2021-08-03",
// 	"pickup_location": "Bandung",
// 	"dropoff_location": "Jakarta"
// }
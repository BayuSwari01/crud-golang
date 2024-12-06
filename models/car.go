package models

type Car struct {
	CarId	int     `json:"car_id" gorm:"primaryKey;not null"`
	CarName	string  `json:"car_name" gorm:"not null"`
	DayRate	float64 `json:"day_rate" gorm:"not null"`
	MonthRate	float64 `json:"month_rate" gorm:"not null"`
	ImageCar string  `json:"image_car" gorm:"not null"`
}
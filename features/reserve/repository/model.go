package repository

import (
	"be13/project/features/reserve"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID     uint
	Quantity   int `json:"quantity" form:"quantity"`
	HomestayID uint
	Title      string
	StartDate  string //`gorm:"unique"`
	EndDate    string //`gorm:"unique"`
}

type Homestay struct {
	gorm.Model
	Title         string
	Price         uint
	ReservationID []Reservation //`gorm:"foreignKey:homestay_id"`
}

type Results struct {
	ID         uint
	Quantity   int
	Title      string
	Images     string
	Price      int
	UserID     uint
	HomestayID uint
	StarDate   string
	EndDate    string
	Deleted_At string
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	// Homestay    []Homestay
	// Reservation []Reservation
}

func FromCore(dataCore reserve.Reservation) Reservation {
	resModel := Reservation{
		UserID:     dataCore.UserID,
		Quantity:   dataCore.Quantity,
		HomestayID: dataCore.HomestayID,
		StartDate:  dataCore.StartDate,
		EndDate:    dataCore.EndDate,
	}
	return resModel
}

func (res *Results) toCore() reserve.Reservation {
	return reserve.Reservation{
		ID:         res.ID,
		UserID:     res.UserID,
		HomestayID: res.HomestayID,
		StartDate:  res.StarDate,
		EndDate:    res.EndDate,
	}
}

func toCoreList(dataCart []Results) []reserve.Reservation {
	var dataCore []reserve.Reservation
	for key := range dataCart {
		if dataCart[key].Deleted_At == "" {
			dataCore = append(dataCore, dataCart[key].toCore())
		}
	}

	return dataCore

}

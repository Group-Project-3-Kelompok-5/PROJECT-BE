package delivery

import (
	"be13/project/features/reserve"
)

type Booking struct {
	UserID     uint   `json:"user_id" form:"user_id"`
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	Quantity   uint   `json:"quantity" form:"quantity"`
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
}

func toCore(data Booking) reserve.Reservation {
	return reserve.Reservation{
		UserID:     data.UserID,
		HomestayID: data.HomestayID,
		Quantity:   int(data.Quantity),
		StartDate:  data.StartDate,
		EndDate:    data.EndDate,
	}
}

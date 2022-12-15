package delivery

import (
	"be13/project/features/reserve"
)

type BookingResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Price     uint   `json:"price"`
}

type TotalRes struct {
	Duration   int `json:"duration"`
	TotalPrice int `json:"total_price"`
}

func fromCore(data reserve.Reservation) BookingResponse {
	return BookingResponse{
		ID:        data.ID,
		Title:     data.Title,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		Price:     data.Price,
	}
}

func fromCoreList(data []reserve.Reservation) ([]BookingResponse, TotalRes) {

	var dataRes []BookingResponse
	var totalRes TotalRes
	for i, v := range data {

		dataRes = append(dataRes, fromCore(v))
		totalRes.Duration += int(v.Price)
		totalRes.TotalPrice += data[i].Quantity * int(data[i].Price)

	}

	return dataRes, totalRes

}

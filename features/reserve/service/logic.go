package service

import (
	"be13/project/features/reserve"
	"errors"
)

type ReservService struct {
	ResData reserve.RepositoryEntities
}

func New(data reserve.RepositoryEntities) reserve.ServiceEntities {
	return &ReservService{
		ResData: data,
	}
}

func (rs *ReservService) AddReserv(data reserve.Reservation) (int, error) {
	if data.HomestayID == 0 {
		return -1, errors.New("data tidak boleh kosong")
	}
	row, err := rs.ResData.PostReserv(data)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (rs *ReservService) GetByToken(token int) ([]reserve.Reservation, error) {
	res, err := rs.ResData.SelectByToken(token)
	if err != nil {
		return nil, err
	}

	return res, nil
}

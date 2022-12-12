package service

import (
	"be13/project/features/homestay"

	"github.com/go-playground/validator/v10"
)

// bisnis Logic
type homeStayService struct {
	homeStayRepository homestay.RepositoryEntities //data repository dri entities
	validate           *validator.Validate
}

func NewHome(repo homestay.RepositoryEntities) homestay.ServiceEntities { //dengan kembalian user.service
	return &homeStayService{
		homeStayRepository: repo,
		validate:           validator.New(),
	}
}

// Create implements homestay.ServiceEntities
func (service *homeStayService) Create(input homestay.CoreHomestay) (err error) {
	err = service.homeStayRepository.Create(input)
	return err
}

// DeleteById implements homestay.ServiceEntities
func (*homeStayService) DeleteById(id int) (homestay.CoreHomestay, error) {
	panic("unimplemented")
}

// GetAllhomestay implements homestay.ServiceEntities
func (*homeStayService) GetAllhomestay() (data []homestay.CoreHomestay, err error) {
	panic("unimplemented")
}

// GetById implements homestay.ServiceEntities
func (*homeStayService) GetById(id int) (data homestay.CoreHomestay, err error) {
	panic("unimplemented")
}

// GetBytime implements homestay.ServiceEntities
func (service *homeStayService) GetBytime(start string, end string) (data []homestay.CoreHomestay, err error) {
	data, err = service.homeStayRepository.GetBytime(start, end)
	return data, err
}

// Update implements homestay.ServiceEntities
func (*homeStayService) Update(id int, input homestay.CoreHomestay) error {
	panic("unimplemented")
}

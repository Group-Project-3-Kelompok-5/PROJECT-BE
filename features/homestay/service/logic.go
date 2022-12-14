package service

import (
	"be13/project/features/homestay"
	"errors"

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
func (service *homeStayService) DeleteById(id int, userId int) (homestay.CoreHomestay, error) {
	data, err := service.homeStayRepository.DeleteById(id, userId) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return data, err
}

// GetAllhomestay implements homestay.ServiceEntities
func (service *homeStayService) GetAllhomestay() (data []homestay.CoreHomestay, err error) {
	data, err = service.homeStayRepository.GetAll() // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}

// GetById implements homestay.ServiceEntities
func (service *homeStayService) GetById(id int) (data homestay.CoreHomestay, err error) {
	data, err = service.homeStayRepository.GetById(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}

// GetBytime implements homestay.ServiceEntities

// Update implements homestay.ServiceEntities
func (service *homeStayService) Update(id int, userId int, input homestay.CoreHomestay) error {
	errUpdate := service.homeStayRepository.Update(id, userId, input)
	if errUpdate != nil {
		return errors.New("GAGAL mengupdate data , QUERY ERROR")
	}

	return nil
}

// GethHomestaybyidUser implements homestay.ServiceEntities
func (service *homeStayService) GethHomestaybyidUser(user_id int) (data []homestay.CoreHomestay, err error) {
	data, err = service.homeStayRepository.GethHomestaybyidUser(user_id)
	return data, err
}

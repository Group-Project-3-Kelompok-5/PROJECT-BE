package repository

import (
	"be13/project/features/reserve"

	"gorm.io/gorm"
)

type resData struct {
	db *gorm.DB
}

func New(db *gorm.DB) reserve.RepositoryEntities {
	return &resData{
		db: db,
	}
}

func (repo *resData) PostReserv(data reserve.Reservation) (int, error) {

	var dbCek Reservation
	repo.db.First(&dbCek, "homestay_id = ? AND user_id = ? ", data.HomestayID, data.UserID)

	if dbCek.ID != 0 {
		return 2, nil
	}

	var dbCekHomestay Homestay
	repo.db.First(&dbCekHomestay, "id = ? AND user_id = ? ", data.HomestayID, data.UserID)

	if dbCekHomestay.ID > 0 {
		return 3, nil
	}

	dataModel := FromCore(data)

	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil

}
func (repo *resData) SelectByToken(token int) ([]reserve.Reservation, error) {
	var dataResCek []Results
	tx := repo.db.Model(&Homestay{}).Select("reservations.id, reservations.start_date, reservations.title, reservations.end_date, reservations.quantity").Joins("inner join reservations on reservations.homestay_id = homestays.id").Where("reservations.user_id = ?", token).Scan(&dataResCek)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return toCoreList(dataResCek), nil
}

package reserve

type Reservation struct {
	ID         uint
	UserID     uint
	HomestayID uint
	Quantity   int
	Title      string
	StartDate  string
	EndDate    string
	Price      uint
}

type Homestay struct {
	ID    uint
	Title string
	Price uint
}

type ServiceEntities interface {
	AddReserv(data Reservation) (int, error)
	GetByToken(token int) ([]Reservation, error)
}

type RepositoryEntities interface {
	PostReserv(data Reservation) (int, error)
	SelectByToken(token int) ([]Reservation, error)
}

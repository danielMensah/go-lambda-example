package repository

type Repository struct {
	CustomerRepo Customer
	OrderRepo Order
}

func NewRepository() Repository {
	return Repository{}
}

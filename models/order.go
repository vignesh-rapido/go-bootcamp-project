package models

type Order struct {
	orderID string
	status  string
	amount  float64
	user    User
	product Product
	rating  Reviews
}

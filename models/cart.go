package models

type Cart struct {
	items  []Product
	total  float64
	user   string
	id     string
	date   string
	status string
}

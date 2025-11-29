package models

type Product struct {
	Name        string
	Price       float64
	description string
	quantity    int
	reviews     Reviews
}

type Reviews struct {
	rating    float64
	count     int
	avgRating float64
	feedback  []string
}

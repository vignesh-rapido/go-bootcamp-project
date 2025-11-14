package main

import (
	"fmt"
	"github.com/fatih/color"
	"practise/rating"
)

func main() {

	color.Green("Welcome to the rating app")
	rating := &rating.Rating{
		ProductId: "Apple Iphone",
	}
	addRatings(rating, "Vignesh", 4.5, "Great product")
	addRatings(rating, "Arjun", 3.5, "Worst product")
	addRatings(rating, "Sahil", 3, "Worst product")
	fmt.Println(rating)
}

func addRatings(r *rating.Rating, userName string, rating float64, comment string) {

	e := r.AddRatings(userName, rating, comment)
	if e != nil {
		fmt.Println("Error adding rating:", e)
	} else {
		fmt.Printf("Rating added successfully for user:%s\n", userName)
	}
}

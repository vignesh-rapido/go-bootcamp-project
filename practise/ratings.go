package main

import (
	"fmt"
	"github.com/fatih/color"
)

type Ratings struct {
	name     string
	ratings  int
	comments string
}

func getRatings() []Ratings {
	customerRatings := []Ratings{}
	customerRatings = append(customerRatings, Ratings{name: "Vignesh", ratings: 4, comments: "Very Professional agent"})
	customerRatings = append(customerRatings, Ratings{name: "Ram", ratings: 2, comments: "Agent was rude and unprofessional"})
	customerRatings = append(customerRatings, Ratings{name: "Raju", ratings: 1, comments: "Horrific behiaviour with the agent"})
	customerRatings = append(customerRatings, Ratings{name: "Reetha", ratings: 4, comments: "Very freidnly agent"})
	customerRatings = append(customerRatings, Ratings{name: "Reenu", ratings: 3, comments: "Good agent"})
	return customerRatings

}

func computeRatings() {
	allRatings := getRatings()
	for _, r := range allRatings {
		if r.ratings > 3 {
			fmt.Printf("Customer %s has rated %d and says %s\n", r.name, r.ratings, r.comments)
			color.Green("Thank-you for your feedback 😊👍")
		} else {
			fmt.Printf("Customer %s has rated %d and says %s\n", r.name, r.ratings, r.comments)
			color.Red("We are really sorry for your experience 😞")
		}
		// Print stars dynamically
		stars := r.ratings
		for i := 0; i < stars; i++ {
			fmt.Print("*")
		}
		fmt.Println() // extra newline for readability
	}

}

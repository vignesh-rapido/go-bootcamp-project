package rating

import (
	"fmt"
	"time"
)

type UserRating struct {
	userName string
	rating   float64
	comments Comment
}

type Rating struct {
	ProductId     string
	averageRating float64
	Ratings       []UserRating
	totalRating   float64
}

func (rating Rating) String() string {
	return fmt.Sprintf("Rating{productId: %s, averageRating: %.1f, ratings: %v}", rating.ProductId, rating.averageRating, rating.Ratings)
}

func (userRating UserRating) String() string {
	return fmt.Sprintf("UserRating{userName: %s, rating: %.1f, comment: %s}", userRating.userName, userRating.rating, userRating.comments)
}

func (r *Rating) AddRatings(userName string, rating float64, comment string) error {

	if rating < 0 || rating > 5 {
		return fmt.Errorf("Not a valid rating : %v", rating)
	}
	r.Ratings = append(r.Ratings, UserRating{userName: userName, rating: rating, comments: Comment{comment: comment, date: time.Now()}})
	r.totalRating += rating
	r.averageRating = r.totalRating / float64(len(r.Ratings))
	return nil
}

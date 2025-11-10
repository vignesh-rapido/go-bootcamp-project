package main

//
//import (
//	"fmt"
//	"time"
//)
//
//type UserRating struct {
//	uid      int
//	rating   float64
//	comments Comment
//}
//
//type Comment struct {
//	comment string
//	date    time.Time
//}
//
//type Rating struct {
//	id        int
//	avgRating float64
//	rating   []UserRating
//}
//
//func (r Rating) Add(uid int, rating float64, comment string) {
//	r.rating = append(r.rating, UserRating{uid: uid, rating: rating, comments: Comment{comment: comment, date: time.Now()}})
//	r.avgRating = (r.avgRating + rating) / float64(len(r.rating))
//}
//func (r Rating) String() string {
//	return fmt.Sprintf("Vignesh{id: %d, avgRating: %f, rating: %v}", r.id, r.avgRating, r.rating)
//}
//
//func (receiver UserRating) String() string {
//	return fmt.Sprintf("UserRating{uid: %d, rating: %f, comments: %v}", receiver.uid, receiver.rating, receiver.comments)
//}
//
//func (receiver Comment) String() string {
//	return fmt.Sprintf("Comment{comment: %s, date: %s}", receiver.comment, receiver.date)
//}
//
//func main() {
//
//	rating := []Rating{}
//	rating.Add(2, 4.5, "Great product")
//	rating.Add(3, 4.5, "Great product")
//	fmt.Println(rating)
//}

package rating

import (
	"fmt"
	"time"
)

type Comment struct {
	comment string
	date    time.Time
}

func (receiver Comment) String() string {
	return fmt.Sprintf("Comment{comment: %s, date: %s}", receiver.comment, receiver.date)
}

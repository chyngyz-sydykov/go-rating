package rating

import "time"

type BookMessage struct {
	BookId   int       `json:"book_id"`
	EditedAt time.Time `json:"EditedAt"`
	Event    string    `json:"event"`
}

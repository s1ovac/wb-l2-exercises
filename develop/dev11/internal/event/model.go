package event

import "time"

type Event struct {
	UserID      int       `json:"user_id"`
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

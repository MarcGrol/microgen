package news

import (
	"time"
)

type News struct {
	Year     int
	NewItems []NewsItem
}

type NewsItem struct {
	Message   string
	Sender    string
	Timestamp time.Time
}

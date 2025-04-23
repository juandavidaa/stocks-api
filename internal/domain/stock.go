package domain

import "time"

type Stock struct {
	ID         int
	Ticker     string
	Company    string
	Brokerage  string
	Action     string
	RatingFrom string
	RatingTo   string
	TargetFrom float64
	TargetTo   float64
	EventTime  time.Time
}

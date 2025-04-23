package dto

type CreateStock struct {
	Ticker     string  `json:"ticker"`
	Company    string  `json:"company"`
	Brokerage  string  `json:"brokerage"`
	Action     string  `json:"action"`
	RatingFrom string  `json:"rating_from"`
	RatingTo   string  `json:"rating_to"`
	TargetFrom string  `json:"target_from"`
	TargetTo   string  `json:"target_to"`
	Time       string  `json:"time"`
	LastPrice  float32 `json:"last_price"`

	UpsidePct     float32
	RatingDelta   int
	RecencyWeight float32
	ScoreBase     float32
}

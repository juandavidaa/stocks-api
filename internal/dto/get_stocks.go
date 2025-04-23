package dto

type GetStocks struct {
	Risk  string `json:"risk" form:"risk" binding:"required,oneof=low medium high"`
	Page  int    `json:"page" form:"page" default:"1"`
	Query string `json:"query" form:"query"`
}

package models

type Layout struct {
	Id    int64  `json:"id" validate:"gte=1"`
	Name  string `json:"name" validate:"len=3"`
	Ranks int    `json:"ranks" validate:"gte=1"`
	Rows  int    `json:"rows" validate:"gte=1"`
	// Label  string  `json:"label" validate:"len=3"`
	// Value  float64 `json:"value" validate:"gte=0.0001"`
	// UserId int64   `json:"userId" validate:"gte=1"`
}

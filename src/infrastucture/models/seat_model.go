package models

type Seat struct {
	Id        int64 `json:"id" validate:"gte=1"`
	RankId    int   `json:"ranks" validate:"gte=1"`
	Row       int   `json:"rows" validate:"gte=1"`
	AisleSeat bool  `json:"aisleSeat"`

	// Label  string  `json:"label" validate:"len=3"`
	// Value  float64 `json:"value" validate:"gte=0.0001"`
	// UserId int64   `json:"userId" validate:"gte=1"`
}

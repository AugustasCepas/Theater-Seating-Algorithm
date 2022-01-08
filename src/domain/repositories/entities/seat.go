package entities

type Seat struct {
	Id            int `json:"id" validate:"gte=1"`
	Index         int `json:"index" validate:"gte=1"`
	Number        int `json:"number" validate:"gte=1"`
	RankId        int `json:"rankId" validate:"gte=1"`
	RowId         int `json:"rowId" validate:"gte=1"`
	ReservationId int `json:"reservationId" validate:"gte=0"`
	// AisleSeat bool  `json:"aisleSeat"`
}

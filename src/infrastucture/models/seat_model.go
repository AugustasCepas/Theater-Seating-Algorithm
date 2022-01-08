package models

type Seat struct {
	Id            int  `json:"id" validate:"gte=1"`
	Index         int  `json:"index" validate:"gte=1"`  //[1...6]
	Number        int  `json:"number" validate:"gte=1"` // for example [1,3,5,6,4]
	RankId        int  `json:"rankId" validate:"gte=1"`
	RowId         int  `json:"rowId" validate:"gte=1"`
	ReservationId int  `json:"reservationId" validate:"gte=0"`
	AisleSeat     bool `json:"aisleSeat"`
}

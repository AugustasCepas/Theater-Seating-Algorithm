package models

type Rank struct {
	Id    int64 `json:"id" validate:"gte=1"`
	Price int64 `json:"price" validate:"gte=3"` //(7900 = $79.00) https://forum.golangbridge.org/t/what-is-the-proper-golang-equivalent-to-decimal-when-dealing-with-money/413/9
	Row   int64 `json:"row" validate:"gte=1"`
}

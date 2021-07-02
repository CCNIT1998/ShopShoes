package model

type Product struct {
	ID             int
	Name           string
	Description    string
	Image          string
	PriceOld       float64
	PriceCurrent   float64
	Sale           int
	Madein         string
	CategoryID     uint
	Category       *Category
	HistoricalSold uint
	RatingStar     uint
	Quantity       uint
}

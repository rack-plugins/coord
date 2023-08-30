package coord

type Coordinates struct {
	From      string  `json:"from" form:"from" validate:"required"`
	To        string  `json:"to" form:"to" validate:"required"`
	Latitude  float64 `json:"latitude" form:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" form:"longitude" validate:"required"`
	H         bool    `json:"h" form:"h"`
}

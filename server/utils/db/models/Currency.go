package models

// Currency model for the database
type Currency struct {
	ID   int64   `json:"id"`
	Name string  `json:"name"`
	Buy  float64 `json:"buy"`
	Sell float64 `json:"sell"`
}

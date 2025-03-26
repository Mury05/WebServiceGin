package models

type Album struct {
	ID     uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

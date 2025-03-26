package models

type Album struct {
	ID     string  `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Mingus Ah Um", Artist: "Charles Mingus", Price: 23.99},
	{ID: "5", Title: "Kind of Blue", Artist: "Miles Davis", Price: 19.99},
	{ID: "6", Title: "A Love Supreme", Artist: "John Coltrane", Price: 29.99},
}

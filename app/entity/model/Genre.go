package model

type Genre struct {
	GenreId uint   `gorm:"primaryKey" json:"genre_id"`
	Name    string `gorm:"not null" json:"name"`
	Slug    string `gorm:"not null" json:"slug"`
	Status  int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}

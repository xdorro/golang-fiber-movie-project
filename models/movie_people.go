package models

type MoviePeople struct {
	MoviePeopleId uint `gorm:"primarykey"`
	MovieId       uint `gorm:"not null"`
	ProfileId     uint `gorm:"not null"`
	Status        int8 `gorm:"default:1"`
	BaseModel
}

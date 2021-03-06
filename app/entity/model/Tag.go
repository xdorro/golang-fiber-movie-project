package model

type Tag struct {
	TagId  int64  `gorm:"primaryKey" json:"tag_id"`
	Name   string `gorm:"not null" json:"name"`
	Slug   string `gorm:"index:,not null" json:"slug"`
	Status int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}

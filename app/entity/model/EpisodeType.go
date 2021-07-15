package model

type EpisodeType struct {
	EpisodeTypeId uint   `gorm:"primaryKey" json:"episode_type_id"`
	Name          string `gorm:"not null;unique" json:"name"`
	Status        int    `gorm:"index:,default:1" json:"status"`
	BaseModel
}

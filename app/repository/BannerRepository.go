package repository

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"gorm.io/gorm"
	"log"
	"sync"
)

type BannerRepository struct {
	db *gorm.DB
}

func NewBannerRepository() *BannerRepository {
	if bannerRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if bannerRepository == nil {
				bannerRepository = &BannerRepository{
					db: db,
				}
				log.Println("Create new BannerRepository")
			}
		})
	}

	return bannerRepository
}

// FindAllBannersByStatus : Find banner by BannerId and Status
func (obj *BannerRepository) FindAllBannersByStatus(status int) (*[]model.Banner, error) {
	banners := make([]model.Banner, 0)

	err := db.Model(model.Banner{}).
		Find(&banners, "status = ?", status).Error

	return &banners, err
}

func (obj *BannerRepository) FindAllBannersByStatusNot(status int) (*[]model.Banner, error) {
	banners := make([]model.Banner, 0)

	err := db.Model(model.Banner{}).
		Find(&banners, "status <> ?", status).Error

	return &banners, err
}

// FindBannerByIdAndStatus : Find banner by BannerId and Status
func (obj *BannerRepository) FindBannerByIdAndStatus(id string, status int) (*model.Banner, error) {
	var banner model.Banner

	err := obj.db.Model(model.Banner{}).
		Where("banner_id = ? AND status = ?", id, status).
		Find(&banner).Error

	return &banner, err
}

func (obj *BannerRepository) FindBannerByIdAndStatusNot(id string, status int) (*model.Banner, error) {
	var banner model.Banner

	err := obj.db.Model(model.Banner{}).
		Where("banner_id = ? AND status <> ?", id, status).
		Find(&banner).Error

	return &banner, err
}

func (obj *BannerRepository) SaveBanner(banner model.Banner) (*model.Banner, error) {
	err := obj.db.Model(model.Banner{}).
		Save(&banner).Error

	return &banner, err
}
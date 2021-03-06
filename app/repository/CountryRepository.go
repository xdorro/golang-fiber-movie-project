package repository

import (
	"log"
	"sync"

	"gorm.io/gorm"

	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
	"github.com/xdorro/golang-fiber-movie-project/pkg/util"
)

type CountryRepository struct {
	db *gorm.DB
}

func NewCountryRepository() *CountryRepository {
	if countryRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if countryRepository == nil {
				countryRepository = &CountryRepository{
					db: db,
				}
				log.Println("Create new CountryRepository")
			}
		})
	}

	return countryRepository
}

func (obj *CountryRepository) CountAllCountriesStatusNotIn(status []int) (int64, error) {
	var count int64

	err := db.
		Model(&model.Country{}).
		Select("countries.country_id").
		Where("countries.status NOT IN ?", status).
		Count(&count).Error

	return count, err
}

func (obj *CountryRepository) FindAllCountriesByStatusNot(status int) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	err := db.Model(model.Country{}).
		Find(&countries, "status <> ?", status).Error

	return &countries, err
}

func (obj *CountryRepository) FindAllCountriesByStatusNotIn(status []int) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	err := db.Model(model.Country{}).
		Find(&countries, "status NOT IN ?", status).Error

	return &countries, err
}

func (obj *CountryRepository) FindCountryByIdAndStatusNot(id string, status int) (*model.Country, error) {
	uid := util.ParseStringToInt64(id)

	var country model.Country
	err := db.Model(model.Country{}).
		Where("country_id = ? AND status <> ?", uid, status).
		Find(&country).Error

	return &country, err
}

func (obj *CountryRepository) FindAllCountriesByCountryIdsInAndStatusNotIn(
	countryIds []int64, status []int,
) (*[]model.Country, error) {
	countries := make([]model.Country, 0)

	err := db.Model(model.Country{}).
		Find(&countries, "country_id IN ? AND status NOT IN ?", countryIds, status).Error

	return &countries, err
}

func (obj *CountryRepository) SaveCountry(country model.Country) (*model.Country, error) {
	err := db.Model(model.Country{}).
		Save(&country).Error

	return &country, err
}

func (obj *CountryRepository) UpdateCountry(countryId string, country model.Country) (*model.Country, error) {
	err := db.Model(model.Country{}).
		Where("country_id = ?", countryId).
		Save(&country).Error

	return &country, err
}

func (obj *CountryRepository) FindCountryBySlugAndCountryIdNotAndStatusNotIn(
	slug string, id string, status []int,
) (*model.Country, error) {
	var country model.Country

	err := obj.db.
		Where("country_id <> ?", id).
		Where("slug = ? AND status NOT IN ?", slug, status).
		Find(&country).Error

	return &country, err
}

func (obj *CountryRepository) FindCountryBySlugAndStatusNotIn(slug string, status []int) (*model.Country, error) {
	var country model.Country

	err := obj.db.
		Where("slug = ? AND status NOT IN ?", slug, status).
		Find(&country).Error

	return &country, err
}

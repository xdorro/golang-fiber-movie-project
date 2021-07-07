package repository

import (
	"errors"
	"github.com/xdorro/golang-fiber-base-project/app/entity/dto"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"gorm.io/gorm"
)

// FindAllMoviesByStatus : Find movie by MovieId and Status
func FindAllMoviesByStatus(status int) (*[]model.Movie, error) {
	movies := make([]model.Movie, 0)

	if err := db.Find(&movies, "status = ?", status).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movies, nil
}

func FindAllMoviesByStatusNot(status int) (*[]dto.SearchMovieDTO, error) {
	movies := make([]dto.SearchMovieDTO, 0)

	if err := db.
		Model(&model.Movie{}).
		Select("movies.movie_id, movies.name, movies.slug, movies.status, movies.movie_type_id, movie_types.name movie_type_name").
		Joins("LEFT JOIN movie_types on movies.movie_type_id = movie_types.movie_type_id").
		Where("movies.status <> ?", status).
		Find(&movies).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movies, nil
}

// FindMovieByIdAndStatus : Find movie by MovieId and Status
func FindMovieByIdAndStatus(id string, status int) (*model.Movie, error) {
	uid := util.ParseStringToUInt(id)

	var movie model.Movie
	if err := db.Where("movie_id = ? AND status = ?", uid, status).Find(&movie).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movie, nil
}

func FindMovieByIdAndStatusNot(id string, status int) (*model.Movie, error) {
	uid := util.ParseStringToUInt(id)

	var movie model.Movie
	if err := db.Where("movie_id = ? AND status <> ?", uid, status).Find(&movie).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &movie, nil
}

// SaveMovie : Save Movie
func SaveMovie(movie model.Movie) (*model.Movie, error) {
	if err := db.Save(&movie).Error; err != nil {
		return nil, err
	}

	return &movie, nil
}

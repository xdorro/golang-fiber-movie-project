package repository

import (
	"log"
	"sync"

	"gorm.io/gorm"

	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
)

type MovieGenreRepository struct {
	db *gorm.DB
}

func NewMovieGenreRepository() *MovieGenreRepository {
	if movieGenreRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if movieGenreRepository == nil {
				movieGenreRepository = &MovieGenreRepository{
					db: db,
				}

				log.Println("Create new MovieGenreRepository")
			}
		})
	}

	return movieGenreRepository
}

// CreateMovieGenreByMovieId : Create MovieGenre By MovieId
func (obj *MovieGenreRepository) CreateMovieGenreByMovieId(movieGenres []model.MovieGenre) error {
	err := db.
		Model(&model.MovieGenre{}).
		Create(&movieGenres).Error

	return err
}

func (obj *MovieGenreRepository) RemoveMovieGenreByMovieIdAndGenreIds(movieId int64, genreIds []int64) error {
	err := db.
		Model(&model.MovieGenre{}).
		Where("movie_id = ? AND genre_id IN ?", movieId, genreIds).
		Delete(&model.MovieGenre{}).Error

	return err
}

func (obj *MovieGenreRepository) FindAllGenresByMovieIdAndStatusNotIn(movieId int64, status []int) (
	*[]model.Genre, error,
) {
	genres := make([]model.Genre, 0)

	err := db.
		Model(&model.Genre{}).
		Select("genres.*").
		Joins("LEFT JOIN movie_genres ON genres.genre_id = movie_genres.genre_id").
		Where("movie_genres.movie_id = ? AND genres.status NOT IN ?", movieId, status).
		Find(&genres).Error

	return &genres, err
}

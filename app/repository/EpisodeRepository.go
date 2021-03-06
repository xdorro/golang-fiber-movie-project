package repository

import (
	"log"
	"sync"

	"gorm.io/gorm"

	"github.com/xdorro/golang-fiber-movie-project/app/entity/model"
)

type EpisodeRepository struct {
	db *gorm.DB
}

func NewEpisodeRepository() *EpisodeRepository {
	if episodeRepository == nil {
		once = &sync.Once{}

		once.Do(func() {
			if episodeRepository == nil {
				episodeRepository = &EpisodeRepository{
					db: db,
				}
				log.Println("Create new EpisodeRepository")
			}
		})
	}

	return episodeRepository
}

func (obj *EpisodeRepository) FindAllEpisodesByMovieIdAndStatusNot(movieId string, status int) (
	*[]model.Episode, error,
) {
	episodes := make([]model.Episode, 0)

	err := db.Model(model.Episode{}).
		Find(&episodes, "movie_id = ? AND status <> ?", movieId, status).Error

	return &episodes, err
}

func (obj *EpisodeRepository) FindAllEpisodesByMovieIdAndStatusNotIn(movieId string, status []int) (
	*[]model.Episode, error,
) {
	episodes := make([]model.Episode, 0)

	err := db.Model(model.Episode{}).
		Select("episodes.*").
		Joins("JOIN movies ON episodes.movie_id = movies.movie_id").
		Where("movies.status NOT IN ?", status).
		Where("episodes.status NOT IN ?", status).
		Where("movies.movie_id = ?", movieId).
		Find(&episodes).Error

	return &episodes, err
}

func (obj *EpisodeRepository) FindEpisodeByIdAndStatusNot(episodeId string, status int) (*model.Episode, error) {
	episode := new(model.Episode)

	err := db.Model(model.Episode{}).
		Find(&episode, "episode_id = ? AND status <> ?", episodeId, status).Error

	return episode, err
}

// SaveEpisode : Save Episode
func (obj *EpisodeRepository) SaveEpisode(episode model.Episode) (*model.Episode, error) {
	err := db.Model(model.Episode{}).
		Save(&episode).Error

	return &episode, err
}

func (obj *EpisodeRepository) UpdateEpisode(episodeId string, episode model.Episode) (*model.Episode, error) {
	err := db.Model(model.Episode{}).
		Where("episode_id = ?", episodeId).
		Save(&episode).Error

	return &episode, err
}

func (obj *EpisodeRepository) UpdateStatusByEpisodeId(episodeId string, status int) error {
	err := obj.db.Model(&model.Episode{}).
		Where("episode_id = ?", episodeId).
		Update("status", status).Error

	return err
}

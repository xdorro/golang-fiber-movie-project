package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/entity/model"
	"github.com/xdorro/golang-fiber-base-project/app/entity/request"
	"github.com/xdorro/golang-fiber-base-project/app/entity/response"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"github.com/xdorro/golang-fiber-base-project/pkg/validator"
	"log"
	"sync"
)

type EpisodeController struct {
	episodeRepository       *repository.EpisodeRepository
	episodeDetailRepository *repository.EpisodeDetailRepository
}

func NewEpisodeController() *EpisodeController {
	if episodeController == nil {
		once = &sync.Once{}

		once.Do(func() {
			episodeController = &EpisodeController{
				episodeRepository:       repository.NewEpisodeRepository(),
				episodeDetailRepository: repository.NewEpisodeDetailRepository(),
			}
			log.Println("Create new EpisodeController")
		})
	}

	return episodeController
}

func (obj *EpisodeController) FindAllEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("id")

	if _, err := validator.ValidateMovieId(movieId); err != nil {
		return err
	}

	episodes, err := obj.episodeRepository.FindAllEpisodesByMovieIdAndStatusNot(movieId, util.StatusDeleted)

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", episodes)
}

func (obj *EpisodeController) FindEpisodeByEpisodeId(c *fiber.Ctx) error {
	episodeId := c.Params("episodeId")

	episode, err := validator.ValidateEpisodeId(episodeId)

	if err != nil {
		return err
	}

	episodeDetails, err := obj.episodeDetailRepository.FindEpisodeDetailsByIdAndStatusNot(episodeId, []int{util.StatusDeleted})

	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := &response.MovieEpisodeDetailResponse{
		Episode: model.Episode{
			EpisodeId: episode.EpisodeId,
			Name:      episode.Name,
			MovieId:   episode.MovieId,
			Status:    episode.Status,
		},
		EpisodeDetails: *episodeDetails,
	}

	return util.ResponseSuccess("Thành công", result)
}

func (obj *EpisodeController) CreateEpisodesByMovieId(c *fiber.Ctx) error {
	movieId := c.Params("id")

	movie, err := validator.ValidateMovieId(movieId)
	if err != nil {
		return err
	}

	episodeRequest := new(request.EpisodeRequest)
	if err = c.BodyParser(episodeRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	newEpisode := model.Episode{
		MovieId: movie.MovieId,
		Name:    episodeRequest.Name,
		Status:  episodeRequest.Status,
	}

	episode, err := obj.episodeRepository.SaveEpisode(newEpisode)
	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	if episode.EpisodeId == 0 {
		return util.ResponseBadRequest("Thêm mới thất bại", nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *EpisodeController) UpdateEpisodesByEpisodeId(c *fiber.Ctx) error {
	episodeId := c.Params("episodeId")

	episode, err := validator.ValidateEpisodeId(episodeId)
	if err != nil {
		return err
	}

	episodeRequest := new(request.EpisodeRequest)
	if err = c.BodyParser(episodeRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	episode.Name = episodeRequest.Name
	episode.Status = episodeRequest.Status

	episode, err = obj.episodeRepository.UpdateEpisode(episodeId, *episode)
	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	if episode.EpisodeId == 0 {
		return util.ResponseBadRequest("Thêm mới thất bại", nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

func (obj *EpisodeController) DeleteEpisodesByEpisodeId(c *fiber.Ctx) error {
	episodeId := c.Params("episodeId")

	_, err := validator.ValidateEpisodeId(episodeId)
	if err != nil {
		return err
	}

	// Update episode status
	if err = obj.episodeRepository.UpdateStatusByEpisodeId(episodeId, util.StatusDeleted); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	// Delete episodeDetails
	if err = obj.episodeDetailRepository.UpdateStatusByEpisodeId(episodeId, util.StatusDeleted); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	return util.ResponseSuccess("Thành công", nil)
}

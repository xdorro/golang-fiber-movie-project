package controller

import "sync"

var (
	once *sync.Once

	tagController *TagController

	genreController *GenreController

	episodeController *EpisodeController

	episodeTypeController *EpisodeTypeController

	movieController *MovieController
)
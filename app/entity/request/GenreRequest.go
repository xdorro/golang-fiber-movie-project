package request

type GenreRequest struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Status int    `json:"status"`
}

package models

type JsonResponse struct {
	Songs  []Song   `json:"songs"`
	Movies []Movie  `json:"movies"`
	TvShow []TvShow `json:"tvShow"`
	Person []Person `json:"person"`
}

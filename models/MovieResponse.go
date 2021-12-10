package models

type results struct {
	ArtistName string  `json:"artistName"`
	TrackName  string  `json:"trackName"`
	TrackPrice float64 `json:"trackPrice"`
	Kind       string  `json:"kind"`
}

type MovieResponse struct {
	ResultCont int       `json:"resultCount"`
	Results    []results `json:"results"`
}

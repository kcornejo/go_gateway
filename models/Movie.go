package models

type Movie struct {
	ArtistName string  `json:"artistName"`
	TrackName  string  `json:"trackName"`
	TrackPrice float64 `json:"trackPrice"`
}

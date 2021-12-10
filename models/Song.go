package models

type Song struct {
	ArtistName string  `json:"artistName"`
	TrackName  string  `json:"trackName"`
	TrackPrice float64 `json:"trackPrice"`
}

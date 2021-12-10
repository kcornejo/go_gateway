package models

type show struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Status     string  `json:"status"`
	Language   string  `json:"language"`
	Genres     string  `json:"genres"`
	ScoreMatch float64 `json:"scoreMatch"`
	Premiered  string  `json:"premiered"`
	Ended      string  `json:"ended"`
}
type TvShowResponse struct {
	Score float64 `json:"score"`
	Show  show    `json:"show"`
}

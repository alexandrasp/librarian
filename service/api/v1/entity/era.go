package entity

type Era struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	MinYear int32  `json:"minYear"`
	MaxYear int32  `json:"maxYear"`
}

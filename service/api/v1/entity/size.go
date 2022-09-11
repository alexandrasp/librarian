package entity

type Size struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	MinPages int32  `json:"minPages"`
	MaxPages int32  `json:"maxPages"`
}

package models

type Buses struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Start string `json:"start"`
	End string `json:"end"`
}
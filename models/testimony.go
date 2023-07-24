package models

type Testimony struct {
	Id          int    `json:"id"`
	User        string `json:"user"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

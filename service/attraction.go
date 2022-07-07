package service

import "time"

// Attraction is view spot that contains id, visitors, etc.
type Attraction struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Visitors    int64     `json:"visitors"`
	Square      float32   `json:"square"`
	Heritage    bool      `json:"heritage"`
	CreateAt    time.Time `json:"create_at"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Link        string    `json:"link"`
	States      string    `json:"state"`
}

// Search defines the interface exposed by this package.
type Search interface {
	GetAttractions(name string) ([]Attraction, error)
}

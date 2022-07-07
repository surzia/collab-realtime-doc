package service

import "time"

type Service struct {
	// TODO: connect to elasticsearch
	db []Attraction
}

// NewService instantiates a new Service.
func NewService() Search {
	// TODO: init db
	db := []Attraction{
		{
			Id:          "park_carlsbad-caverns",
			Title:       "Carlsbad Caverns",
			Visitors:    466773,
			Square:      189.3,
			Heritage:    true,
			CreateAt:    time.Date(1930, time.May, 14, 0, 0, 0, 0, time.Local),
			Image:       "https://storage.googleapis.com/public-demo-assets.swiftype.info/swiftype-dot-com-search-ui-national-parks-demo/1CF45171-AA55-D49A-F7F2E2AF45DBEB8E.jpeg",
			Description: "Carlsbad Caverns has 117 caves, the longest of which is over 120 miles (190 km) long. The Big Room is almost 4,000 feet (1,200 m) long, and the caves are home to over 400,000 Mexican free-tailed bats and sixteen other species. Above ground are the Chihuahuan Desert and Rattlesnake Springs.",
			Location:    "32.17,-104.44",
			Link:        "https://www.nps.gov/cave/index.htm",
			States:      "New Mexico",
		},
	}
	return &Service{
		db: db,
	}
}

func (s *Service) GetAttractions(name string) ([]Attraction, error) {
	return s.db, nil
}

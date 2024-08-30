package entity

import "fmt"

type Location struct {
	Id          int    `db:"id"`
	Name        string `json:"name" db:"name"`
	Country     string `json:"country" db:"country"`
	NearestTown string `json:"nearest_town" db:"nearest_town"`
}

type Locations []*Location

type CreateLocationInput struct {
	Name        string `json:"name"`
	Country     string `json:"country"`
	NearestTown string `json:"nearest_town"`
}

func (input *CreateLocationInput) IsValid() error {
	var err error

	switch {
	case input.Name == "":
		err = fmt.Errorf("invalid location name")
	case input.Country == "":
		err = fmt.Errorf("invalid location country")
	case input.NearestTown == "":
		err = fmt.Errorf("invalid location nearest town")
	}

	return err
}

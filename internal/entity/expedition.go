package entity

import (
	"fmt"
	"time"
)

type Expedition struct {
	Id         int       `db:"id"`
	LocationId int       `json:"location_id" db:"location_id"`
	StartDate  time.Time `json:"start_date" db:"start_date"`
	EndDate    time.Time `json:"end_date" db:"end_date"`
}

type Expeditions []*Expedition

type CreateExpeditionInput struct {
	LocationId int    `json:"location_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

func (input *CreateExpeditionInput) IsValid() error {
	var err error

	switch {
	case input.StartDate == "":
		err = fmt.Errorf("invalid expedition start date")
	case input.EndDate == "":
		err = fmt.Errorf("invalid expedition end date")
	case input.StartDate >= input.EndDate:
		err = fmt.Errorf("invalid expedition dates")
	}

	return err
}

package entity

import "fmt"

type Equipment struct {
	Id           int    `db:"id"`
	ExpeditionId int    `json:"expedition_id" db:"expedition_id"`
	Name         string `json:"name" db:"name"`
	Amount       int    `json:"amount" db:"amount"`
}

type Equipments []*Equipment

type CreateEquipmentInput struct {
	ExpeditionId int    `json:"expedition_id"`
	Name         string `json:"name"`
	Amount       int    `json:"amount" db:"amount"`
}

func (input *CreateEquipmentInput) IsValid() error {
	var err error

	switch {
	case input.Name == "":
		err = fmt.Errorf("invalid equipment name")
	case input.Amount < 1:
		err = fmt.Errorf("invalid equipment amount")
	}

	return err
}

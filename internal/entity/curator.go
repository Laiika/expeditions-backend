package entity

import "fmt"

type Curator struct {
	Id   int    `db:"id"`
	Name string `json:"name" db:"name"`
}

type Curators []*Curator

type CreateCuratorInput struct {
	Name string `json:"name"`
}

func (input *CreateCuratorInput) IsValid() error {
	var err error

	if input.Name == "" {
		err = fmt.Errorf("invalid curator name")
	}

	return err
}

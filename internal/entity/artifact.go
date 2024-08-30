package entity

import "fmt"

type Artifact struct {
	Id         int    `db:"id"`
	LocationId int    `json:"location_id" db:"location_id"`
	Name       string `json:"name" db:"name"`
	Age        int    `json:"age" db:"age"`
}

type Artifacts []*Artifact

type CreateArtifactInput struct {
	LocationId int    `json:"location_id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
}

func (input *CreateArtifactInput) IsValid() error {
	var err error

	switch {
	case input.Name == "":
		err = fmt.Errorf("invalid artifact name")
	case input.Age < 1:
		err = fmt.Errorf("invalid artifact age")
	}

	return err
}

package entity

import "fmt"

type Leader struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
	Login       string `db:"login"`
	Password    string `db:"password"`
}

type Leaders []*Leader

type CreateLeaderInput struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Login       string `json:"login"`
	Password    string `json:"password"`
}

func (input *CreateLeaderInput) IsValid() error {
	var err error

	switch {
	case input.Name == "":
		err = fmt.Errorf("invalid leader name")
	case input.PhoneNumber == "":
		err = fmt.Errorf("invalid leader phone number")
	case input.Login == "":
		err = fmt.Errorf("invalid leader login")
	case input.Password == "":
		err = fmt.Errorf("invalid leader password")
	}

	return err
}

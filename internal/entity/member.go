package entity

import "fmt"

type Member struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
	Role        string `db:"role"`
	Login       string `db:"login"`
	Password    string `db:"password"`
}

type Members []*Member

type CreateMemberInput struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	Login       string `json:"login"`
	Password    string `json:"password"`
}

func (input *CreateMemberInput) IsValid() error {
	var err error

	switch {
	case input.Name == "":
		err = fmt.Errorf("invalid member name")
	case input.PhoneNumber == "":
		err = fmt.Errorf("invalid member phone number")
	case input.Role == "":
		err = fmt.Errorf("invalid member role")
	case input.Login == "":
		err = fmt.Errorf("invalid member login")
	case input.Password == "":
		err = fmt.Errorf("invalid member password")
	}

	return err
}

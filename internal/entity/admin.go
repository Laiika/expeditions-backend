package entity

import "fmt"

type Admin struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

type Admins []*Admin

type LoginInput struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required,password"`
}

type CreateAdminInput struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (input *CreateAdminInput) IsValid() error {
	var err error

	switch {
	case input.Name == "":
		err = fmt.Errorf("invalid admin name")
	case input.Login == "":
		err = fmt.Errorf("invalid admin login")
	case input.Password == "":
		err = fmt.Errorf("invalid admin password")
	}

	return err
}

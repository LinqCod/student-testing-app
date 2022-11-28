package entity

import "github.com/linqcod/student-testing-app/app/pkg/security"

type User struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type UserDTO struct {
	Id       int64  `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type UserRegistrationDTO struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) EncryptPassword(password string) (string, error) {
	hashPassword, err := security.Hash(password)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

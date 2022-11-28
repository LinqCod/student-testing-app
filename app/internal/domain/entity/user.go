package entity

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

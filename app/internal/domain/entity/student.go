package entity

type Student struct {
	User
	PersonalNumber string `json:"personal_number"`
	Group          Group  `json:"group"`
}

type StudentDTO struct {
	Id             int64  `json:"id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	PersonalNumber string `json:"personal_number"`
	Group          string `json:"group"`
}

type StudentRegistrationDTO struct {
	FirstName      string `json:"first_name"`
	SecondName     string `json:"second_name"`
	Patronymic     string `json:"patronymic"`
	PersonalNumber string `json:"personal_number"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Group          Group  `json:"group"`
}

type StudentLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

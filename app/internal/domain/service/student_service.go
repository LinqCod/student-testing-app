package service

import (
	"fmt"
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
	"github.com/linqcod/student-testing-app/app/internal/domain/repository"
	"github.com/linqcod/student-testing-app/app/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

type StudentService interface {
	SaveStudent(registration *entity.StudentRegistrationDTO) (*entity.StudentDTO, error)
	GetStudentDetails(studentId int64) (*entity.StudentDTO, error)
	UpdateStudent(student *entity.Student) (*entity.StudentDTO, error)
	GetStudentByEmail(login entity.StudentLoginDTO) (*entity.Student, error)
}

type StudentServiceImpl struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(studentRepo repository.StudentRepository) *StudentServiceImpl {
	return &StudentServiceImpl{
		studentRepo: studentRepo,
	}
}

func (s StudentServiceImpl) SaveStudent(registration *entity.StudentRegistrationDTO) (*entity.StudentDTO, error) {
	var student = entity.Student{
		User: entity.User{
			FullName: fmt.Sprintf("%s %s %s", registration.FirstName, registration.SecondName, registration.Patronymic),
			Email:    registration.Email,
			Role:     entity.StudentRole,
		},
		PersonalNumber: registration.PersonalNumber,
		Group:          registration.Group,
	}

	password, err := security.Hash(registration.Password)
	if err != nil {
		return nil, err
	}

	student.Password = password

	result, err := s.studentRepo.SaveStudent(&student)
	if err != nil {
		return nil, err
	}

	return &entity.StudentDTO{
		Id:             result.Id,
		FullName:       result.FullName,
		Email:          result.Email,
		PersonalNumber: result.PersonalNumber,
		Group:          result.Group.Title,
	}, nil
}

func (s StudentServiceImpl) GetStudentDetails(studentId int64) (*entity.StudentDTO, error) {
	result, err := s.studentRepo.GetStudentDetails(studentId)
	if err != nil {
		return nil, err
	}

	return &entity.StudentDTO{
		Id:             result.Id,
		FullName:       result.FullName,
		Email:          result.Email,
		PersonalNumber: result.PersonalNumber,
		Group:          result.Group.Title,
	}, nil
}

func (s StudentServiceImpl) UpdateStudent(student *entity.Student) (*entity.StudentDTO, error) {
	password, err := security.Hash(student.Password)
	if err != nil {
		return nil, err
	}

	student.Password = password

	result, err := s.studentRepo.UpdateStudent(student)
	if err != nil {
		return nil, err
	}

	return &entity.StudentDTO{
		Id:             result.Id,
		FullName:       result.FullName,
		Email:          result.Email,
		PersonalNumber: result.PersonalNumber,
		Group:          result.Group.Title,
	}, nil
}

func (s StudentServiceImpl) GetStudentByEmail(login entity.StudentLoginDTO) (*entity.Student, error) {
	result, err := s.studentRepo.GetStudentByEmail(login)
	if err != nil {
		return nil, err
	}

	err = security.VerifyPassword(result.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, fmt.Errorf("incorrect password. Error: %s", err.Error())
	}

	return result, nil
}

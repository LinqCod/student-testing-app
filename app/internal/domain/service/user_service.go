package service

import (
	"fmt"
	"github.com/linqcod/student-testing-app/app/internal/domain/entity"
	"github.com/linqcod/student-testing-app/app/internal/domain/repository"
	"github.com/linqcod/student-testing-app/app/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SaveUser(registration *entity.UserRegistrationDTO) (*entity.UserDTO, error)
	GetUserDetails(userId int64) (*entity.UserDTO, error)
	GetAllUsers() ([]entity.UserDTO, error)
	UpdateUser(user *entity.User) (*entity.UserDTO, error)
	DeleteUser(userId int64) error
	GetUserByEmail(login entity.UserLoginDTO) (*entity.User, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) getUserFullName(user entity.User) string {
	return fmt.Sprintf("%s %s", user.FirstName, user.SecondName)
}

func (s *UserServiceImpl) SaveUser(registration *entity.UserRegistrationDTO) (*entity.UserDTO, error) {
	var user = entity.User{
		FirstName:  registration.FirstName,
		SecondName: registration.SecondName,
		Email:      registration.Email,
	}

	password, err := user.EncryptPassword(registration.Password)
	if err != nil {
		return nil, err
	}

	user.Password = password

	result, err := s.userRepo.SaveUser(&user)
	if err != nil {
		return nil, err
	}

	return &entity.UserDTO{
		Id:       result.Id,
		FullName: s.getUserFullName(*result),
		Email:    result.Email,
	}, nil
}

func (s *UserServiceImpl) GetUserDetails(userId int64) (*entity.UserDTO, error) {
	result, err := s.userRepo.GetUserDetails(userId)
	if err != nil {
		return nil, err
	}

	return &entity.UserDTO{
		Id:       result.Id,
		FullName: s.getUserFullName(*result),
		Email:    result.Email,
	}, nil
}

func (s *UserServiceImpl) GetAllUsers() ([]entity.UserDTO, error) {
	result, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var users []entity.UserDTO
	for _, user := range result {
		users = append(users, entity.UserDTO{
			Id:       user.Id,
			FullName: s.getUserFullName(user),
			Email:    user.Email,
		})
	}

	return users, nil
}

func (s *UserServiceImpl) UpdateUser(user *entity.User) (*entity.UserDTO, error) {
	password, err := user.EncryptPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = password

	result, err := s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &entity.UserDTO{
		Id:       result.Id,
		FullName: s.getUserFullName(*result),
		Email:    result.Email,
	}, nil
}

func (s *UserServiceImpl) DeleteUser(userId int64) error {
	err := s.userRepo.DeleteUser(userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImpl) GetUserByEmail(login entity.UserLoginDTO) (*entity.User, error) {
	result, err := s.userRepo.GetUserByEmail(login)
	if err != nil {
		return nil, err
	}

	err = security.VerifyPassword(result.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, fmt.Errorf("incorrect password. Error: %s", err.Error())
	}

	return result, nil
}

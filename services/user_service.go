package services

import (
	"myapi/models"
	"myapi/pkg/hashing"
	"myapi/storage"
)

type UserService struct {
    Storage *storage.UserStorage
}

func NewUserService(storage *storage.UserStorage) *UserService {
    return &UserService{
        Storage: storage,
    }
}

func (s *UserService) GetAllUsers() []models.User {
    return s.Storage.GetAllUsers()
}

func (s *UserService) GetUserByID(id string) (models.User, bool) {
    return s.Storage.GetUserByID(id)
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	hashedPassword, err := hashing.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashedPassword
	return s.Storage.CreateUser(user), nil
}

func (s *UserService) UpdateUser(user models.User) (models.User, error) {
    updatedUser, err := s.Storage.UpdateUser(user)
    if err != nil {
        return models.User{}, err
    }
    return updatedUser, nil
}

func (s *UserService) DeleteUser(id string) error {
    err := s.Storage.DeleteUser(id)
    if err != nil {
        return err
    }
    return nil
}

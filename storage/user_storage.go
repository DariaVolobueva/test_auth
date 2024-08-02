package storage

import (
	"errors"
	"myapi/models"
	"strconv"
)

type UserStorage struct {
    users   []models.User
    nextID  int
}

func NewUserStorage() *UserStorage {
    return &UserStorage{
        users: []models.User{
            {ID: "1", Email: "alice@example.com", Password: "$2a$10$7HU18HH.QSiStY25oT/GNOW5byFOEQ3R14LMj6OvhghAjkSMciEC6", Role: "admin"}, // password: "password"
            {ID: "2", Email: "bob@example.com", Password: "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjr9k68GJ/3qF9F0fO/kXt1yP/Qc4aK", Role: "user"},  // password: "password"
        },
        nextID: 3,
    }
}

func (s *UserStorage) FindUserByEmail(email string) (models.User, bool) {
    for _, user := range s.users {
        if user.Email == email {
            return user, true
        }
    }
    return models.User{}, false
}

func (s *UserStorage) CreateUser(user models.User) models.User {
    user.ID = strconv.Itoa(s.nextID)
    s.nextID++
    s.users = append(s.users, user)
    return user
}

func (s *UserStorage) GetAllUsers() []models.User {
    return s.users
}

func (s *UserStorage) GetUserByID(id string) (models.User, bool) {
    for _, user := range s.users {
        if user.ID == id {
            return user, true
        }
    }
    return models.User{}, false
}

func (s *UserStorage) UpdateUser(updatedUser models.User) (models.User, error) {
    for i, user := range s.users {
        if user.ID == updatedUser.ID {
            s.users[i] = updatedUser
            return updatedUser, nil
        }
    }
    return models.User{}, errors.New("user not found")
}

func (s *UserStorage) DeleteUser(id string) error {
    for i, user := range s.users {
        if user.ID == id {
            s.users = append(s.users[:i], s.users[i+1:]...)
            return nil
        }
    }
    return errors.New("user not found")
}

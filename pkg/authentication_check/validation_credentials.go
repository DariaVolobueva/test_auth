package authentication_check

import (
	"errors"
	"myapi/models"
	"myapi/pkg/hashing"
	"myapi/storage"
)

func ValidateCredentials(u *models.User, s *storage.UserStorage) error {
    user, found := s.FindUserByEmail(u.Email)
    if !found {
        return errors.New("credentials invalid")
    }

    passwordIsValid := hashing.CheckPasswordHash(u.Password, user.Password)
    if !passwordIsValid {
        return errors.New("credentials password invalid")
    }

    u.ID = user.ID
    u.Role = user.Role
    return nil
}

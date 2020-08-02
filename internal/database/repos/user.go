package repos

import (
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func CreateUser(db *gorm.DB, username, email, encryptedPasswordHash string) (*entities.User, error) {
	user := &entities.User{
		Username:      username,
		Email:         email,
		PasswordHash:  encryptedPasswordHash,
		Role:          constants.ROLE_USER,
		ProfileURL:    utils.RandomProfileURL(),
		EmailVerified: false,
	}
	err := db.Model(&user).Create(user).Error
	if err != nil {
		log.E("Failed to create user: %v", err)
		return nil, fmt.Errorf("Failed to create user")
	}
	preferences := mappers.DefaultPreferences(user.ID)
	err = db.Model(&preferences).Create(preferences).Error
	if err != nil {
		log.E("Failed to create preferences: %v", err)
		return nil, fmt.Errorf("Failed to create preferences")
	}
	return user, nil
}

func VerifyUserEmail(db *gorm.DB, existingUser *entities.User) (*entities.User, error) {
	existingUser.EmailVerified = true
	err := db.Model(existingUser).Update(*existingUser).Error
	if err != nil {
		log.E("Failed to update user for [%+v]: %v", existingUser, err)
		return nil, fmt.Errorf("Failed to update user with id='%s'", existingUser.ID)
	}
	return existingUser, err
}

func FindUserByID(db *gorm.DB, userID string) (*entities.User, error) {
	user := &entities.User{}
	err := db.Unscoped().Where("id = ?", userID).Find(user).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No user found with id='%s'", userID)
	}
	return user, nil
}

func FindUserByUsername(db *gorm.DB, username string) (*entities.User, error) {
	user := &entities.User{}
	err := db.Where("username = ?", username).Find(user).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No user found with username='%s'", username)
	}
	return user, nil
}

func FindUserByEmail(db *gorm.DB, email string) (*entities.User, error) {
	user := &entities.User{}
	err := db.Where("email = ?", email).Find(user).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No user found with email='%s'", email)
	}
	return user, nil
}

func FindUserByUsernameOrEmail(db *gorm.DB, usernameOrEmail string) (*entities.User, error) {
	user := &entities.User{}
	err := db.Where("email = ? OR username = ?", usernameOrEmail, usernameOrEmail).Find(user).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No user found with email or username = '%s'", usernameOrEmail)
	}
	return user, nil
}

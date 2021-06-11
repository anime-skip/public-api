package repos

import (
	"strings"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/constants"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func CreateUser(db *gorm.DB, username, email, encryptedPasswordHash string) (*entities.User, error) {
	user := &entities.User{
		Username:      username,
		Email:         strings.ToLower(email),
		PasswordHash:  encryptedPasswordHash,
		Role:          constants.ROLE_USER,
		ProfileURL:    utils.RandomProfileURL(),
		EmailVerified: false,
	}
	err := db.Model(&user).Create(user).Error
	if err != nil {
		log.E("Failed to create user: %v", err)
		return nil, err
	}
	preferences := mappers.DefaultPreferences(user.ID)
	err = db.Model(&preferences).Create(preferences).Error
	if err != nil {
		log.E("Failed to create preferences: %v", err)
		return nil, err
	}
	return user, nil
}

func VerifyUserEmail(db *gorm.DB, existingUser *entities.User) (*entities.User, error) {
	existingUser.EmailVerified = true
	err := db.Model(existingUser).Update(*existingUser).Error
	if err != nil {
		log.E("Failed to update user for [%+v]: %v", existingUser, err)
		return nil, err
	}
	return existingUser, err
}

func FindUserByID(db *gorm.DB, userID string) (*entities.User, error) {
	user := &entities.User{}
	err := db.Where("id = ?", userID).Find(user).Error
	if err != nil {
		log.E("No user found with id='%s': %v", userID, err)
		return nil, err
	}
	return user, nil
}

func FindUserByUsername(db *gorm.DB, username string) (*entities.User, error) {
	user := &entities.User{}
	err := db.Where("username = ?", username).Find(user).Error
	if err != nil {
		log.E("No user found with username='%s': %v", username, err)
		return nil, err
	}
	return user, nil
}

func FindUserByEmail(db *gorm.DB, email string) (*entities.User, error) {
	user := &entities.User{}
	err := db.Where("email = ?", strings.ToLower(email)).Find(user).Error
	if err != nil {
		log.E("No user found with email='%s': %v", email, err)
		return nil, err
	}
	return user, nil
}

func FindUserByUsernameOrEmail(db *gorm.DB, usernameOrEmail string) (*entities.User, error) {
	user := &entities.User{}
	err := db.Where("email = ? OR username = ?", strings.ToLower(usernameOrEmail), usernameOrEmail).Find(user).Error
	if err != nil {
		log.E("No user found with email or username='%s': %v", usernameOrEmail, err)
		return nil, err
	}
	return user, nil
}

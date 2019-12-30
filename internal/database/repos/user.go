package repos

import (
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

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

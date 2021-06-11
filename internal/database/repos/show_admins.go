package repos

import (
	"errors"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func CreateShowAdmin(db *gorm.DB, showInput models.InputShowAdmin) (*entities.ShowAdmin, error) {
	showAdmin := mappers.ShowAdminInputModelToEntity(showInput, &entities.ShowAdmin{})
	err := db.Model(&showAdmin).Create(showAdmin).Error
	if err != nil {
		log.E("Failed to create show admin with [%+v]: %v", showInput, err)
		return nil, err
	}
	return showAdmin, nil
}

func UpdateShowAdmin(db *gorm.DB, newShowAdmin models.InputShowAdmin, existingShowAdmin *entities.ShowAdmin) (*entities.ShowAdmin, error) {
	data := mappers.ShowAdminInputModelToEntity(newShowAdmin, existingShowAdmin)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update show admin for [%+v]: %v", data, err)
		return nil, err
	}
	return data, err
}

func DeleteShowAdmin(db *gorm.DB, showAdminID string) error {
	err := db.Delete(entities.ShowAdmin{}, "id = ?", showAdminID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.E("Failed to delete show admin for id='%s': %v", showAdminID, err)
		return err
	}
	return nil
}

func FindShowAdminByID(db *gorm.DB, showAdminID string) (*entities.ShowAdmin, error) {
	showAdmin := &entities.ShowAdmin{}
	err := db.Where("id = ?", showAdminID).Find(showAdmin).Error
	if err != nil {
		log.V("No show admin found with id='%s' (%v)", showAdminID, err)
		return nil, err
	}
	return showAdmin, nil
}

func FindShowAdminsByUserID(db *gorm.DB, userID string) ([]*entities.ShowAdmin, error) {
	showAdmin := []*entities.ShowAdmin{}
	err := db.Where("user_id = ?", userID).Find(&showAdmin).Error
	if err != nil {
		log.V("Query for show admins for user_id='%s' failed: %v", userID, err)
		return nil, err
	}
	return showAdmin, nil
}

func FindShowAdminsByShowID(db *gorm.DB, showID string) ([]*entities.ShowAdmin, error) {
	showAdmin := []*entities.ShowAdmin{}
	err := db.Where("show_id = ?", showID).Find(&showAdmin).Error
	if err != nil {
		log.V("Query for show admins for show_id='%s' failed: %v", showID, err)
		return nil, err
	}
	return showAdmin, nil
}

func FindShowAdminsByUserIDShowID(db *gorm.DB, userID string, showID string) (*entities.ShowAdmin, error) {
	showAdmin := &entities.ShowAdmin{}
	err := db.Where("user_id = ? AND show_id = ?", userID, showID).Find(&showAdmin).Error
	if err != nil {
		log.V("Query for show admins for user_id='%s' AND show_id='%s' failed: %v", userID, showID, err)
		return nil, err
	}
	return showAdmin, nil
}

package repos

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func CreateShowAdmin(ctx context.Context, db *gorm.DB, showInput models.InputShowAdmin) (*entities.ShowAdmin, error) {
	showAdmin := mappers.ShowAdminInputModelToEntity(showInput, &entities.ShowAdmin{})
	err := db.Model(&showAdmin).Create(showAdmin).Error
	if err != nil {
		log.E("Failed to create show admin with [%+v]: %v", showInput, err)
		return nil, fmt.Errorf("Failed to create show admin: %v", err)
	}
	return showAdmin, nil
}

func UpdateShowAdmin(ctx context.Context, db *gorm.DB, newShowAdmin models.InputShowAdmin, existingShowAdmin *entities.ShowAdmin) (*entities.ShowAdmin, error) {
	data := mappers.ShowAdminInputModelToEntity(newShowAdmin, existingShowAdmin)
	err := db.Model(data).Update(*data).Error
	if err != nil {
		log.E("Failed to update show admin for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update show admin with id='%s'", data.ID)
	}
	return data, err
}

func DeleteShowAdmin(ctx context.Context, db *gorm.DB, showAdmin *entities.ShowAdmin) error {
	err := db.Model(showAdmin).Delete(showAdmin).Error
	if err != nil {
		log.E("Failed to delete show admin for id='%s': %v", showAdmin.ID, err)
		return fmt.Errorf("Failed to delete show admin with id='%s'", showAdmin.ID)
	}
	return err
}

func FindShowAdminByID(ctx context.Context, db *gorm.DB, showAdminID string) (*entities.ShowAdmin, error) {
	showAdmin := &entities.ShowAdmin{}
	err := db.Where("id = ?", showAdminID).Find(showAdmin).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No show admin found with id='%s'", showAdminID)
	}
	return showAdmin, nil
}

func FindShowAdminsByUserID(ctx context.Context, db *gorm.DB, userID string) ([]*entities.ShowAdmin, error) {
	showAdmin := []*entities.ShowAdmin{}
	err := db.Where("user_id = ?", userID).Find(&showAdmin).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("Query for show admins for user_id='%s' failed", userID)
	}
	return showAdmin, nil
}

func FindShowAdminsByShowID(ctx context.Context, db *gorm.DB, showID string) ([]*entities.ShowAdmin, error) {
	showAdmin := []*entities.ShowAdmin{}
	err := db.Where("show_id = ?", showID).Find(&showAdmin).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("Query for show admins for show_id='%s' failed", showID)
	}
	return showAdmin, nil
}

func FindShowAdminsByUserIDShowID(ctx context.Context, db *gorm.DB, userID string, showID string) (*entities.ShowAdmin, error) {
	showAdmin := &entities.ShowAdmin{}
	err := db.Where("user_id = ? AND show_id = ?", userID, showID).Find(&showAdmin).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("Query for show admins for user_id='%s' AND show_id='%s' failed", userID, showID)
	}
	return showAdmin, nil
}

package repos

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
)

func CreateShowAdmin(ctx context.Context, orm *database.ORM, showInput models.InputShowAdmin) (*entities.ShowAdmin, error) {
	showAdmin := mappers.ShowAdminInputModelToEntity(showInput, &entities.ShowAdmin{})
	err := orm.DB.Model(&showAdmin).Create(showAdmin).Error
	if err != nil {
		log.E("Failed to create show admin with [%+v]: %v", showInput, err)
		return nil, fmt.Errorf("Failed to create show admin: %v", err)
	}
	return showAdmin, nil
}

func UpdateShowAdmin(ctx context.Context, orm *database.ORM, newShowAdmin models.InputShowAdmin, existingShowAdmin *entities.ShowAdmin) (*entities.ShowAdmin, error) {
	data := mappers.ShowAdminInputModelToEntity(newShowAdmin, existingShowAdmin)
	err := orm.DB.Model(data).Update(*data).Error
	if err != nil {
		log.E("Failed to update show admin for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update show admin with id='%s'", data.ID)
	}
	return data, err
}

func DeleteShowAdmin(ctx context.Context, orm *database.ORM, showAdmin *entities.ShowAdmin) error {
	err := orm.DB.Model(showAdmin).Delete(showAdmin).Error
	if err != nil {
		log.E("Failed to delete show admin for id='%s': %v", showAdmin.ID, err)
		return fmt.Errorf("Failed to delete show admin with id='%s'", showAdmin.ID)
	}
	return err
}

func FindShowAdminByID(ctx context.Context, orm *database.ORM, showAdminID string) (*entities.ShowAdmin, error) {
	showAdmin := &entities.ShowAdmin{}
	err := orm.DB.Where("id = ?", showAdminID).Find(showAdmin).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No show admin found with id='%s'", showAdminID)
	}
	return showAdmin, nil
}

func FindShowAdminsByUserID(ctx context.Context, orm *database.ORM, userID string) ([]*entities.ShowAdmin, error) {
	showAdmin := []*entities.ShowAdmin{}
	err := orm.DB.Where("user_id = ?", userID).Find(&showAdmin).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("Query for show admins for user_id='%s' failed", userID)
	}
	return showAdmin, nil
}

func FindShowAdminsByShowID(ctx context.Context, orm *database.ORM, showID string) ([]*entities.ShowAdmin, error) {
	showAdmin := []*entities.ShowAdmin{}
	err := orm.DB.Where("show_id = ?", showID).Find(&showAdmin).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("Query for show admins for show_id='%s' failed", showID)
	}
	return showAdmin, nil
}

func FindShowAdminsByUserIDShowID(ctx context.Context, orm *database.ORM, userID string, showID string) (*entities.ShowAdmin, error) {
	showAdmin := &entities.ShowAdmin{}
	err := orm.DB.Where("user_id = ? AND show_id = ?", userID, showID).Find(&showAdmin).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("Query for show admins for user_id='%s' AND show_id='%s' failed", userID, showID)
	}
	return showAdmin, nil
}

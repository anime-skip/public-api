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

func CreateShow(ctx context.Context, db *gorm.DB, showInput models.InputShow) (*entities.Show, error) {
	show := mappers.ShowInputModelToEntity(showInput, &entities.Show{})
	err := db.Model(&show).Create(show).Error
	if err != nil {
		log.E("Failed to create show with [%+v]: %v", showInput, err)
		return nil, fmt.Errorf("Failed to create show: %v", err)
	}
	return show, nil
}

func UpdateShow(ctx context.Context, db *gorm.DB, newShow models.InputShow, existingShow *entities.Show) (*entities.Show, error) {
	data := mappers.ShowInputModelToEntity(newShow, existingShow)
	err := db.Model(data).Update(*data).Error
	if err != nil {
		log.E("Failed to update show for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update show with id='%s'", data.ID)
	}
	return data, err
}

func DeleteShow(ctx context.Context, db *gorm.DB, show *entities.Show) error {
	err := db.Model(show).Delete(show).Error
	if err != nil {
		log.E("Failed to delete show for id='%s': %v", show.ID, err)
		return fmt.Errorf("Failed to delete show with id='%s'", show.ID)
	}
	return err
}

func FindShowByID(ctx context.Context, db *gorm.DB, showID string) (*entities.Show, error) {
	show := &entities.Show{}
	err := db.Where("id = ?", showID).Find(show).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No show found with id='%s'", showID)
	}
	return show, nil
}

func FindShows(ctx context.Context, db *gorm.DB, search string, offset int, limit int, sort string) ([]*entities.Show, error) {
	shows := []*entities.Show{}
	searchVar := "%" + search + "%"
	err := db.Where("LOWER(name) LIKE LOWER(?) OR LOWER(original_name) LIKE LOWER(?)", searchVar, searchVar).Offset(offset).Limit(limit).Order("LOWER(name) " + sort).Find(&shows).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No shows found with name LIKE '%s'", search)
	}
	return shows, nil
}

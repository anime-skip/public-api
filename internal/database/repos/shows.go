package repos

import (
	"fmt"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func CreateShow(db *gorm.DB, showInput models.InputShow) (*entities.Show, error) {
	show := mappers.ShowInputModelToEntity(showInput, &entities.Show{})
	err := db.Model(&show).Create(show).Error
	if err != nil {
		log.E("Failed to create show with [%+v]: %v", showInput, err)
		return nil, fmt.Errorf("Failed to create show: %v", err)
	}
	return show, nil
}

func UpdateShow(db *gorm.DB, newShow models.InputShow, existingShow *entities.Show) (*entities.Show, error) {
	data := mappers.ShowInputModelToEntity(newShow, existingShow)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update show for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update show with id='%s'", data.ID)
	}
	return data, err
}

func DeleteShow(db *gorm.DB, inTransaction bool, showID string) (err error) {
	tx := utils.StartTransaction(db, inTransaction)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Failed to delete show and it's admins: %+v", r)
			tx.Rollback()
		}
	}()

	// Delete the show
	err = tx.Delete(entities.Show{}, "id=?", showID).Error
	if err != nil {
		log.E("Failed to delete show for id='%s': %v", showID, err)
		tx.Rollback()
		return fmt.Errorf("Failed to delete show with id='%s'", showID)
	}

	// Delete the admins for that show
	admins, err := FindShowAdminsByShowID(tx, showID)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, admin := range admins {
		if err = DeleteShowAdmin(tx, admin.ID.String()); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Episodes
	episodes, err := FindEpisodesByShowID(tx, showID)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, episode := range episodes {
		if err = DeleteEpisode(tx, true, episode.ID.String()); err != nil {
			tx.Rollback()
			return err
		}
	}

	utils.CommitTransaction(tx, inTransaction)
	return nil
}

func FindShowByID(db *gorm.DB, showID string) (*entities.Show, error) {
	show := &entities.Show{}
	err := db.Unscoped().Where("id = ?", showID).Find(show).Error
	if err != nil {
		return nil, fmt.Errorf("No show found with id='%s'", showID)
	}
	return show, nil
}

func SearchShows(db *gorm.DB, search string, offset int, limit int, sort string) ([]*entities.Show, error) {
	shows := []*entities.Show{}
	searchVar := "%" + search + "%"
	var sortOrder string
	if sort == "ASC" {
		sortOrder = "LOWER(name) ASC"
	} else {
		sortOrder = "LOWER(name) DESC"
	}
	err := db.Where("LOWER(name) LIKE LOWER(?) OR LOWER(original_name) LIKE LOWER(?)", searchVar, searchVar).Offset(offset).Limit(limit).Order(sortOrder).Find(&shows).Error
	if err != nil {
		return nil, fmt.Errorf("No shows found with name LIKE '%s'", search)
	}
	return shows, nil
}

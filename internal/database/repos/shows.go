package repos

import (
	"errors"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func CreateShow(db *gorm.DB, showInput models.InputShow) (*entities.Show, error) {
	show := mappers.ShowInputModelToEntity(showInput, &entities.Show{})
	err := db.Model(&show).Create(show).Error
	if err != nil {
		log.E("Failed to create show with [%+v]: %v", showInput, err)
		return nil, err
	}
	return show, nil
}

func UpdateShow(db *gorm.DB, newShow models.InputShow, existingShow *entities.Show) (*entities.Show, error) {
	data := mappers.ShowInputModelToEntity(newShow, existingShow)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update show for [%+v]: %v", data, err)
		return nil, err
	}
	return data, err
}

func DeleteShow(tx *gorm.DB, showID string) error {
	// Delete the show
	err := tx.Delete(entities.Show{}, "id=?", showID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.E("Failed to delete show for id='%s': %v", showID, err)
		return err
	}

	// Delete the admins for that show
	admins, err := FindShowAdminsByShowID(tx, showID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	for _, admin := range admins {
		err = DeleteShowAdmin(tx, admin.ID.String())
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	// Delete timestamps attached to the show
	templates, err := FindTemplatesByShowID(tx, showID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	for _, template := range templates {
		err = DeleteTemplate(tx, template.ID.String())
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	// Episodes
	episodes, err := FindEpisodesByShowID(tx, showID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	for _, episode := range episodes {
		err = DeleteEpisode(tx, episode.ID.String())
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return nil
}

func FindShowByID(db *gorm.DB, showID string) (*entities.Show, error) {
	show := &entities.Show{}
	err := db.Where("id = ?", showID).Find(show).Error
	if err != nil {
		log.E("No show found with id='%s' (%v)", showID, err)
		return nil, err
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
		log.E("No shows found with name LIKE '%s': %v", search, err)
		return nil, err
	}
	return shows, nil
}

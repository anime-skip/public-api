package repos

import (
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

func CreateEpisode(db *gorm.DB, showID string, episodeInput models.InputEpisode) (*entities.Episode, error) {
	episode := mappers.EpisodeInputModelToEntity(episodeInput, &entities.Episode{
		ShowID: uuid.FromStringOrNil(showID),
	})
	err := db.Model(&episode).Create(episode).Error
	if err != nil {
		log.E("Failed to create episode with [%+v]: %v", episodeInput, err)
		return nil, fmt.Errorf("Failed to create episode: %v", err)
	}
	return episode, nil
}

func UpdateEpisode(db *gorm.DB, newEpisode models.InputEpisode, existingEpisode *entities.Episode) (*entities.Episode, error) {
	data := mappers.EpisodeInputModelToEntity(newEpisode, existingEpisode)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update episode for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update episode with id='%s'", data.ID)
	}
	return data, err
}

func DeleteEpisode(db *gorm.DB, inTransaction bool, episodeID string) (err error) {
	tx := utils.StartTransaction(db, inTransaction)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Failed to delete episode and it's admins: %+v", r)
			tx.Rollback()
		}
	}()

	// Delete the episode
	err = tx.Delete(entities.Episode{}, "id = ?", episodeID).Error
	if err != nil {
		log.E("Failed to delete episode for id='%s': %v", episodeID, err)
		tx.Rollback()
		return fmt.Errorf("Failed to delete episode with id='%s'", episodeID)
	}

	// Delete the timestamps for that episode
	timestamps, err := FindTimestampsByEpisodeID(tx, episodeID)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, timestamp := range timestamps {
		if err = DeleteTimestamp(tx, true, timestamp.ID.String()); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Delete the urls for that episode
	urls, err := FindEpisodeURLsByEpisodeID(tx, episodeID)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, url := range urls {
		if _, err = DeleteEpisodeURL(tx, true, url.URL); err != nil {
			tx.Rollback()
			return err
		}
	}

	utils.CommitTransaction(tx, inTransaction)
	return nil
}

func FindEpisodeByID(db *gorm.DB, episodeID string) (*entities.Episode, error) {
	episode := &entities.Episode{}
	err := db.Unscoped().Where("id = ?", episodeID).Find(episode).Error
	if err != nil {
		log.V("Failed query: %v", err)
		return nil, fmt.Errorf("No episode found with id='%s'", episodeID)
	}
	return episode, nil
}

func FindEpisodesByShowID(db *gorm.DB, showID string) ([]*entities.Episode, error) {
	episodes := []*entities.Episode{}
	err := db.Where("show_id = ?", showID).Order("season ASC, number ASC, absolute_number ASC").Find(&episodes).Error
	if err != nil {
		log.V("Failed query: %v", err)
		return nil, fmt.Errorf("No episodes found with show_id='%s'", showID)
	}
	return episodes, nil
}

func SearchEpisodes(db *gorm.DB, search string, showID *string, offset int, limit int, sort string) ([]*entities.Episode, error) {
	episodes := []*entities.Episode{}
	searchVars := []interface{}{"%" + search + "%"}
	queryString := "LOWER(name) LIKE LOWER(?)"
	if showID != nil {
		searchVars = append(searchVars, *showID)
		queryString += " AND show_id = ?"
	}
	err := db.Where(queryString, searchVars...).Offset(offset).Limit(limit).Order("LOWER(name) " + sort).Find(&episodes).Error
	if err != nil {
		log.V("Failed query: %v", err)
		return nil, fmt.Errorf("No episodes found with name LIKE '%s'", search)
	}
	return episodes, nil
}

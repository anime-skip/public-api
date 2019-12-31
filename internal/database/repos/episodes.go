package repos

import (
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
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
	err := db.Model(data).Update(*data).Error
	if err != nil {
		log.E("Failed to update episode for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update episode with id='%s'", data.ID)
	}
	return data, err
}

func DeleteEpisode(db *gorm.DB, episode *entities.Episode) (err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Failed to delete episode and it's admins: %+v", r)
			tx.Rollback()
		}
	}()

	// Delete the episode
	err = tx.Model(episode).Delete(episode).Error
	if err != nil {
		log.E("Failed to delete episode for id='%s': %v", episode.ID, err)
		tx.Rollback()
		return fmt.Errorf("Failed to delete episode with id='%s'", episode.ID)
	}

	log.W("TODO - Delete timestamps when deleting a episode")

	tx.Commit()
	return nil
}

func FindEpisodeByID(db *gorm.DB, episodeID string) (*entities.Episode, error) {
	episode := &entities.Episode{}
	err := db.Unscoped().Where("id = ?", episodeID).Find(episode).Error
	if err != nil {
		log.W("Failed query: %v", err)
		return nil, fmt.Errorf("No episode found with id='%s'", episodeID)
	}
	return episode, nil
}

func FindEpisodesByShowID(db *gorm.DB, showID string) ([]*entities.Episode, error) {
	episodes := []*entities.Episode{}
	err := db.Where("show_id = ?", showID).Order("season ASC, number ASC, absolute_number ASC").Find(&episodes).Error
	if err != nil {
		log.W("Failed query: %v", err)
		return nil, fmt.Errorf("No episodes found with show_id='%s'", showID)
	}
	return episodes, nil
}

func SearchEpisodes(db *gorm.DB, search string, offset int, limit int, sort string) ([]*entities.Episode, error) {
	episodes := []*entities.Episode{}
	searchVar := "%" + search + "%"
	err := db.Where("LOWER(name) LIKE LOWER(?)", searchVar).Offset(offset).Limit(limit).Order("LOWER(name) " + sort).Find(&episodes).Error
	if err != nil {
		log.W("Failed query: %v", err)
		return nil, fmt.Errorf("No episodes found with name LIKE '%s'", search)
	}
	return episodes, nil
}

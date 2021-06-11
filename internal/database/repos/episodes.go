package repos

import (
	"errors"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/log"
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
		return nil, err
	}
	return episode, nil
}

func UpdateEpisode(db *gorm.DB, newEpisode models.InputEpisode, existingEpisode *entities.Episode) (*entities.Episode, error) {
	data := mappers.EpisodeInputModelToEntity(newEpisode, existingEpisode)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update episode for [%+v]: %v", data, err)
		return nil, err
	}
	return data, err
}

func DeleteEpisode(tx *gorm.DB, episodeID string) error {
	// Delete the episode
	err := tx.Delete(entities.Episode{}, "id = ?", episodeID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.E("Failed to delete episode for id='%s': %v", episodeID, err)
		return err
	}

	// Delete the template if it exists
	template, err := FindTemplateBySourceEpisodeID(tx, episodeID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if template != nil {
		templateID := template.ID.String()
		err = DeleteTemplate(tx, templateID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	// Delete the timestamps for that episode
	timestamps, err := FindTimestampsByEpisodeID(tx, episodeID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	for _, timestamp := range timestamps {
		err = DeleteTimestamp(tx, timestamp.ID.String())
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	// Delete the urls for that episode
	urls, err := FindEpisodeURLsByEpisodeID(tx, episodeID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	for _, url := range urls {
		err = DeleteEpisodeURL(tx, url.URL)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return nil
}

func FindEpisodeByID(db *gorm.DB, episodeID string) (*entities.Episode, error) {
	episode := &entities.Episode{}
	err := db.Where("id = ?", episodeID).Find(episode).Error
	if err != nil {
		log.V("No episode found with id='%s': %v", episodeID, err)
		return nil, err
	}
	return episode, nil
}

func FindEpisodesByExactName(db *gorm.DB, name string) ([]*entities.Episode, error) {
	episodes := []*entities.Episode{}
	err := db.Where("name = ?", name).Find(&episodes).Error
	if err != nil {
		log.V("No episode found with name='%s': %v", name, err)
		return nil, err
	}
	return episodes, nil
}

func FindEpisodesByShowID(db *gorm.DB, showID string) ([]*entities.Episode, error) {
	episodes := []*entities.Episode{}
	err := db.Where("show_id = ?", showID).Order("season ASC, number ASC, absolute_number ASC").Find(&episodes).Error
	if err != nil {
		log.V("No episodes found with show_id='%s': %v", showID, err)
		return nil, err
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
	var sortOrder string
	if sort == "ASC" {
		sortOrder = "LOWER(name) ASC"
	} else {
		sortOrder = "LOWER(name) DESC"
	}
	err := db.Where(queryString, searchVars...).Offset(offset).Limit(limit).Order(sortOrder).Find(&episodes).Error
	if err != nil {
		return nil, err
	}
	return episodes, nil
}

func RecentlyAddedEpisodes(db *gorm.DB, limit, offset int) ([]*entities.Episode, error) {
	episodes := []*entities.Episode{}
	err := db.Raw(`
		SELECT * FROM (
			SELECT 
				episodes.*,
				(
					SELECT COUNT(*) FROM timestamps WHERE episode_id = episodes.id
				) AS "timestamp_count"
			FROM "episodes"
			LEFT JOIN timestamps ON timestamps.episode_id = episodes.id
			WHERE "episodes"."deleted_at" IS NULL
			GROUP BY timestamps.episode_id, episodes.id
			ORDER BY episodes.created_at DESC NULLS LAST
		) as episodes
		WHERE episodes.timestamp_count > 0
		LIMIT ?
		OFFSET ?;
	`, limit, offset).Scan(&episodes).Error

	if err != nil {
		log.V("Failed to select recent episodes: %v", err)
		return nil, err
	}
	return episodes, nil
}

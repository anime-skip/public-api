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

func CreateShow(ctx context.Context, orm *database.ORM, showInput models.InputShow) (*entities.Show, error) {
	show := mappers.ShowInputModelToEntity(showInput, &entities.Show{})
	err := orm.DB.Model(&show).Create(show).Error
	if err != nil {
		log.E("Failed to create show with [%+v]: %v", showInput, err)
		return nil, fmt.Errorf("Failed to create show: %v", err)
	}
	return show, nil
}

func UpdateShow(ctx context.Context, orm *database.ORM, newShow models.InputShow, existingShow *entities.Show) (*entities.Show, error) {
	data := mappers.ShowInputModelToEntity(newShow, existingShow)
	err := orm.DB.Model(data).Update(*data).Error
	if err != nil {
		log.E("Failed to update show for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update show with id='%s'", data.ID)
	}
	return data, err
}

func FindShowByID(ctx context.Context, orm *database.ORM, showID string) (*entities.Show, error) {
	show := &entities.Show{}
	err := orm.DB.Where("id = ?", showID).Find(show).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No show found with id='%s'", showID)
	}
	return show, nil
}

func FindShows(ctx context.Context, orm *database.ORM, search string, offset int, limit int, sort string) ([]*entities.Show, error) {
	shows := []*entities.Show{}
	err := orm.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+search+"%").Offset(offset).Limit(limit).Order("name " + sort).Find(&shows).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No shows found with name LIKE '%s'", search)
	}
	return shows, nil
}

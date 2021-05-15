package validation

import (
	"errors"
	"fmt"

	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/constants"
	"github.com/jinzhu/gorm"
)

func CreateTemplateInput(db *gorm.DB, templateInput models.InputTemplate) error {
	hasSeasons := templateInput.Seasons != nil && len(templateInput.Seasons) > 0

	// No seasons on a show template
	if templateInput.Type == models.TemplateTypeShow && hasSeasons {
		return errors.New("A show template cannot contain any seasons")
	}

	// Require seasons on a season template
	if templateInput.Type == models.TemplateTypeSeasons && !hasSeasons {
		return errors.New("A seasons template requires the at least 1 season")
	}

	existingTemplates, err := repos.FindTemplatesByShowID(db, templateInput.ShowID)
	if err != nil {
		return err
	}

	// Only 1 show template per show
	if templateInput.Type == models.TemplateTypeShow {
		for _, existingTemplate := range existingTemplates {
			if existingTemplate.Type == constants.TEMPLATE_TYPE_SHOW {
				return fmt.Errorf("Cannot create a show wide template, one already exists (id='%s')", existingTemplate.ID.String())
			}
		}
	}

	// Only 1 season template for a season
	if templateInput.Type == models.TemplateTypeSeasons {
		for _, existingTemplate := range existingTemplates {
			if existingTemplate.Type == constants.TEMPLATE_TYPE_SEASONS {
				for _, season := range templateInput.Seasons {
					if utils.StringArrayIncludes(existingTemplate.Seasons, season) {
						return fmt.Errorf("Cannot create a template for season '%s', another template already existing with that season (id='%s')", season, existingTemplate.ID.String())
					}
				}
			}
		}
	}

	return nil
}

package resolvers

import (
	"context"
	"fmt"
	"strings"

	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/log"
	"anime-skip.com/backend/internal/utils/validation"
	"github.com/jinzhu/gorm"
)

// Helpers

func templateByID(db *gorm.DB, templateID string) (*models.Template, error) {
	template, err := repos.FindTemplateByID(db, templateID)
	if err != nil {
		return nil, err
	}
	return mappers.TemplateEntityToModel(template), nil
}

func templateBySourceEpisodeID(db *gorm.DB, sourceEpisodeID string) (*models.Template, error) {
	template, err := repos.FindTemplateBySourceEpisodeID(db, sourceEpisodeID)
	if err != nil {
		return nil, err
	}
	return mappers.TemplateEntityToModel(template), nil
}

func templatesByShowID(db *gorm.DB, showID string) ([]*models.Template, error) {
	templates, err := repos.FindTemplatesByShowID(db, showID)
	if err != nil {
		return nil, err
	}
	templateModels := make([]*models.Template, len(templates))
	for index, episode := range templates {
		templateModels[index] = mappers.TemplateEntityToModel(episode)
	}
	return templateModels, nil
}

// Query Resolvers

func (r *queryResolver) FindTemplate(ctx context.Context, templateID string) (*models.Template, error) {
	return templateByID(r.DB(ctx).Unscoped(), templateID)
}

func (r *queryResolver) FindTemplatesByShowID(ctx context.Context, showID string) ([]*models.Template, error) {
	return templatesByShowID(r.DB(ctx), showID)
}

func (r *queryResolver) FindTemplateByDetails(ctx context.Context, episodeID *string, showName *string, season *string) (*models.Template, error) {
	db := r.DB(ctx)

	// Lookup by episode ID first
	if episodeID != nil {
		log.D("Looking up template by id")
		template, err := templateBySourceEpisodeID(db, *episodeID)
		if template != nil {
			return template, nil
		}
		log.D("Failed to get template by episode id: %v", err)
	}

	// Get templates by show name
	if showName != nil {
		log.V("Show name passed, looking up shows with the exact name '%s'", *showName)
		show, err := repos.SearchShows(db, *showName, 0, 1, "ASC")
		if err != nil {
			return nil, err
		}
		if len(show) < 1 {
			return nil, fmt.Errorf("Show name '%v' did not match any known shows", *showName)
		}
		templates, err := templatesByShowID(db, show[0].ID.String())
		if err != nil {
			return nil, err
		}

		if season != nil {
			log.V("Getting template by show name and season '%v'", *season)
			// When season is passed, return the template that includes that season
			for _, template := range templates {
				if template.Type == models.TemplateTypeSeasons && template.Seasons != nil && utils.StringArrayIncludes(template.Seasons, *season) {
					return template, nil
				}
			}
		}

		log.V("Getting template by just show name")
		// When season is not passed, return the template for the show
		for _, template := range templates {
			if template.Type == models.TemplateTypeShow {
				return template, nil
			}
		}
	}

	return nil, nil
}

// Mutation Resolvers

func (r *mutationResolver) CreateTemplate(ctx context.Context, newTemplate models.InputTemplate) (*models.Template, error) {
	db := r.DB(ctx)
	err := validation.CreateTemplateInput(db, newTemplate)
	if err != nil {
		return nil, err
	}

	template, err := repos.CreateTemplate(db, newTemplate)
	if err != nil {
		return nil, err
	}
	return mappers.TemplateEntityToModel(template), nil
}

func (r *mutationResolver) UpdateTemplate(ctx context.Context, templateID string, newTemplate models.InputTemplate) (*models.Template, error) {
	var err error
	tx, commitOrRollback := utils.StartTransaction2(r.DB(ctx), &err)
	defer commitOrRollback()

	err = validation.CreateTemplateInput(tx, newTemplate)
	if err != nil && !strings.Contains(err.Error(), templateID) {
		return nil, err
	}

	existingTemplate, err := repos.FindTemplateByID(tx, templateID)
	if err != nil {
		return nil, err
	}
	updatedTemplate, err := repos.UpdateTemplate(tx, newTemplate, existingTemplate)
	if err != nil {
		return nil, err
	}

	return mappers.TemplateEntityToModel(updatedTemplate), nil
}

func (r *mutationResolver) DeleteTemplate(ctx context.Context, templateID string) (*models.Template, error) {
	var err error
	tx, commitOrRollback := utils.StartTransaction2(r.DB(ctx), &err)
	defer commitOrRollback()
	err = repos.DeleteTemplate(tx, templateID)
	if err != nil {
		return nil, err
	}

	return templateByID(tx.Unscoped(), templateID)
}

// Field Resolvers

type templateResolver struct{ *Resolver }

func (r *templateResolver) CreatedBy(ctx context.Context, obj *models.Template) (*models.User, error) {
	return userByID(r.DB(ctx), obj.CreatedByUserID)
}

func (r *templateResolver) UpdatedBy(ctx context.Context, obj *models.Template) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UpdatedByUserID)
}

func (r *templateResolver) DeletedBy(ctx context.Context, obj *models.Template) (*models.User, error) {
	return deletedUserByID(r.DB(ctx), obj.DeletedByUserID)
}

func (r *templateResolver) Show(ctx context.Context, obj *models.Template) (*models.Show, error) {
	return showByID(r.DB(ctx), obj.ShowID)
}

func (r *templateResolver) SourceEpisode(ctx context.Context, obj *models.Template) (*models.Episode, error) {
	return episodeByID(r.DB(ctx), obj.SourceEpisodeID)
}

func (r *templateResolver) Timestamps(ctx context.Context, obj *models.Template) ([]*models.Timestamp, error) {
	return timestampsByTemplateId(r.DB(ctx), obj.ID)
}

func (r *templateResolver) TimestampIds(ctx context.Context, obj *models.Template) ([]string, error) {
	return timestampIDsByTemplateId(r.DB(ctx), obj.ID)
}

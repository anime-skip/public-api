package resolvers

import (
	"context"
	"sort"

	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils"
	"github.com/jinzhu/gorm"
)

// Helpers

func templateTimestampByIDs(db *gorm.DB, templateID string, timestampID string) (*models.TemplateTimestamp, error) {
	templateTimestamp, err := repos.FindTemplateTimestampByIDs(db, templateID, timestampID)
	if err != nil {
		return nil, err
	}

	return mappers.TemplateTimestampEntityToModel(templateTimestamp), nil
}

func timestampIDsByTemplateId(db *gorm.DB, templateID string) ([]string, error) {
	templateTimestamps, err := repos.FindTemplateTimestampsByTemplateID(db, templateID)
	if err != nil {
		return nil, err
	}

	templateTimestampModels := make([]string, len(templateTimestamps))
	for index, templateTimestamp := range templateTimestamps {
		templateTimestampModels[index] = templateTimestamp.TimestampID.String()
	}
	return templateTimestampModels, nil
}

func timestampsByTemplateId(db *gorm.DB, templateID string) ([]*models.Timestamp, error) {
	templateTimestamps, err := repos.FindTemplateTimestampsByTemplateID(db, templateID)
	if err != nil {
		return nil, err
	}

	timestamps := make([]*models.Timestamp, len(templateTimestamps))
	for index, templateTimestamp := range templateTimestamps {
		timestamps[index], err = timestampByID(db, templateTimestamp.TimestampID.String())
		if err != nil {
			return nil, err
		}
	}

	sort.Slice(timestamps, func(i, j int) bool {
		return timestamps[i].At < timestamps[j].At
	})

	return timestamps, nil
}

// Mutation Resolvers

func (r *mutationResolver) AddTimestampToTemplate(ctx context.Context, templateTimestamp models.InputTemplateTimestamp) (*models.TemplateTimestamp, error) {
	templateTimestampEntity, err := repos.CreateTemplateTimestamp(r.DB(ctx), templateTimestamp)
	if err != nil {
		return nil, err
	}
	return mappers.TemplateTimestampEntityToModel(templateTimestampEntity), nil
}

func (r *mutationResolver) RemoveTimestampFromTemplate(ctx context.Context, templateTimestamp models.InputTemplateTimestamp) (*models.TemplateTimestamp, error) {
	var err error
	tx, commitOrRollback := utils.StartTransaction2(r.DB(ctx), &err)
	defer commitOrRollback()

	err = repos.DeleteTemplateTimestamp(tx, templateTimestamp.TemplateID, templateTimestamp.TimestampID)
	if err != nil {
		return nil, err
	}

	return &models.TemplateTimestamp{
		TemplateID:  templateTimestamp.TemplateID,
		TimestampID: templateTimestamp.TimestampID,
	}, nil
}

// func (r *mutationResolver) UpdateTemplateTimestamps(ctx context.Context, templateID string, add []*models.InputTemplateTimestamp, remove []*models.InputTemplateTimestamp) ([]*models.TemplateTimestamp, error) {
// 	panic("TODO")
// }

// Field Resolvers

type templateTimestampResolver struct{ *Resolver }

func (r *templateTimestampResolver) Template(ctx context.Context, obj *models.TemplateTimestamp) (*models.Template, error) {
	return templateByID(r.DB(ctx), obj.TemplateID)
}

func (r *templateTimestampResolver) Timestamp(ctx context.Context, obj *models.TemplateTimestamp) (*models.Timestamp, error) {
	return timestampByID(r.DB(ctx), obj.TimestampID)
}

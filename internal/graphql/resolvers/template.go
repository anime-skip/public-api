package resolvers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/mappers"
	"anime-skip.com/public-api/internal/utils"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getTemplateByID(ctx context.Context, id *uuid.UUID) (*internal.Template, error) {
	if id == nil {
		return nil, nil
	}
	template, err := r.TemplateService.Get(ctx, internal.TemplatesFilter{
		ID:             id,
		IncludeDeleted: true,
	})
	if err != nil {
		return nil, err
	}
	return &template, nil
}

func (r *Resolver) getTemplateByEpisodeID(ctx context.Context, episodeID *uuid.UUID) (*internal.Template, error) {
	template, err := r.TemplateService.Get(ctx, internal.TemplatesFilter{
		SourceEpisodeID: episodeID,
	})
	if err != nil {
		return nil, err
	}
	return &template, nil
}

func (r *Resolver) getTemplatesByShowID(ctx context.Context, showID *uuid.UUID) ([]*internal.Template, error) {
	templates, err := r.TemplateService.List(ctx, internal.TemplatesFilter{
		ShowID: showID,
	})
	if err != nil {
		return nil, err
	}
	return utils.PtrSlice(templates), nil
}

// Mutations

func (r *mutationResolver) CreateTemplate(ctx context.Context, input internal.InputTemplate) (*internal.Template, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	newTemplate := internal.Template{}
	mappers.ApplyGraphqlInputTemplate(input, &newTemplate)

	created, err := r.TemplateService.Create(ctx, newTemplate, auth.UserID)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *mutationResolver) UpdateTemplate(ctx context.Context, templateID *uuid.UUID, newTemplate internal.InputTemplate) (*internal.Template, error) {
	log.V("Updating: %v", templateID)
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	existing, err := r.TemplateService.Get(ctx, internal.TemplatesFilter{
		ID: templateID,
	})
	if err != nil {
		return nil, err
	}
	mappers.ApplyGraphqlInputTemplate(newTemplate, &existing)
	log.V("Updating to %+v", existing)
	template, err := r.TemplateService.Update(ctx, existing, auth.UserID)
	if err != nil {
		log.V("Failed to update: %v", err)
		return nil, err
	}

	return &template, nil
}

func (r *mutationResolver) DeleteTemplate(ctx context.Context, templateID *uuid.UUID) (*internal.Template, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	deleted, err := r.TemplateService.Delete(ctx, *templateID, auth.UserID)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

// Queries

func (r *queryResolver) FindTemplate(ctx context.Context, templateID *uuid.UUID) (*internal.Template, error) {
	return r.getTemplateByID(ctx, templateID)
}

func (r *queryResolver) FindTemplatesByShowID(ctx context.Context, showID *uuid.UUID) ([]*internal.Template, error) {
	return r.getTemplatesByShowID(ctx, showID)
}

func (r *queryResolver) FindTemplateByDetails(ctx context.Context, episodeID *uuid.UUID, showName *string, season *string) (*internal.Template, error) {
	templates := []internal.Template{}

	// 1. Matching source episodeId
	if episodeID != nil {
		templates, err := r.TemplateService.List(ctx, internal.TemplatesFilter{
			SourceEpisodeID: episodeID,
		})
		if err != nil {
			return nil, err
		} else if len(templates) > 0 {
			return &templates[0], nil
		}
	}

	if showName != nil {
		show, err := r.ShowService.Get(ctx, internal.ShowsFilter{
			Name: showName,
		})
		if err != nil {
			return nil, err
		}

		// 2. Matching showname (case sensitive) and season (case sensitive)
		if season != nil {
			templates, err = r.TemplateService.List(ctx, internal.TemplatesFilter{
				ShowID: show.ID,
				Season: season,
				Type:   utils.Ptr(internal.TEMPLATE_TYPE_SEASONS),
			})
			if err != nil {
				return nil, err
			} else if len(templates) > 0 {
				return &templates[0], nil
			}
		}

		// 3. Matching showname (case sensitive)
		templates, err = r.TemplateService.List(ctx, internal.TemplatesFilter{
			ShowID: show.ID,
			Type:   utils.Ptr(internal.TEMPLATE_TYPE_SHOW),
		})
		if err != nil {
			return nil, err
		} else if len(templates) > 0 {
			return &templates[0], nil
		}
	}

	return nil, internal.NewNotFound("Template", "FindTemplateByDetails")
}

// Fields

func (r *templateResolver) CreatedBy(ctx context.Context, obj *internal.Template) (*internal.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *templateResolver) UpdatedBy(ctx context.Context, obj *internal.Template) (*internal.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *templateResolver) DeletedBy(ctx context.Context, obj *internal.Template) (*internal.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}

func (r *templateResolver) Show(ctx context.Context, obj *internal.Template) (*internal.Show, error) {
	return r.getShowById(ctx, obj.ShowID)
}

func (r *templateResolver) SourceEpisode(ctx context.Context, obj *internal.Template) (*internal.Episode, error) {
	return r.getEpisodeByID(ctx, obj.SourceEpisodeID)
}

func (r *templateResolver) Timestamps(ctx context.Context, obj *internal.Template) ([]*internal.Timestamp, error) {
	templateTimestamps, err := r.TemplateTimestampService.List(ctx, internal.TemplateTimestampsFilter{
		TemplateID: obj.ID,
	})
	if err != nil {
		return nil, err
	}

	timestamps := []internal.Timestamp{}
	for _, templateTimestamp := range templateTimestamps {
		timestamp, err := r.TimestampService.Get(ctx, internal.TimestampsFilter{
			ID: templateTimestamp.TimestampID,
		})
		if err != nil {
			return nil, err
		}
		timestamps = append(timestamps, timestamp)
	}
	return utils.PtrSlice(timestamps), nil
}

func (r *templateResolver) TimestampIds(ctx context.Context, obj *internal.Template) ([]*uuid.UUID, error) {
	templateTimestamps, err := r.TemplateTimestampService.List(ctx, internal.TemplateTimestampsFilter{
		TemplateID: obj.ID,
	})
	if err != nil {
		return nil, err
	}

	ids := []*uuid.UUID{}
	for _, timestamp := range templateTimestamps {
		ids = append(ids, timestamp.TimestampID)
	}
	return ids, nil
}

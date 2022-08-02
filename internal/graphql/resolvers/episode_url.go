package resolvers

import (
	"regexp"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/mappers"
	"anime-skip.com/public-api/internal/utils"
	"github.com/gofrs/uuid"
)

// Helpers

// cleanEpisodeURL takes in a URL, strips any unnecessary query params, transforms the domain, and returns a cleaned version of the URL
func cleanEpisodeURL(url string) string {
	nineAnimeRegexp := regexp.MustCompile("9anime\\.\\w+")
	url = nineAnimeRegexp.ReplaceAllString(url, "9anime.to")
	return url
}

func (r *Resolver) getEpisodeURLByURL(ctx context.Context, url string) (*internal.EpisodeURL, error) {
	url = cleanEpisodeURL(url)
	episodeURL, err := r.EpisodeURLService.Get(ctx, internal.EpisodeURLsFilter{
		URL: &url,
	})
	if err != nil {
		return nil, err
	}
	return &episodeURL, nil
}

func (r *Resolver) getEpisodeURLsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*internal.EpisodeURL, error) {
	episodeURLs, err := r.EpisodeURLService.List(ctx, internal.EpisodeURLsFilter{
		EpisodeID: episodeID,
	})
	if err != nil {
		return nil, err
	}
	return utils.PtrSlice(episodeURLs), nil
}

// Mutations

func (r *mutationResolver) CreateEpisodeURL(ctx context.Context, episodeID *uuid.UUID, episodeURLInput internal.InputEpisodeURL) (*internal.EpisodeURL, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	newEpisodeURL := internal.EpisodeURL{
		EpisodeID: episodeID,
	}
	episodeURLInput.URL = cleanEpisodeURL(episodeURLInput.URL)
	mappers.ApplyGraphqlInputEpisodeURL(episodeURLInput, &newEpisodeURL)

	created, err := r.EpisodeURLService.Create(ctx, newEpisodeURL, auth.UserID)
	if err != nil {
		return nil, err
	}
	return &created, nil
}

func (r *mutationResolver) DeleteEpisodeURL(ctx context.Context, episodeURL string) (*internal.EpisodeURL, error) {
	episodeURL = cleanEpisodeURL(episodeURL)
	deleted, err := r.EpisodeURLService.Delete(ctx, episodeURL)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

func (r *mutationResolver) UpdateEpisodeURL(ctx context.Context, url string, newEpisodeURL internal.InputEpisodeURL) (*internal.EpisodeURL, error) {
	url = cleanEpisodeURL(url)
	log.V("Updating: %v", url)
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	existing, err := r.EpisodeURLService.Get(ctx, internal.EpisodeURLsFilter{
		URL: &url,
	})
	if err != nil {
		return nil, err
	}
	mappers.ApplyGraphqlInputEpisodeURL(newEpisodeURL, &existing)
	log.V("Updating to %+v", existing)
	updated, err := r.EpisodeURLService.Update(ctx, existing, auth.UserID)
	if err != nil {
		log.V("Failed to update: %v", err)
		return nil, err
	}

	return &updated, nil
}

// Queries

func (r *queryResolver) FindEpisodeURL(ctx context.Context, episodeURL string) (*internal.EpisodeURL, error) {
	return r.getEpisodeURLByURL(ctx, episodeURL)
}

func (r *queryResolver) FindEpisodeUrlsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*internal.EpisodeURL, error) {
	return r.getEpisodeURLsByEpisodeID(ctx, episodeID)
}

// Fields

func (r *episodeUrlResolver) CreatedBy(ctx context.Context, obj *internal.EpisodeURL) (*internal.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *episodeUrlResolver) UpdatedBy(ctx context.Context, obj *internal.EpisodeURL) (*internal.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *episodeUrlResolver) Episode(ctx context.Context, obj *internal.EpisodeURL) (*internal.Episode, error) {
	return r.getEpisodeByID(ctx, obj.EpisodeID)
}

func (r *episodeUrlResolver) URL(ctx context.Context, obj *internal.EpisodeURL) (string, error) {
	return cleanEpisodeURL(obj.URL), nil
}

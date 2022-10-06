package resolvers

import (
	"net/url"
	"regexp"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/validation"
	"github.com/gofrs/uuid"
	"github.com/samber/lo"
)

// Helpers

func (r *Resolver) getExternalLinksByShowID(ctx context.Context, showID *uuid.UUID) ([]*internal.ExternalLink, error) {
	links, err := r.ExternalLinkService.List(ctx, internal.ExternalLinksFilter{
		ShowID: showID,
	})
	if err != nil {
		return nil, err
	}
	return lo.ToSlicePtr(links), nil
}

// Mutations

func (r *mutationResolver) AddExternalLink(ctx context.Context, showID *uuid.UUID, url string) (*internal.ExternalLink, error) {
	cleanURL, err := validation.SanitizeExternalLinkURL(url)
	if err != nil {
		return nil, err
	}

	newLink := internal.ExternalLink{
		URL:    cleanURL,
		ShowID: showID,
	}
	created, err := r.ExternalLinkService.Create(ctx, newLink)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *mutationResolver) RemoveExternalLink(ctx context.Context, showID *uuid.UUID, url string) (*internal.ExternalLink, error) {
	removed, err := r.ExternalLinkService.Delete(ctx, url, *showID)
	if err != nil {
		return nil, err
	}
	return &removed, nil
}

// Queries

// Fields

func (r *externalLinkResolver) URL(ctx context.Context, obj *internal.ExternalLink) (string, error) {
	return validation.SanitizeExternalLinkURL(obj.URL)
}

func (r *externalLinkResolver) Show(ctx context.Context, obj *internal.ExternalLink) (*internal.Show, error) {
	return r.getShowByID(ctx, obj.ShowID)
}

func (r *externalLinkResolver) Service(ctx context.Context, obj *internal.ExternalLink) (string, error) {
	u, err := url.Parse(obj.URL)
	if err != nil {
		return "", err
	}
	return u.Hostname(), nil
}

func (r *externalLinkResolver) ServiceID(ctx context.Context, obj *internal.ExternalLink) (*string, error) {
	u, err := url.Parse(obj.URL)
	if err != nil {
		return nil, err
	}
	switch u.Hostname() {
	case internal.ExternalServiceAnilist.Hostname():
		re := regexp.MustCompile(`^\/anime\/([0-9]+)\/?.*?$`)
		match := re.FindStringSubmatch(u.Path)
		if len(match) == 2 {
			return lo.ToPtr(match[1]), nil
		}
		return nil, nil
	default:
		return nil, nil
	}
}

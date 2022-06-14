package resolvers

import (
	"net/url"
	"regexp"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/utils"
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

// Queries

// Fields

func (r *externalLinkResolver) URL(ctx context.Context, obj *internal.ExternalLink) (string, error) {
	return utils.SanitizeUrlString(obj.URL)
}

func (r *externalLinkResolver) Show(ctx context.Context, obj *internal.ExternalLink) (*internal.Show, error) {
	return r.getShowById(ctx, obj.ShowID)
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
		return lo.ToPtr(re.FindStringSubmatch(u.Path)[1]), nil
	default:
		return nil, nil
	}
}

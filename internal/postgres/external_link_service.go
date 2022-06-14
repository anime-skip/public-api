package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"github.com/gofrs/uuid"
)

type externalLinkService struct {
	db internal.Database
}

func NewExternalLinkService(db internal.Database) internal.ExternalLinkService {
	return &externalLinkService{db}
}

func (s *externalLinkService) Get(ctx context.Context, filter internal.ExternalLinksFilter) (internal.ExternalLink, error) {
	return inTx(ctx, s.db, false, internal.ZeroExternalLink, func(tx internal.Tx) (internal.ExternalLink, error) {
		return findExternalLink(ctx, tx, filter)
	})
}

func (s *externalLinkService) List(ctx context.Context, filter internal.ExternalLinksFilter) ([]internal.ExternalLink, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.ExternalLink, error) {
		return findExternalLinks(ctx, tx, filter)
	})
}

func (s *externalLinkService) Create(ctx context.Context, newExternalLink internal.ExternalLink) (internal.ExternalLink, error) {
	return inTx(ctx, s.db, true, internal.ZeroExternalLink, func(tx internal.Tx) (internal.ExternalLink, error) {
		return createExternalLink(ctx, tx, newExternalLink)
	})
}

func (s *externalLinkService) Delete(ctx context.Context, url string, showID uuid.UUID) (internal.ExternalLink, error) {
	return inTx(ctx, s.db, true, internal.ZeroExternalLink, func(tx internal.Tx) (internal.ExternalLink, error) {
		existing, err := findExternalLink(ctx, tx, internal.ExternalLinksFilter{
			URL: &url,
		})
		if err != nil {
			return internal.ZeroExternalLink, err
		}
		return deleteCascadeExternalLink(ctx, tx, existing)
	})
}

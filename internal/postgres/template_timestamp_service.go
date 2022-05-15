package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
)

type templateTimestampService struct {
	db internal.Database
}

func NewTemplateTimestampService(db internal.Database) internal.TemplateTimestampService {
	return &templateTimestampService{db}
}

func (s *templateTimestampService) Get(ctx context.Context, filter internal.TemplateTimestampsFilter) (internal.TemplateTimestamp, error) {
	return inTx(ctx, s.db, false, internal.ZeroTemplateTimestamp, func(tx internal.Tx) (internal.TemplateTimestamp, error) {
		return findTemplateTimestamp(ctx, tx, filter)
	})
}

func (s *templateTimestampService) List(ctx context.Context, filter internal.TemplateTimestampsFilter) ([]internal.TemplateTimestamp, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.TemplateTimestamp, error) {
		return findTemplateTimestamps(ctx, tx, filter)
	})
}

func (s *templateTimestampService) Create(ctx context.Context, newTemplateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	return inTx(ctx, s.db, true, internal.ZeroTemplateTimestamp, func(tx internal.Tx) (internal.TemplateTimestamp, error) {
		return createTemplateTimestamp(ctx, tx, newTemplateTimestamp)
	})
}

func (s *templateTimestampService) Update(ctx context.Context, newTemplateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	return inTx(ctx, s.db, true, internal.ZeroTemplateTimestamp, func(tx internal.Tx) (internal.TemplateTimestamp, error) {
		return updateTemplateTimestamp(ctx, tx, newTemplateTimestamp)
	})
}

func (s *templateTimestampService) Delete(ctx context.Context, templateTimestamp internal.InputTemplateTimestamp) (internal.TemplateTimestamp, error) {
	return inTx(ctx, s.db, true, internal.ZeroTemplateTimestamp, func(tx internal.Tx) (internal.TemplateTimestamp, error) {
		existing, err := findTemplateTimestamp(ctx, tx, internal.TemplateTimestampsFilter{
			TemplateID:  templateTimestamp.TemplateID,
			TimestampID: templateTimestamp.TemplateID,
		})
		if err != nil {
			return internal.ZeroTemplateTimestamp, err
		}
		return deleteCascadeTemplateTimestamp(ctx, tx, existing)
	})
}

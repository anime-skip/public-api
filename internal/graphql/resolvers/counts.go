package resolvers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
)

// Helpers

// Mutations

// Queries

// Counts implements graphql.QueryResolver
func (*queryResolver) Counts(ctx context.Context) (*internal.TotalCounts, error) {
	return &internal.TotalCounts{}, nil
}

// Fields

// EpisodeUrls implements graphql.TotalCountsResolver
func (r *totalCountsResolver) EpisodeUrls(ctx context.Context, obj *internal.TotalCounts) (int, error) {
	return r.EpisodeURLService.Count(ctx)
}

// Episodes implements graphql.TotalCountsResolver
func (r *totalCountsResolver) Episodes(ctx context.Context, obj *internal.TotalCounts) (int, error) {
	return r.EpisodeService.Count(ctx)
}

// Shows implements graphql.TotalCountsResolver
func (r *totalCountsResolver) Shows(ctx context.Context, obj *internal.TotalCounts) (int, error) {
	return r.ShowService.Count(ctx)
}

// Templates implements graphql.TotalCountsResolver
func (r *totalCountsResolver) Templates(ctx context.Context, obj *internal.TotalCounts) (int, error) {
	return r.TemplateService.Count(ctx)
}

// TimestampTypes implements graphql.TotalCountsResolver
func (r *totalCountsResolver) TimestampTypes(ctx context.Context, obj *internal.TotalCounts) (int, error) {
	return r.TimestampTypeService.Count(ctx)
}

// Timestmaps implements graphql.TotalCountsResolver
func (r *totalCountsResolver) Timestamps(ctx context.Context, obj *internal.TotalCounts) (int, error) {
	return r.TimestampService.Count(ctx)
}

// Users implements graphql.TotalCountsResolver
func (r *totalCountsResolver) Users(ctx context.Context, obj *internal.TotalCounts) (int, error) {
	return r.UserService.Count(ctx)
}

package resolvers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/utils"
	"anime-skip.com/public-api/internal/validation"
	"github.com/gofrs/uuid"
	"github.com/samber/lo"
)

// Mutations

// CreateUserReport implements graphql.MutationResolver
func (r *mutationResolver) CreateUserReport(ctx context.Context, inputReport *internal.InputUserReport) (*internal.UserReport, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	report, err := validation.InputUserReport(*inputReport)
	if err != nil {
		return nil, err
	}

	createdReport, err := r.UserReportService.Create(
		ctx,
		internal.UserReport{
			Message:          report.Message,
			ReportedFromURL:  report.ReportedFromURL,
			TimestampID:      report.TimestampID,
			EpisodeID:        report.EpisodeID,
			EpisodeURLString: report.EpisodeURL,
			ShowID:           report.ShowID,
		},
		auth.UserID,
	)
	if err != nil {
		return nil, err
	}
	return &createdReport, nil
}

// ResolveUserReport implements graphql.MutationResolver
func (r *mutationResolver) ResolveUserReport(ctx context.Context, id *uuid.UUID, resolvedMessage *string) (*internal.UserReport, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}
	updatedReport, err := r.UserReportService.Resolve(ctx, *id, resolvedMessage, auth.UserID)
	if err != nil {
		return nil, err
	}
	return &updatedReport, nil
}

// Queries

// FindUserReport implements graphql.QueryResolver
func (r *queryResolver) FindUserReport(ctx context.Context, id *uuid.UUID) (*internal.UserReport, error) {
	filter := internal.UserReportsFilter{
		ID:             id,
		IncludeDeleted: true,
	}
	report, err := r.UserReportService.Get(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

// FindUserReports implements graphql.QueryResolver
func (r *queryResolver) FindUserReports(ctx context.Context, resolved *bool, offset *int, limit *int, sort *string) ([]*internal.UserReport, error) {
	filter := internal.UserReportsFilter{
		Pagination: &internal.Pagination{
			Offset: utils.ValueOr(offset, 0),
			Limit:  utils.ValueOr(limit, 10),
		},
		Sort:     utils.ValueOr(sort, "DESC"),
		Resolved: resolved,
	}
	reports, err := r.UserReportService.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return lo.ToSlicePtr(reports), nil
}

// Fields

func (r *userReportResolver) CreatedBy(ctx context.Context, obj *internal.UserReport) (*internal.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *userReportResolver) UpdatedBy(ctx context.Context, obj *internal.UserReport) (*internal.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *userReportResolver) DeletedBy(ctx context.Context, obj *internal.UserReport) (*internal.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}

func (r *userReportResolver) Timestamp(ctx context.Context, obj *internal.UserReport) (*internal.Timestamp, error) {
	return r.getTimestampByID(ctx, obj.TimestampID)
}

func (r *userReportResolver) Episode(ctx context.Context, obj *internal.UserReport) (*internal.Episode, error) {
	return r.getEpisodeByID(ctx, obj.EpisodeID)
}

func (r *userReportResolver) EpisodeURL(ctx context.Context, obj *internal.UserReport) (*internal.EpisodeURL, error) {
	if obj.EpisodeURLString == nil {
		return nil, nil
	}
	return r.getEpisodeURLByURL(ctx, *obj.EpisodeURLString)
}

func (r *userReportResolver) Show(ctx context.Context, obj *internal.UserReport) (*internal.Show, error) {
	return r.getShowByID(ctx, obj.ShowID)
}

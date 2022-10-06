package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"github.com/gofrs/uuid"
)

type userReportService struct {
	db internal.Database
}

func NewUserReportService(db internal.Database) internal.UserReportService {
	return &userReportService{db}
}

func (s *userReportService) Get(ctx context.Context, filter internal.UserReportsFilter) (internal.UserReport, error) {
	return inTx(ctx, s.db, false, internal.ZeroUserReport, func(tx internal.Tx) (internal.UserReport, error) {
		return findUserReport(ctx, tx, filter)
	})
}

func (s *userReportService) List(ctx context.Context, filter internal.UserReportsFilter) ([]internal.UserReport, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.UserReport, error) {
		return findUserReports(ctx, tx, filter)
	})
}

func (s *userReportService) Create(ctx context.Context, newReport internal.UserReport, createdBy uuid.UUID) (internal.UserReport, error) {
	return inTx(ctx, s.db, true, internal.ZeroUserReport, func(tx internal.Tx) (internal.UserReport, error) {
		return createUserReport(ctx, tx, newReport, createdBy)
	})
}

func (s *userReportService) Update(ctx context.Context, newReport internal.UserReport, updatedBy uuid.UUID) (internal.UserReport, error) {
	return inTx(ctx, s.db, true, internal.ZeroUserReport, func(tx internal.Tx) (internal.UserReport, error) {
		return updateUserReport(ctx, tx, newReport, updatedBy)
	})
}

func (s *userReportService) Resolve(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (internal.UserReport, error) {
	return inTx(ctx, s.db, true, internal.ZeroUserReport, func(tx internal.Tx) (internal.UserReport, error) {
		existing, err := findUserReport(ctx, tx, internal.UserReportsFilter{
			ID: &id,
		})
		if err != nil {
			return internal.ZeroUserReport, err
		}
		existing.Resolved = true
		return deleteCascadeUserReport(ctx, tx, existing, deletedBy)
	})
}

func (s *userReportService) Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (internal.UserReport, error) {
	return inTx(ctx, s.db, true, internal.ZeroUserReport, func(tx internal.Tx) (internal.UserReport, error) {
		existing, err := findUserReport(ctx, tx, internal.UserReportsFilter{
			ID: &id,
		})
		if err != nil {
			return internal.ZeroUserReport, err
		}
		return deleteCascadeUserReport(ctx, tx, existing, deletedBy)
	})
}

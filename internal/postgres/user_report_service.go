package postgres

import (
	"context"
	"fmt"
	"strings"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"github.com/gofrs/uuid"
)

type userReportService struct {
	db      internal.Database
	alerter internal.Alerter
}

func NewUserReportService(db internal.Database, alerter internal.Alerter) internal.UserReportService {
	return &userReportService{
		db:      db,
		alerter: alerter,
	}
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
		createdReport, err := createUserReport(ctx, tx, newReport, createdBy)
		if err != nil {
			return internal.ZeroUserReport, err
		}
		alertErr := s.sendNewReportAlert(createdReport)
		if alertErr != nil {
			log.E("Failed to send alert for new user report: %v", alertErr)
		} else {
			log.I("Sent alert for new user report: %s", createdReport.ID.String())
		}
		return createdReport, err
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

func (s *userReportService) sendNewReportAlert(report internal.UserReport) error {
	lines := []string{}
	lines = append(lines, fmt.Sprintf("New User Report `%s`", report.ID.String()))
	for _, messageLine := range strings.Split(report.Message, "\n") {
		lines = append(lines, fmt.Sprintf("> %s", messageLine))
	}
	lines = append(lines, fmt.Sprintf("> %s", report.ReportedFromURL))
	return s.alerter.Send(strings.Join(lines, "\n"))
}

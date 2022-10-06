package validation

import (
	"strings"

	"anime-skip.com/public-api/internal"
)

func InputUserReport(report internal.InputUserReport) (internal.InputUserReport, error) {
	// Cleanup
	report.Message = strings.TrimSpace(report.Message)
	report.ReportedFromURL = strings.TrimSpace(report.ReportedFromURL)
	if report.EpisodeURL != nil {
		trimmedUrl := strings.TrimSpace(*report.EpisodeURL)
		if trimmedUrl == "" {
			report.EpisodeURL = nil
		} else {
			*report.EpisodeURL = trimmedUrl
		}
	}

	// Validate
	ids := 0
	if report.TimestampID != nil {
		ids++
	}
	if report.EpisodeID != nil {
		ids++
	}
	if report.EpisodeURL != nil {
		ids++
	}
	if report.ShowID != nil {
		ids++
	}
	if ids == 0 {
		return report, &internal.Error{
			Code:    internal.EINVALID,
			Message: "You must provide one of timestampId, episodeId, episodeUrl, or showId",
			Op:      "validate.InputUserReport",
		}
	}
	if ids > 1 {
		return report, &internal.Error{
			Code:    internal.EINVALID,
			Message: "You must provide ONLY one of timestampId, episodeId, episodeUrl, or showId",
			Op:      "validate.InputUserReport",
		}
	}
	if report.ReportedFromURL == "" {
		return report, &internal.Error{
			Code:    internal.EINVALID,
			Message: "reportedFromUrl cannot be empty",
			Op:      "validate.InputUserReport",
		}
	}
	if report.Message == "" {
		return report, &internal.Error{
			Code:    internal.EINVALID,
			Message: "message cannot be empty",
			Op:      "validate.InputUserReport",
		}
	}
	return report, nil
}

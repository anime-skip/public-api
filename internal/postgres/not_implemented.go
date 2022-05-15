package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	uuid "github.com/gofrs/uuid"
)

// Episode

func findEpisodes(ctx context.Context, tx internal.Tx, filter internal.EpisodesFilter) ([]internal.Episode, error) {
	return []internal.Episode{}, internal.NewNotImplemented("findEpisodes")
}

func findEpisode(ctx context.Context, tx internal.Tx, filter internal.EpisodesFilter) (internal.Episode, error) {
	all, err := findEpisodes(ctx, tx, filter)
	if err != nil {
		return internal.ZeroEpisode, err
	} else if len(all) == 0 {
		return internal.ZeroEpisode, internal.NewNotFound("Episode", "findEpisode")
	}
	return all[0], nil
}

func createEpisode(ctx context.Context, tx internal.Tx, episode internal.Episode, createdBy uuid.UUID) (internal.Episode, error) {
	return episode, internal.NewNotImplemented("createEpisode")
}

func updateEpisode(ctx context.Context, tx internal.Tx, episode internal.Episode, updatedBy uuid.UUID) (internal.Episode, error) {
	return episode, internal.NewNotImplemented("updateEpisode")
}

func deleteEpisode(ctx context.Context, tx internal.Tx, episode internal.Episode, deletedBy uuid.UUID) (internal.Episode, error) {
	return episode, internal.NewNotImplemented("deleteEpisode")
}

// EpisodeURL

func findEpisodeURLs(ctx context.Context, tx internal.Tx, filter internal.EpisodeURLsFilter) ([]internal.EpisodeURL, error) {
	return []internal.EpisodeURL{}, internal.NewNotImplemented("findEpisodeURLs")
}

func findEpisodeURL(ctx context.Context, tx internal.Tx, filter internal.EpisodeURLsFilter) (internal.EpisodeURL, error) {
	all, err := findEpisodeURLs(ctx, tx, filter)
	if err != nil {
		return internal.ZeroEpisodeURL, err
	} else if len(all) == 0 {
		return internal.ZeroEpisodeURL, internal.NewNotFound("Episode URL", "findEpisodeURL")
	}
	return all[0], nil
}

func createEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL, createdBy uuid.UUID) (internal.EpisodeURL, error) {
	return episodeURL, internal.NewNotImplemented("createEpisodeURL")
}

func updateEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL, updatedBy uuid.UUID) (internal.EpisodeURL, error) {
	return episodeURL, internal.NewNotImplemented("updateEpisodeURL")
}

func deleteEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL) (internal.EpisodeURL, error) {
	return episodeURL, internal.NewNotImplemented("deleteEpisodeURL")
}

// Preferences

func findPreferences(ctx context.Context, tx internal.Tx, filter internal.PreferencesFilter) (internal.Preferences, error) {
	return internal.ZeroPreferences, internal.NewNotImplemented("findPreferences")
}

func updatePreferences(ctx context.Context, tx internal.Tx, preferences internal.Preferences) (internal.Preferences, error) {
	return preferences, internal.NewNotImplemented("updatePreferences")
}

func deletePreferences(ctx context.Context, tx internal.Tx, preferences internal.Preferences) (internal.Preferences, error) {
	return preferences, internal.NewNotImplemented("deletePreferences")
}

// ShowAdmin

func findShowAdmins(ctx context.Context, tx internal.Tx, filter internal.ShowAdminsFilter) ([]internal.ShowAdmin, error) {
	return []internal.ShowAdmin{}, internal.NewNotImplemented("findShowAdmins")
}

func findShowAdmin(ctx context.Context, tx internal.Tx, filter internal.ShowAdminsFilter) (internal.ShowAdmin, error) {
	all, err := findShowAdmins(ctx, tx, filter)
	if err != nil {
		return internal.ZeroShowAdmin, err
	} else if len(all) == 0 {
		return internal.ZeroShowAdmin, internal.NewNotFound("Show admin", "findShowAdmin")
	}
	return all[0], nil
}

func createShowAdmin(ctx context.Context, tx internal.Tx, showAdmin internal.ShowAdmin, createdBy uuid.UUID) (internal.ShowAdmin, error) {
	return showAdmin, internal.NewNotImplemented("createShowAdmin")
}

func updateShowAdmin(ctx context.Context, tx internal.Tx, showAdmin internal.ShowAdmin, updatedBy uuid.UUID) (internal.ShowAdmin, error) {
	return showAdmin, internal.NewNotImplemented("updateShowAdmin")
}

func deleteShowAdmin(ctx context.Context, tx internal.Tx, showAdmin internal.ShowAdmin, deletedBy uuid.UUID) (internal.ShowAdmin, error) {
	return showAdmin, internal.NewNotImplemented("deleteShowAdmin")
}

// Show

// func (s *showService) Search(ctx context.Context, filter internal.ShowsFilter) ([]internal.Show, error) {
// 	return
//
// 	where := []WhereCondition{}
// 	if filter.Search != "" {
// 		where = append(where, WhereLike{
// 			value:      "%" + filter.Search + "%",
// 			column:     "name",
// 			ignoreCase: true,
// 		})
// 	}
//
// 	limitOffset := &LimitOffset{
// 		Limit:  filter.Limit,
// 		Offset: filter.Offset,
// 	}
// 	orderBy := &OrderBy{
// 		Column:    "name",
// 		Direction: filter.Sort,
// 	}
//
// 	return searchShows(ctx, s.db, where, orderBy, limitOffset)
// }
func findShows(ctx context.Context, tx internal.Tx, filter internal.ShowsFilter) ([]internal.Show, error) {
	return []internal.Show{}, internal.NewNotImplemented("findShows")
}

func findShow(ctx context.Context, tx internal.Tx, filter internal.ShowsFilter) (internal.Show, error) {
	all, err := findShows(ctx, tx, filter)
	if err != nil {
		return internal.ZeroShow, err
	} else if len(all) == 0 {
		return internal.ZeroShow, internal.NewNotFound("Show", "findShow")
	}
	return all[0], nil
}

func countShowSeasons(ctx context.Context, tx internal.Tx, id uuid.UUID) (int, error) {
	return 0, internal.NewNotImplemented("countShowSeasons")
}

func createShow(ctx context.Context, tx internal.Tx, show internal.Show, createdBy uuid.UUID) (internal.Show, error) {
	return show, internal.NewNotImplemented("createShow")
}

func updateShow(ctx context.Context, tx internal.Tx, show internal.Show, updatedBy uuid.UUID) (internal.Show, error) {
	return show, internal.NewNotImplemented("updateShow")
}

func deleteShow(ctx context.Context, tx internal.Tx, show internal.Show, deletedBy uuid.UUID) (internal.Show, error) {
	return show, internal.NewNotImplemented("deleteShow")
}

// Template

func findTemplates(ctx context.Context, tx internal.Tx, filter internal.TemplatesFilter) ([]internal.Template, error) {
	return []internal.Template{}, internal.NewNotImplemented("findTemplates")
}

func findTemplate(ctx context.Context, tx internal.Tx, filter internal.TemplatesFilter) (internal.Template, error) {
	all, err := findTemplates(ctx, tx, filter)
	if err != nil {
		return internal.ZeroTemplate, err
	} else if len(all) == 0 {
		return internal.ZeroTemplate, internal.NewNotFound("Template", "findTemplate")
	}
	return all[0], nil
}

func createTemplate(ctx context.Context, tx internal.Tx, template internal.Template, createdBy uuid.UUID) (internal.Template, error) {
	return template, internal.NewNotImplemented("createTemplate")
}

func updateTemplate(ctx context.Context, tx internal.Tx, template internal.Template, updatedBy uuid.UUID) (internal.Template, error) {
	return template, internal.NewNotImplemented("updateTemplate")
}

func deleteTemplate(ctx context.Context, tx internal.Tx, template internal.Template, deletedBy uuid.UUID) (internal.Template, error) {
	return template, internal.NewNotImplemented("deleteTemplate")
}

// TemplateTimestamp

func findTemplateTimestamps(ctx context.Context, tx internal.Tx, filter internal.TemplateTimestampsFilter) ([]internal.TemplateTimestamp, error) {
	return []internal.TemplateTimestamp{}, internal.NewNotImplemented("findTemplateTimestamps")
}

func findTemplateTimestamp(ctx context.Context, tx internal.Tx, filter internal.TemplateTimestampsFilter) (internal.TemplateTimestamp, error) {
	all, err := findTemplateTimestamps(ctx, tx, filter)
	if err != nil {
		return internal.ZeroTemplateTimestamp, err
	} else if len(all) == 0 {
		return internal.ZeroTemplateTimestamp, internal.NewNotFound("Template timestamp", "findTemplateTimestamp")
	}
	return all[0], nil
}

func createTemplateTimestamp(ctx context.Context, tx internal.Tx, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	return templateTimestamp, internal.NewNotImplemented("createTemplateTimestamp")
}

func updateTemplateTimestamp(ctx context.Context, tx internal.Tx, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	return templateTimestamp, internal.NewNotImplemented("updateTemplateTimestamp")
}

func deleteTemplateTimestamp(ctx context.Context, tx internal.Tx, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	return templateTimestamp, internal.NewNotImplemented("deleteTemplateTimestamp")
}

// Timestamp

func findTimestamps(ctx context.Context, tx internal.Tx, filter internal.TimestampsFilter) ([]internal.Timestamp, error) {
	return []internal.Timestamp{}, internal.NewNotImplemented("findTimestamps")
}

func findTimestamp(ctx context.Context, tx internal.Tx, filter internal.TimestampsFilter) (internal.Timestamp, error) {
	all, err := findTimestamps(ctx, tx, filter)
	if err != nil {
		return internal.ZeroTimestamp, err
	} else if len(all) == 0 {
		return internal.ZeroTimestamp, internal.NewNotFound("Timestamp", "findTimestamp")
	}
	return all[0], nil
}

func createTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp, createdBy uuid.UUID) (internal.Timestamp, error) {
	return timestamp, internal.NewNotImplemented("createTimestamp")
}

func updateTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp, updatedBy uuid.UUID) (internal.Timestamp, error) {
	return timestamp, internal.NewNotImplemented("updateTimestamp")
}

func deleteTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp, deletedBy uuid.UUID) (internal.Timestamp, error) {
	return timestamp, internal.NewNotImplemented("deleteTimestamp")
}

// TimestampType

func findTimestampTypes(ctx context.Context, tx internal.Tx, filter internal.TimestampTypesFilter) ([]internal.TimestampType, error) {
	return []internal.TimestampType{}, internal.NewNotImplemented("findTimestampTypes")
}

func findTimestampType(ctx context.Context, tx internal.Tx, filter internal.TimestampTypesFilter) (internal.TimestampType, error) {
	all, err := findTimestampTypes(ctx, tx, filter)
	if err != nil {
		return internal.ZeroTimestampType, err
	} else if len(all) == 0 {
		return internal.ZeroTimestampType, internal.NewNotFound("Timestamp type", "findTimestampType")
	}
	return all[0], nil
}

func createTimestampType(ctx context.Context, tx internal.Tx, timestampType internal.TimestampType, createdBy uuid.UUID) (internal.TimestampType, error) {
	return timestampType, internal.NewNotImplemented("createTimestampType")
}

func updateTimestampType(ctx context.Context, tx internal.Tx, timestampType internal.TimestampType, updatedBy uuid.UUID) (internal.TimestampType, error) {
	return timestampType, internal.NewNotImplemented("updateTimestampType")
}

func deleteTimestampType(ctx context.Context, tx internal.Tx, timestampType internal.TimestampType, deletedBy uuid.UUID) (internal.TimestampType, error) {
	return timestampType, internal.NewNotImplemented("deleteTimestampType")
}

// User

func updateUser(ctx context.Context, tx internal.Tx, user internal.FullUser) (internal.FullUser, error) {
	return user, internal.NewNotImplemented("updateUser")
}

func deleteUser(ctx context.Context, tx internal.Tx, user internal.FullUser) (internal.FullUser, error) {
	return user, internal.NewNotImplemented("deleteUser")
}

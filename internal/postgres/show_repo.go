package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	uuid "github.com/gofrs/uuid"
)

func deleteCascadeShow(ctx context.Context, tx internal.Tx, show internal.Show, deletedBy uuid.UUID) (internal.Show, error) {
	log.V("Deleting show: %v", show.ID)
	deletedShow, err := deleteShow(ctx, tx, show, deletedBy)
	if err != nil {
		return internal.Show{}, err
	}

	log.V("Deleting show admins")
	admins, err := findShowAdmins(ctx, tx, internal.ShowAdminsFilter{
		ShowID: show.ID,
	})
	if err != nil {
		return internal.Show{}, err
	}
	for _, admin := range admins {
		_, err := deleteCascadeShowAdmin(ctx, tx, admin, deletedBy)
		if err != nil {
			return internal.Show{}, err
		}
	}

	log.V("Deleting show templates")
	templates, err := findTemplates(ctx, tx, internal.TemplatesFilter{
		ShowID: show.ID,
	})
	if err != nil {
		return internal.Show{}, err
	}
	for _, template := range templates {
		_, err := deleteCascadeTemplate(ctx, tx, template, deletedBy)
		if err != nil {
			return internal.Show{}, err
		}
	}

	log.V("Deleting show episodes")
	episodes, err := findEpisodes(ctx, tx, internal.EpisodesFilter{
		ShowID: show.ID,
	})
	if err != nil {
		return internal.Show{}, err
	}
	for _, episode := range episodes {
		_, err := deleteCascadeEpisode(ctx, tx, episode, deletedBy)
		if err != nil {
			return internal.Show{}, err
		}
	}

	log.V("Done deleting show: %v", show.ID)
	return deletedShow, nil
}

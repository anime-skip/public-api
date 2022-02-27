package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/log"
)

func deleteCascadeShow(ctx context.Context, tx internal.Tx, show internal.Show) (internal.Show, error) {
	log.V("Deleting show: %v", show.ID)
	deletedShow, err := deleteShowInTx(ctx, tx, show)
	if err != nil {
		return internal.Show{}, err
	}

	log.V("Deleting show admins")
	admins, err := getShowAdminsByShowIDInTx(ctx, tx, show.ID)
	if err != nil {
		return internal.Show{}, err
	}
	for _, admin := range admins {
		_, err := deleteCascadeShowAdmin(ctx, tx, admin)
		if err != nil {
			return internal.Show{}, err
		}
	}

	log.V("Deleting show templates")
	templates, err := getTemplatesByShowIDInTx(ctx, tx, show.ID)
	if err != nil {
		return internal.Show{}, err
	}
	for _, template := range templates {
		_, err := deleteCascadeTemplate(ctx, tx, template)
		if err != nil {
			return internal.Show{}, err
		}
	}

	log.V("Deleting show episodes")
	episodes, err := getEpisodesByShowIDInTx(ctx, tx, show.ID)
	if err != nil {
		return internal.Show{}, err
	}
	for _, episode := range episodes {
		_, err := deleteCascadeEpisode(ctx, tx, episode)
		if err != nil {
			return internal.Show{}, err
		}
	}

	log.V("Done deleting show: %v", show.ID)
	return deletedShow, nil
}

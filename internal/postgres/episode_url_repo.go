package postgres

import (
	"context"

	internal "anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/log"
)

func deleteCascadeEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL) (internal.EpisodeURL, error) {
	log.V("Deleting episode url (nothing to cascade): %s", episodeURL.URL)
	return episodeURL, deleteEpisodeURLInTx(ctx, tx, episodeURL)
}
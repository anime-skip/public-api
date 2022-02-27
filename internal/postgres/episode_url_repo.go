package postgres

import (
	"context"

	internal "anime-skip.com/timestamps-service/internal"
)

func deleteCascadeEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL) (internal.EpisodeURL, error) {
	return episodeURL, deleteEpisodeURLInTx(ctx, tx, episodeURL)
}

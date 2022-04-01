package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
)

func deleteCascadeAPIClient(ctx context.Context, tx internal.Tx, apiClient internal.APIClient) (internal.APIClient, error) {
	log.V("Deleting api client: %v", apiClient.ID)
	return deleteAPIClientInTx(ctx, tx, apiClient)
}

package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/log"
)

func deleteCascadeAPIClient(ctx context.Context, tx internal.Tx, apiClient internal.APIClient) (internal.APIClient, error) {
	log.V("Deleting api client: %v", apiClient.ID)
	return deleteAPIClientInTx(ctx, tx, apiClient)
}

package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
)

func ToGraphqlTimestampType(timestampType internal.TimestampType) graphql.TimestampType {
	return graphql.TimestampType{
		ID:              &timestampType.ID,
		CreatedAt:       timestampType.CreatedAt,
		CreatedByUserID: &timestampType.CreatedByUserID,
		UpdatedAt:       timestampType.UpdatedAt,
		UpdatedByUserID: &timestampType.UpdatedByUserID,
		DeletedAt:       timestampType.DeletedAt,
		DeletedByUserID: timestampType.DeletedByUserID,

		Name:        timestampType.Name,
		Description: timestampType.Description,
	}
}

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

func toGraphqlTimestampTypePointer(timestamp internal.TimestampType) *graphql.TimestampType {
	value := ToGraphqlTimestampType(timestamp)
	return &value
}

func ToGraphqlTimestampTypePointers(timestampTypes []internal.TimestampType) []*graphql.TimestampType {
	result := []*graphql.TimestampType{}
	for _, timestampType := range timestampTypes {
		result = append(result, toGraphqlTimestampTypePointer(timestampType))
	}
	return result
}

package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
)

func ToGraphqlTimestamp(timestamp internal.Timestamp) graphql.Timestamp {
	return graphql.Timestamp{
		ID:              &timestamp.ID,
		CreatedAt:       timestamp.CreatedAt,
		CreatedByUserID: &timestamp.CreatedByUserID,
		UpdatedAt:       timestamp.UpdatedAt,
		UpdatedByUserID: &timestamp.UpdatedByUserID,
		DeletedAt:       timestamp.DeletedAt,
		DeletedByUserID: timestamp.DeletedByUserID,

		At:        timestamp.At,
		Source:    ToTimestampSourceEnum(timestamp.Source),
		TypeID:    &timestamp.TypeID,
		EpisodeID: &timestamp.EpisodeID,
	}
}

func toGraphqlTimestampPointer(timestamp internal.Timestamp) *graphql.Timestamp {
	value := ToGraphqlTimestamp(timestamp)
	return &value
}

func ToGraphqlTimestampPointers(timestamps []internal.Timestamp) []*graphql.Timestamp {
	result := []*graphql.Timestamp{}
	for _, timestamp := range timestamps {
		result = append(result, toGraphqlTimestampPointer(timestamp))
	}
	return result
}

func ApplyGraphqlInputTimestamp(input graphql.InputTimestamp, output *internal.Timestamp) {
	output.At = input.At
	output.TypeID = *input.TypeID
	output.Source = ToTimestampSourceInt(input.Source)
}

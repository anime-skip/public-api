package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
)

func ToGraphqlShow(entity internal.Show) graphql.Show {
	return graphql.Show{
		ID:              &entity.ID,
		CreatedAt:       entity.CreatedAt,
		CreatedByUserID: &entity.CreatedByUserID,
		UpdatedAt:       entity.UpdatedAt,
		UpdatedByUserID: &entity.UpdatedByUserID,
		DeletedAt:       entity.DeletedAt,
		DeletedByUserID: entity.DeletedByUserID,

		Name:         entity.Name,
		OriginalName: entity.OriginalName,
		Website:      entity.Website,
		Image:        entity.Image,
	}
}
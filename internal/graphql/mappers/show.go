package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
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

func ApplyGraphqlInputShow(input graphql.InputShow, output *internal.Show) {
	output.Name = input.Name
	output.OriginalName = input.OriginalName
	output.Website = input.Website
	output.Image = input.Image
}

package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
)

func ToGraphqlShowAdmin(admin internal.ShowAdmin) graphql.ShowAdmin {
	return graphql.ShowAdmin{
		ID:              &admin.ID,
		CreatedAt:       admin.CreatedAt,
		CreatedByUserID: &admin.CreatedByUserID,
		UpdatedAt:       admin.UpdatedAt,
		UpdatedByUserID: &admin.UpdatedByUserID,
		DeletedAt:       admin.DeletedAt,
		DeletedByUserID: admin.DeletedByUserID,

		ShowID: &admin.ShowID,
		UserID: &admin.UserID,
	}
}

func toGraphqlShowAdminPointer(timestamp internal.ShowAdmin) *graphql.ShowAdmin {
	value := ToGraphqlShowAdmin(timestamp)
	return &value
}

func ToGraphqlShowAdminPointers(showAdmins []internal.ShowAdmin) []*graphql.ShowAdmin {
	result := []*graphql.ShowAdmin{}
	for _, showAdmin := range showAdmins {
		result = append(result, toGraphqlShowAdminPointer(showAdmin))
	}
	return result
}

func ApplyGraphqlInputShowAdmin(input graphql.InputShowAdmin, output *internal.ShowAdmin) {
	output.ShowID = *input.ShowID
	output.UserID = *input.UserID
}

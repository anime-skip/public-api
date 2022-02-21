package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
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

func ToGraphqlShowAdmins(admins []internal.ShowAdmin) []graphql.ShowAdmin {
	result := []graphql.ShowAdmin{}
	for _, admin := range admins {
		result = append(result, ToGraphqlShowAdmin(admin))
	}
	return result
}

func ToGraphqlShowAdminPointers(admins []internal.ShowAdmin) []*graphql.ShowAdmin {
	result := []*graphql.ShowAdmin{}
	for _, admin := range ToGraphqlShowAdmins(admins) {
		result = append(result, &admin)
	}
	return result
}

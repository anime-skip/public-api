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

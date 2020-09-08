package directives

import (
	"context"
	"fmt"

	"anime-skip.com/backend/internal/database"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/constants"
	"github.com/99designs/gqlgen/graphql"
)

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role models.Role) (interface{}, error) {
	if err := isAuthorized(ctx); err != nil {
		return nil, err
	}

	roleInt := mappers.RoleEnumToInt(role)
	userID, err := utils.UserIDFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("500 Internal Error [005]")
	}
	user, err := repos.FindUserByID(database.ORMInstance.DB, userID)
	if err != nil {
		return nil, err
	}
	if user.Role == roleInt || user.Role == constants.ROLE_DEV {
		return next(ctx)
	}
	return nil, fmt.Errorf("403 Forebidden")
}

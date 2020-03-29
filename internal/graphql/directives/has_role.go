package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
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

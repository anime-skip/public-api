package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
)

func _findShowID(ctx context.Context, obj interface{}) (string, error) {
	args, ok := obj.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("[%+v] must be a map, but was a %T", obj, obj)
	}

	// showId
	if showID, ok := args["showId"]; ok {
		showIDStr, isString := showID.(string)
		if !isString {
			return "", fmt.Errorf("args['%+v'] must be a string, but was %v (%T)", "showID", showID, showID)
		}
		return showIDStr, nil
	}

	// showAdminId
	if showAdminID, ok := args["showAdminId"]; ok {
		showAdminIDStr, isString := showAdminID.(string)
		if !isString {
			return "", fmt.Errorf("args['%+v'] must be a string, but was %v (%T)", "showAdminID", showAdminID, showAdminID)
		}
		showAdmin, err := repos.FindShowAdminByID(ctx, database.ORMInstance, showAdminIDStr)
		if err != nil {
			return "", err
		}
		return showAdmin.ShowID.String(), nil
	}

	// showAdmin
	if showAdmin, ok := args["showAdmin"]; ok {
		inputShowAdmin, isInputShowAdmin := showAdmin.(*models.InputShowAdmin)
		if isInputShowAdmin {
			return inputShowAdmin.ShowID, nil
		}
		return "", fmt.Errorf("args['%+v'] must be a InputShowAdmin, but was %v (%T)", "showAdmin", showAdmin, showAdmin)
	}

	return "", fmt.Errorf("isShowAdmin directive not implemented for the provided arguments: %+v", obj)
}

func IsShowAdmin(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	log.D("1")
	userID, err := utils.UserIDFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("500 Internal Error [004]")
	}
	log.D("2: userID=%s", userID)
	showID, err := _findShowID(ctx, obj)
	if err != nil {
		return nil, err
	}

	// Basic User that is an admin for the specified show
	log.D("2: showID=%s", showID)
	log.D("orm=%v", database.ORMInstance)
	_, err = repos.FindShowAdminsByUserIDShowID(ctx, database.ORMInstance, userID, showID)
	if err != nil {
		return next(ctx)
	}

	// Admin User
	log.D("4")
	user, err := repos.FindUserByID(ctx, database.ORMInstance, userID)
	if err != nil {
		return nil, err
	}
	if user.Role == constants.ROLE_ADMIN || user.Role == constants.ROLE_DEV {
		return next(ctx)
	}

	log.D("5")
	return nil, fmt.Errorf("403 Forebidden")
}

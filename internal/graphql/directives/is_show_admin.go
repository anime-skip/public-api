package directives

import (
	"context"
	"fmt"

	"anime-skip.com/backend/internal/database"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/constants"
	"github.com/99designs/gqlgen/graphql"
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
			return "", fmt.Errorf("args['%+v'] must be a string, but was %v (%T)", "showId", showID, showID)
		}
		return showIDStr, nil
	}

	// showAdminId
	db := database.ORMInstance.DB
	if showAdminID, ok := args["showAdminId"]; ok {
		showAdminIDStr, isString := showAdminID.(string)
		if !isString {
			return "", fmt.Errorf("args['%+v'] must be a string, but was %v (%T)", "showAdminId", showAdminID, showAdminID)
		}
		showAdmin, err := repos.FindShowAdminByID(db, showAdminIDStr)
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

	// timestampId
	if timestampID, ok := args["timestampId"]; ok {
		timestampIDStr, isString := timestampID.(string)
		if !isString {
			return "", fmt.Errorf("args['%+v'] must be a string, but was %v (%T)", "timestampId", timestampID, timestampID)
		}
		timestamp, err := repos.FindTimestampByID(db, timestampIDStr)
		if err != nil {
			return "", err
		}
		episode, err := repos.FindEpisodeByID(db, timestamp.EpisodeID.String())
		if err != nil {
			return "", err
		}
		return episode.ShowID.String(), nil
	}

	// episodeUrl
	if epiosdeURL, ok := args["episodeUrl"]; ok {
		epiosdeURLStr, isString := epiosdeURL.(string)
		if !isString {
			return "", fmt.Errorf("args['%+v'] must be a string, but was %v (%T)", "episodeUrl", epiosdeURL, epiosdeURL)
		}
		episodeURL, err := repos.FindEpisodeURLByURL(db, epiosdeURLStr)
		if err != nil {
			return "", err
		}
		episode, err := repos.FindEpisodeByID(db, episodeURL.EpisodeID.String())
		if err != nil {
			return "", err
		}
		return episode.ShowID.String(), nil
	}

	// episodeId
	if episodeID, ok := args["episodeId"]; ok {
		episodeIDStr, isString := episodeID.(string)
		if !isString {
			return "", fmt.Errorf("args['%+v'] must be a string, but was %v (%T)", "episodeId", episodeID, episodeID)
		}
		episode, err := repos.FindEpisodeByID(db, episodeIDStr)
		if err != nil {
			return "", err
		}
		return episode.ShowID.String(), nil
	}

	return "", fmt.Errorf("isShowAdmin directive not implemented for the provided arguments: %+v", obj)
}

func IsShowAdmin(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if err := isAuthorized(ctx); err != nil {
		return nil, err
	}

	userID, err := utils.UserIDFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("500 Internal Error [004]")
	}
	showID, err := _findShowID(ctx, obj)
	if err != nil {
		return nil, err
	}

	// Basic User that is an admin for the specified show
	db := database.ORMInstance.DB
	_, err = repos.FindShowAdminsByUserIDShowID(db, userID, showID)
	if err == nil {
		return next(ctx)
	}

	// Basic User that created the show and there are no admins present
	showAdmins, err := repos.FindShowAdminsByShowID(db, showID)
	if err != nil {
		return nil, err
	}
	if len(showAdmins) == 0 {
		show, err := repos.FindShowByID(db, showID)
		if err != nil {
			return nil, err
		}
		if show.CreatedByUserID.String() == userID {
			return next(ctx)
		}
	}

	// Admin User
	user, err := repos.FindUserByID(db, userID)
	if err != nil {
		return nil, err
	}
	if user.Role == constants.ROLE_ADMIN || user.Role == constants.ROLE_DEV {
		return next(ctx)
	}

	return nil, fmt.Errorf("403 Forebidden - you are not a show admin")
}

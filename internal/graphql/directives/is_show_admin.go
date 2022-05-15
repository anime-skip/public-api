package directives

import (
	context2 "context"
	"fmt"
	"strings"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/log"
	graphql2 "github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
)

type showIDGetter = func(ctx context2.Context, s internal.Services, arg any) (uuid.UUID, error)

var showIDGetters = map[string]showIDGetter{
	"showId": func(
		ctx context2.Context, s internal.Services, arg any,
	) (uuid.UUID, error) {
		log.V("@isShowAdmin.showId: (%T) %+v", arg, arg)
		return uuid.FromString(arg.(string))
	},
	"showAdminInput": func(
		ctx context2.Context, s internal.Services, arg any,
	) (uuid.UUID, error) {
		log.V("@isShowAdmin.showAdminInput: (%T) %+v", arg, arg)
		showAdmin := arg.(*internal.InputShowAdmin)
		return *showAdmin.ShowID, nil
	},
	"showAdminId": func(
		ctx context2.Context, s internal.Services, arg any,
	) (uuid.UUID, error) {
		log.V("@isShowAdmin.showAdminId: (%T) %+v", arg, arg)
		showAdminId, err := uuid.FromString(arg.(string))
		if err != nil {
			return uuid.UUID{}, err
		}
		showAdmin, err := s.ShowAdminService.Get(ctx, internal.ShowAdminsFilter{
			ID: &showAdminId,
		})
		if err != nil {
			return uuid.UUID{}, err
		}
		return *showAdmin.ID, nil
	},
	"episodeId": func(
		ctx context2.Context, s internal.Services, arg any,
	) (uuid.UUID, error) {
		log.V("@isShowAdmin.episodeId: (%T) %+v", arg, arg)
		episodeID, err := uuid.FromString(arg.(string))
		if err != nil {
			return uuid.UUID{}, err
		}
		episode, err := s.EpisodeService.Get(ctx, internal.EpisodesFilter{
			ID: &episodeID,
		})
		if err != nil {
			return uuid.UUID{}, err
		}
		return *episode.ShowID, nil
	},
	"episodeUrl": func(
		ctx context2.Context, s internal.Services, arg any,
	) (uuid.UUID, error) {
		log.V("@isShowAdmin.episodeUrl: (%T) %+v", arg, arg)
		url := arg.(string)
		episodeURL, err := s.EpisodeURLService.Get(ctx, internal.EpisodeURLsFilter{
			URL: &url,
		})
		if err != nil {
			return uuid.UUID{}, err
		}
		episode, err := s.EpisodeService.Get(ctx, internal.EpisodesFilter{
			ID: episodeURL.EpisodeID,
		})
		if err != nil {
			return uuid.UUID{}, err
		}
		return *episode.ShowID, nil
	},
	"templateId": func(
		ctx context2.Context, s internal.Services, arg any,
	) (uuid.UUID, error) {
		log.V("@isShowAdmin.templateId: (%T) %+v", arg, arg)
		templateID, err := uuid.FromString(arg.(string))
		if err != nil {
			return uuid.UUID{}, err
		}
		template, err := s.TemplateService.Get(ctx, internal.TemplatesFilter{
			ID: &templateID,
		})
		if err != nil {
			return uuid.UUID{}, err
		}
		return *template.ShowID, nil
	},
}

func getShowIdFromParams(ctx context2.Context, params map[string]any, services internal.Services) (uuid.UUID, error) {
	names := []string{}
	for name, value := range params {
		if getter, ok := showIDGetters[name]; ok {
			return getter(ctx, services, value)
		}
		names = append(names, name)
	}
	return uuid.UUID{}, &internal.Error{
		Code: internal.EINTERNAL,
		Message: fmt.Sprintf(
			"No show id getter implemented for any of the args (%s)",
			strings.Join(names, ", "),
		),
		Op: "IsShowAdmin",
	}
}

func IsShowAdmin(ctx context2.Context, params any, next graphql2.Resolver) (any, error) {
	log.V("@isShowAdmin(%+v)", params)

	// Authenticate first, arg directives run before field directives (notably, `@authenticated``)
	ctx, err := authenticate(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}
	if auth.IsAdmin || auth.IsDev {
		log.V("@isShowAdmin - elevated role")
		return next(ctx)
	}

	services := context.GetServices(ctx)
	showID, err := getShowIdFromParams(ctx, params.(map[string]any), services)
	if err != nil {
		return nil, err
	}
	admins, err := services.ShowAdminService.List(ctx, internal.ShowAdminsFilter{
		ShowID: &showID,
	})
	if err != nil {
		return nil, err
	}

	for _, admin := range admins {
		if *admin.UserID == auth.UserID {
			return next(ctx)
		}
	}
	return nil, fmt.Errorf("You are is not an admin of this show (id=%s)", showID)
}

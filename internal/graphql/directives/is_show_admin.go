package directives

import (
	context2 "context"
	"fmt"
	"strings"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/context"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/log"
	graphql2 "github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
)

type showIDGetter = func(ctx context2.Context, s internal.DirectiveServices, arg interface{}) (uuid.UUID, error)

var showIDGetters = map[string]showIDGetter{
	"showId": func(
		ctx context2.Context, s internal.DirectiveServices, arg interface{},
	) (uuid.UUID, error) {
		log.V("@isShowAdmin.showId: (%T) %+v", arg, arg)
		return uuid.FromString(arg.(string))
	},
	"showAdminInput": func(
		ctx context2.Context, s internal.DirectiveServices, arg interface{},
	) (uuid.UUID, error) {
		log.V("@isShowAdmin.showAdminInput: (%T) %+v", arg, arg)
		showAdmin := arg.(*graphql.InputShowAdmin)
		return *showAdmin.ShowID, nil
	},
	"showAdminId": func(
		ctx context2.Context, s internal.DirectiveServices, arg interface{},
	) (uuid.UUID, error) {
		log.V("@isShowAdmin.showAdminId: (%T) %+v", arg, arg)
		// showAdminId := arg.(*uuid.UUID)
		// return *showAdmin.ShowID, nil
		panic("not implemented")
	},
	// episodeId: UUID
	// episodeUrl: String
	// newTemplate: InputTemplate
	// templateId: UUID
}

func getShowIdFromParams(ctx context2.Context, params map[string]interface{}, services internal.DirectiveServices) (uuid.UUID, error) {
	names := []string{}
	for name, value := range params {
		if getter, ok := showIDGetters[name]; ok {
			return getter(ctx, services, value)
		}
		names = append(names, name)
	}
	return uuid.UUID{}, fmt.Errorf("Internal error: @isShowAdmin not implemented for args (%s)", strings.Join(names, ", "))
}

func IsShowAdmin(ctx context2.Context, params interface{}, next graphql2.Resolver) (interface{}, error) {
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

	services := context.GetDirectiveServices(ctx)
	showID, err := getShowIdFromParams(ctx, params.(map[string]interface{}), services)
	if err != nil {
		return nil, err
	}
	admins, err := services.ShowAdminService.GetByShowID(ctx, showID)
	if err != nil {
		return nil, err
	}

	for _, admin := range admins {
		if admin.UserID == auth.UserID {
			return next(ctx)
		}
	}
	return nil, fmt.Errorf("You are is not an admin of this show (id=%s)", showID)
}

func AllowShowAdmin(ctx context2.Context, params interface{}, next graphql2.Resolver) (interface{}, error) {
	// Authenticate first, arg directives run before field directives (notably, `@authenticated``)
	ctx, err := authenticate(ctx)
	if err != nil {
		return nil, err
	}
	return next(ctx)
}

package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
)

type Resolver struct {
	// Give resolvers access to all the services
	*internal.Services
	DB internal.Database
}

// ApiClient returns graphql.ApiClientResolver implementation.
func (r *Resolver) ApiClient() graphql.ApiClientResolver { return &apiClientResolver{r} }

// Account returns graphql.AccountResolver implementation.
func (r *Resolver) Account() graphql.AccountResolver { return &accountResolver{r} }

// Episode returns graphql.EpisodeResolver implementation.
func (r *Resolver) Episode() graphql.EpisodeResolver { return &episodeResolver{r} }

// EpisodeUrl returns graphql.EpisodeUrlResolver implementation.
func (r *Resolver) EpisodeUrl() graphql.EpisodeUrlResolver { return &episodeUrlResolver{r} }

// ExternalLink returns graphql.ExternalLinkResolver implementation.
func (r *Resolver) ExternalLink() graphql.ExternalLinkResolver { return &externalLinkResolver{r} }

// Mutation returns graphql.MutationResolver implementation.
func (r *Resolver) Mutation() graphql.MutationResolver { return &mutationResolver{r} }

// Preferences returns graphql.PreferencesResolver implementation.
func (r *Resolver) Preferences() graphql.PreferencesResolver { return &preferencesResolver{r} }

// Query returns graphql.QueryResolver implementation.
func (r *Resolver) Query() graphql.QueryResolver { return &queryResolver{r} }

// Show returns graphql.ShowResolver implementation.
func (r *Resolver) Show() graphql.ShowResolver { return &showResolver{r} }

// ShowAdmin returns graphql.ShowAdminResolver implementation.
func (r *Resolver) ShowAdmin() graphql.ShowAdminResolver { return &showAdminResolver{r} }

// Template returns graphql.TemplateResolver implementation.
func (r *Resolver) Template() graphql.TemplateResolver { return &templateResolver{r} }

// TemplateTimestamp returns graphql.TemplateTimestampResolver implementation.
func (r *Resolver) TemplateTimestamp() graphql.TemplateTimestampResolver {
	return &templateTimestampResolver{r}
}

// ThirdPartyTimestamp returns graphql.ThirdPartyTimestampResolver implementation.
func (r *Resolver) ThirdPartyTimestamp() graphql.ThirdPartyTimestampResolver {
	return &thirdPartyTimestampResolver{r}
}

// Timestamp returns graphql.TimestampResolver implementation.
func (r *Resolver) Timestamp() graphql.TimestampResolver { return &timestampResolver{r} }

// TimestampType returns graphql.TimestampTypeResolver implementation.
func (r *Resolver) TimestampType() graphql.TimestampTypeResolver { return &timestampTypeResolver{r} }

// TimestampType returns graphql.TimestampTypeResolver implementation.
func (r *Resolver) TotalCounts() graphql.TotalCountsResolver { return &totalCountsResolver{r} }

// User returns graphql.UserResolver implementation.
func (r *Resolver) User() graphql.UserResolver { return &userResolver{r} }

type apiClientResolver struct{ *Resolver }
type accountResolver struct{ *Resolver }
type episodeResolver struct{ *Resolver }
type episodeUrlResolver struct{ *Resolver }
type externalLinkResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type preferencesResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type showResolver struct{ *Resolver }
type showAdminResolver struct{ *Resolver }
type templateResolver struct{ *Resolver }
type templateTimestampResolver struct{ *Resolver }
type thirdPartyTimestampResolver struct{ *Resolver }
type timestampResolver struct{ *Resolver }
type timestampTypeResolver struct{ *Resolver }
type totalCountsResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

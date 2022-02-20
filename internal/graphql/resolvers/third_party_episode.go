package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

func (r *thirdPartyEpisodeResolver) Timestamps(ctx context.Context, obj *graphql.ThirdPartyEpisode) ([]*graphql.ThirdPartyTimestamp, error) {
	panic("thirdPartyEpisodeResolver.Timestamps not implemented")
}

func (r *thirdPartyEpisodeResolver) Show(ctx context.Context, obj *graphql.ThirdPartyEpisode) (*graphql.ThirdPartyShow, error) {
	panic("thirdPartyEpisodeResolver.Show not implemented")
}

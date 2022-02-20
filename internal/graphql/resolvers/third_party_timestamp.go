package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

func (r *thirdPartyEpisodeResolver) Timestamps(ctx context.Context, obj *graphql.ThirdPartyEpisode) ([]*graphql.ThirdPartyTimestamp, error) {
	panic("not implemented")
}

func (r *thirdPartyEpisodeResolver) Show(ctx context.Context, obj *graphql.ThirdPartyEpisode) (*graphql.ThirdPartyShow, error) {
	panic("not implemented")
}

func (r *thirdPartyTimestampResolver) Type(ctx context.Context, obj *graphql.ThirdPartyTimestamp) (*graphql.TimestampType, error) {
	panic("not implemented")
}

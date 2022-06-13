package scalars

import (
	"fmt"
	"io"

	"anime-skip.com/public-api/internal"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
)

func MarshalUUID(id *uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		str := fmt.Sprintf(`"%s"`, id.String())
		w.Write([]byte(str))
	})
}

func UnmarshalUUID(v any) (*uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		id, err := uuid.FromString(v)
		return &id, err
	default:
		return nil, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: fmt.Sprintf("%v is not a string", v),
			Op:      "UnmarshalUUID",
		}
	}
}

package scalars

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
)

func MarshalUUID(id *uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		str := fmt.Sprintf(`"%s"`, id.String())
		w.Write([]byte(str))
	})
}

func UnmarshalUUID(v interface{}) (*uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		id, err := uuid.FromString(v)
		return &id, err
	default:
		return nil, fmt.Errorf("%v is not a string", v)
	}
}

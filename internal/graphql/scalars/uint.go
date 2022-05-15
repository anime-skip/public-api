package scalars

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalUInt(i *uint) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		str := fmt.Sprintf("%d", i)
		w.Write([]byte(str))
	})
}

func UnmarshalUInt(v any) (*uint, error) {
	switch v := v.(type) {
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		return UnmarshalUInt(i)
	case int:
		if v < 0 {
			return nil, errors.New("Uint cannot be less than 0")
		}
		ui := uint(v)
		return &ui, nil
	case int64:
		if v < 0 {
			return nil, errors.New("Uint cannot be less than 0")
		}
		ui := uint(v)
		return &ui, nil
	default:
		return nil, fmt.Errorf("%v is not a string", v)
	}
}

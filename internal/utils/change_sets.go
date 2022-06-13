package utils

import (
	"reflect"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

// https://gqlgen.com/reference/changesets/
func ApplyChanges(changes map[string]any, to any) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  false, // This was changed to false because when true, it resets pointer fields to nil (like deleted_at)
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v any) (any, error) {
			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}

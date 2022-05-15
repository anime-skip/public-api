package context

import (
	"context"

	"anime-skip.com/public-api/internal"
)

var ipAddressKey = &contextKey{"ip_address"}

func WithIPAddress(ctx context.Context, ipAddress string) context.Context {
	return context.WithValue(ctx, ipAddressKey, ipAddress)
}

func GetIPAddress(ctx context.Context) (string, error) {
	value, ok := ctx.Value(ipAddressKey).(string)
	if !ok {
		return "", &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "IP Address is not set yet, is it applied in a middleware?",
			Op:      "GetAuthClaims",
		}
	}
	return value, nil
}

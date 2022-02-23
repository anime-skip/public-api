package context

import (
	"context"
	"errors"
)

var ipAddressKey = &contextKey{"ip_address"}

func WithIPAddress(ctx context.Context, ipAddress string) context.Context {
	return context.WithValue(ctx, ipAddressKey, ipAddress)
}

func GetIPAddress(ctx context.Context) (string, error) {
	value, ok := ctx.Value(ipAddressKey).(string)
	if !ok {
		return "", errors.New("IP Address is not set yet, is it applied in a middleware?")
	}
	return value, nil
}

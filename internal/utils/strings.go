package utils

import (
	"fmt"
	"net/url"
	"strings"

	"anime-skip.com/public-api/internal"
	"github.com/samber/lo"
)

type SanitizeURLOptions struct {
	// "https", "http", etc. No ":" at the end
	AllowedSchemes   []string
	AllowedHostnames []string
	KeepQueryParams  []string
	KeepFragment     bool
}

func SanitizeURL(inputURL string, options SanitizeURLOptions) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", &internal.Error{
			Code:    internal.EINVALID,
			Message: "URL could not be parsed",
			Op:      "SanitizeURL",
			Err:     err,
		}
	}

	finalURL := url.URL{}

	// Scheme:
	if len(options.AllowedSchemes) > 0 && !lo.Contains(options.AllowedSchemes, parsedURL.Scheme) {
		return "", &internal.Error{
			Code: internal.EINVALID,
			Message: fmt.Sprintf(
				"URL does not have the required scheme (allowed: %s | url: %s | scheme: %s)",
				strings.Join(options.AllowedSchemes, ", "),
				inputURL,
				parsedURL.Scheme,
			),
			Op: "SanitizeURL",
		}
	}
	finalURL.Scheme = parsedURL.Scheme

	// host.name
	parsedHostname := parsedURL.Hostname()
	if len(options.AllowedHostnames) > 0 && !lo.Contains(options.AllowedHostnames, parsedHostname) {
		return "", &internal.Error{
			Code: internal.EINVALID,
			Message: fmt.Sprintf(
				"URL does not have the required hostname (allowed: %s | url: %s | hostname: %s)",
				strings.Join(options.AllowedHostnames, ", "),
				inputURL,
				parsedHostname,
			),
			Op: "SanitizeURL",
		}
	}
	finalURL.Host = parsedHostname

	// #fragment
	if options.KeepFragment {
		finalURL.Fragment = parsedURL.Fragment
	}

	// ?query=params
	if len(options.KeepQueryParams) > 0 {
		finalQuery := url.Values{}
		for _, param := range options.KeepQueryParams {
			if parsedURL.Query().Has(param) {
				finalQuery.Set(param, parsedURL.Query().Get(param))
			}
		}
		finalURL.RawQuery = finalQuery.Encode()
		_, err := url.ParseQuery(finalURL.RawQuery)
		if err != nil {
			return "", err
		}
	}

	// /path
	finalURL.Path = parsedURL.Path

	return finalURL.String(), nil
}

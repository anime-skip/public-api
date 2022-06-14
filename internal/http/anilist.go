package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/samber/lo"
)

type AnilistService struct {
	client *http.Client
}

func NewAnilistService() *AnilistService {
	return &AnilistService{
		client: http.DefaultClient,
	}
}

func (a *AnilistService) request(query string, variables map[string]any) (map[string]any, error) {
	bodyJson := map[string]any{
		"query":     query,
		"variables": variables,
	}
	body, err := json.Marshal(bodyJson)
	if err != nil {
		return nil, err
	}

	res, err := a.client.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 300 {
		return nil, errors.New(res.Status)
	}

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]any
	err = json.Unmarshal(resBytes, &data)
	if gqlErrors, ok := data["errors"]; ok {
		return nil, fmt.Errorf("%+v", gqlErrors)
	}
	return data, err
}

func (a *AnilistService) FindLink(showName string) (*string, error) {
	data, err := a.request(`
		query ($search: String) {
			anime: Page(perPage: 1) {
				results: media(type: ANIME, search: $search) {
					title {
						english
					}
					siteUrl
				}
			}
		}
	`, map[string]any{"search": showName})
	if err != nil {
		return nil, err
	}
	results := data["data"].(map[string]any)["anime"].(map[string]any)["results"].([]any)
	if len(results) == 0 {
		return nil, nil
	}

	return lo.ToPtr(results[0].(map[string]any)["siteUrl"].(string)), err
}

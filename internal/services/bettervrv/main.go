package bettervrv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"anime-skip.com/backend/internal/graphql/models"
	. "anime-skip.com/backend/internal/services/bettervrv/models"
	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/log"
)

type betterVRVServiceInterface struct{}
type CachedResponse struct {
	Episodes []*models.ThirdPartyEpisode
	CachedAt time.Time
}

var BetterVRV = betterVRVServiceInterface{}
var localCache map[string]*CachedResponse = map[string]*CachedResponse{}

const BASE_URL = "https://parseapi.back4app.com"
const APP_ID_KEY = "x-parse-application-id"
const API_KEY_KEY = "X-Parse-REST-API-Key"

var APP_ID_VALUE = utils.ENV.BETTER_VRV_APP_ID
var API_KEY_VALUE = utils.ENV.BETTER_VRV_API_KEY
var UNKNOWN_EPISODE = &models.ThirdPartyEpisode{}
var CACHE_DURATION = 30 * time.Minute

func createRequest(endpoint string, query map[string]string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("GET", BASE_URL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	if query != nil {
		q := req.URL.Query()
		for key, value := range query {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Add(APP_ID_KEY, APP_ID_VALUE)
	req.Header.Add(API_KEY_KEY, API_KEY_VALUE)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req, nil
}

func (betterVRVService betterVRVServiceInterface) FetchEpisodesByName(episodeName string) ([]*models.ThirdPartyEpisode, error) {
	cachedResult, ok := localCache[episodeName]
	isCached := ok && cachedResult.CachedAt.Add(CACHE_DURATION).After(time.Now())
	if isCached {
		return cachedResult.Episodes, nil
	}

	log.V("Fetching new episode from BetterVRV")
	remoteResult, err := fetchRemoteEpisodesByName(episodeName)
	if err != nil || remoteResult == nil {
		return nil, err
	}
	localCache[episodeName] = &CachedResponse{
		Episodes: remoteResult,
		CachedAt: time.Now(),
	}
	return remoteResult, nil
}

func (betterVRVService betterVRVServiceInterface) FetchShowById(id string) (*models.ThirdPartyShow, error) {
	log.V("Fetching show from BetterVRV")
	return fetchRemoteShowsById(id)
}

func fetchRemoteEpisodesByName(episodeName string) ([]*models.ThirdPartyEpisode, error) {
	inputEpisodeName := strings.ReplaceAll(episodeName, "\"", "\\\"")
	inputEpisodeName = strings.ReplaceAll(inputEpisodeName, "\\", "\\\\")
	queryParams := map[string]string{
		"where": fmt.Sprintf("{ \"episodeTitle\": \"%s\" }", inputEpisodeName),
	}
	req, err := createRequest("/classes/Timestamps", queryParams, nil)
	if err != nil {
		return nil, err
	}
	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Parse response
	response := &BetterVRVEpisodeResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	// Map list
	episodes := []*models.ThirdPartyEpisode{}
	if response.Results != nil {
		for _, item := range response.Results {
			mappedItem := MapBetterVRVEpisodeToThirdPartyEpisode(&item)
			if mappedItem != nil {
				episodes = append(episodes, mappedItem)
			}
		}
	}

	return episodes, nil
}

func fetchRemoteShowsById(id string) (*models.ThirdPartyShow, error) {
	queryParams := map[string]string{
		"where": fmt.Sprintf("{ \"objectId\": \"%s\" }", id),
	}
	req, err := createRequest("/classes/Series", queryParams, nil)
	if err != nil {
		return nil, err
	}
	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Parse response
	response := &BetterVRVShowResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	// Map list
	return MapBetterVRVShowToThirdPartyShow(&response.Results[0]), nil
}

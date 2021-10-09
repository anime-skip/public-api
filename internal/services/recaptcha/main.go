package recaptcha

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"anime-skip.com/backend/internal/utils/env"
	log "anime-skip.com/backend/internal/utils/log"
)

var recaptcha_secret = env.RECAPTCHA_SECRET
var recaptcha_response_allowlist = env.RECAPTCHA_RESPONSE_ALLOWLIST

const recaptchaURL = "https://www.google.com/recaptcha/api/siteverify?secret=%s&response=%s&remoteip=%s"
const errorMessage = "Recaptcha validation failed"

func Verify(response, ipAddress string) error {
	if contains(recaptcha_response_allowlist, response) {
		return nil
	}
	resp, err := http.Post(fmt.Sprintf(recaptchaURL, recaptcha_secret, response, ipAddress), "application/json", nil)
	if err != nil {
		log.E("(VerifyRecaptcha) Failed to communicate: %v", err)
		return errors.New(errorMessage)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.E("(VerifyRecaptcha) Could not read response body: %v", err)
		return errors.New(errorMessage)
	}

	var responseJson map[string]interface{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		log.E("(VerifyRecaptcha) Response body was not valid JSON: %v", err)
		return errors.New(errorMessage)
	}
	if responseJson["success"] != true {
		return errors.New(errorMessage)
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

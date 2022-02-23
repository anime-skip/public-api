package http

import (
	go_context "context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/context"
	"anime-skip.com/timestamps-service/internal/log"
	"anime-skip.com/timestamps-service/internal/utils"
)

type googleRecaptchaService struct {
	secret            string
	responseAllowList []string
}

func NewGoogleRecaptchaService(secret string, responseAllowList []string) internal.RecaptchaService {
	return &googleRecaptchaService{
		secret:            secret,
		responseAllowList: responseAllowList,
	}
}

const errorMessage = "Recaptcha validation failed"

func (s *googleRecaptchaService) Verify(ctx go_context.Context, response string) error {
	ipAddress, err := context.GetIPAddress(ctx)
	if err != nil {
		return errors.New("Could not get ip address from request")
	}

	// Skip http verification when response matches allowlist
	// The allowlist is only setup for local development so you can create accounts without
	// needing a UI with recaptcha
	if utils.StringSliceIncludes(s.responseAllowList, response) {
		return nil
	}

	url := fmt.Sprintf(
		"https://www.google.com/recaptcha/api/siteverify?secret=%s&response=%s&remoteip=%s",
		s.secret,
		response,
		ipAddress,
	)
	resp, err := http.Post(url, "application/json", nil)
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

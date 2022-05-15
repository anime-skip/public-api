package http

import (
	go_context "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/utils"
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

var genericRecaptchaFailure = &internal.Error{
	Code:    internal.EINTERNAL,
	Message: "Recaptcha validation failed",
	Op:      "RecaptchaService.Verify",
}

func (s *googleRecaptchaService) Verify(ctx go_context.Context, response string) error {
	ipAddress, err := context.GetIPAddress(ctx)
	if err != nil {
		return &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Could not get ip address from request",
			Op:      "RecaptchaService.Verify",
			Err:     err,
		}
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
		return genericRecaptchaFailure
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.E("(VerifyRecaptcha) Could not read response body: %v", err)
		return genericRecaptchaFailure
	}

	var responseJson map[string]any
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		log.E("(VerifyRecaptcha) Response body was not valid JSON: %v", err)
		return genericRecaptchaFailure
	}
	if responseJson["success"] != true {
		return genericRecaptchaFailure
	}

	return nil
}

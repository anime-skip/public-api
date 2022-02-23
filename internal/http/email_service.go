package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"anime-skip.com/timestamps-service/internal"
)

type animeSkipEmailService struct {
	client  *http.Client
	host    string
	secret  string
	enabled bool
}

func NewAnimeSkipEmailService(
	host string,
	secret string,
	enabled bool,
) internal.EmailService {
	return &animeSkipEmailService{
		client:  &http.Client{},
		host:    host,
		secret:  secret,
		enabled: enabled,
	}
}

var httpClient *http.Client = &http.Client{}

func (s *animeSkipEmailService) sendEmail(endpoint string, body map[string]interface{}) error {
	if !s.enabled {
		return nil
	}

	url := fmt.Sprintf("http://%s/%s", s.host, endpoint)
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Secret "+s.secret)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("Send email request failed with status %s", resp.Status)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (s *animeSkipEmailService) SendWelcome(user internal.User) error {
	return s.sendEmail("welcome", map[string]interface{}{
		"emails":   []string{user.Email},
		"username": user.Username,
	})
}

func (s *animeSkipEmailService) SendVerification(user internal.User, token string) error {
	return s.sendEmail("verification", map[string]interface{}{
		"emails": []string{user.Email},
		"token":  token,
	})
}

func (s *animeSkipEmailService) SendResetPassword(user internal.User, token string) error {
	return s.sendEmail("reset-password", map[string]interface{}{
		"emails": []string{user.Email},
		"token":  token,
	})
}

package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"anime-skip.com/public-api/internal"
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

func (s *animeSkipEmailService) sendEmail(ctx context.Context, endpoint string, body map[string]any) error {
	if !s.enabled {
		return nil
	}

	client := http.Client{
		Timeout: 10 * time.Second,
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
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Secret "+s.secret)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return &internal.Error{
			Code:    internal.EINTERNAL,
			Message: fmt.Sprintf("Send email request failed with status %s", resp.Status),
		}
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (s *animeSkipEmailService) SendWelcome(ctx context.Context, user internal.FullUser) error {
	return s.sendEmail(ctx, "welcome", map[string]any{
		"emails":   []string{user.Email},
		"username": user.Username,
	})
}

func (s *animeSkipEmailService) SendVerification(ctx context.Context, user internal.FullUser, token string) error {
	return s.sendEmail(ctx, "verification", map[string]any{
		"emails": []string{user.Email},
		"token":  token,
	})
}

func (s *animeSkipEmailService) SendResetPassword(ctx context.Context, user internal.FullUser, token string) error {
	return s.sendEmail(ctx, "reset-password", map[string]any{
		"emails": []string{user.Email},
		"token":  token,
	})
}

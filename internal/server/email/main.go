package email

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
)

var EMAIL_HOST string = utils.EnvString("EMAIL_SERVICE_HOST")
var EMAIL_SECRET string = utils.EnvString("EMAIL_SERVICE_SECRET")
var httpClient *http.Client = &http.Client{}

func sendEmail(endpoint string, body map[string]interface{}) error {
	url := fmt.Sprintf("http://%s/%s", EMAIL_HOST, endpoint)
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Secret "+EMAIL_SECRET)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return errors.New(fmt.Sprintf("Send email request failed with status %s", resp.Status))
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func SendWelcome(account *entities.User) error {
	return sendEmail("welcome", map[string]interface{}{
		"emails":   []string{account.Email},
		"username": account.Username,
	})
}

func SendVerification(account *entities.User, token string) error {
	return sendEmail("verification", map[string]interface{}{
		"emails": []string{account.Email},
		"token":  token,
	})
}

func SendChangePassword(account *entities.User, token string) error {
	return sendEmail("change_password", map[string]interface{}{
		"emails":   []string{account.Email},
		"username": account.Username,
		"token":    token,
	})
}

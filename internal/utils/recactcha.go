package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var recaptcha_secret string = EnvString("RECAPTCHA_SECRET")
var recaptcha_response_allowlist []string = EnvStringArray("RECAPTCHA_RESPONSE_ALLOWLIST")

const recaptchaURL = "https://www.google.com/recaptcha/api/siteverify?secret=%s&response=%s&remoteip=%s"
const errorMessage = "Recaptacha validation failed"

func VerifyRecaptcha(response, ipAddress string) error {
	if contains(recaptcha_response_allowlist, response) {
		time.Sleep(2 * time.Second)
		return nil
	}
	resp, err := http.Post(fmt.Sprintf(recaptchaURL, recaptcha_secret, response, ipAddress), "application/json", nil)
	if err != nil {
		fmt.Println(err)
		return errors.New(errorMessage)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return errors.New(errorMessage)
	}
	fmt.Println(string(body))

	var responseJson map[string]interface{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		fmt.Println(err)
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

package main

import (
	"fmt"
	"os"

	"github.com/aklinker1/anime-skip-backend/internal/server/email"
)

func main() {
	err := email.Email{
		To:       []string{"aaronklinker1@gmail.com"},
		Subject:  "HTML Test",
		Template: "email_welcome.html",
		TemplateData: map[string]string{
			"username": "aklinker1",
		},
	}.Send()
	if err == nil {
		os.Exit(0)
	}
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}

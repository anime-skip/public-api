package email

import "github.com/aklinker1/anime-skip-backend/internal/database/entities"

func SendWelcome(account *entities.User) error {
	email := Email{
		To:       []string{account.Email},
		Subject:  "Welcome to Anime Skip",
		Template: "email_welcome.html",
		TemplateData: map[string]string{
			"username": account.Username,
		},
	}
	return email.Send()
}

func SendEmailAddressValidation(account *entities.User, token string) error {
	email := Email{
		To:       []string{account.Email},
		Subject:  "Validate Email Address",
		Template: "email_validate.html",
		TemplateData: map[string]string{
			"token": token,
		},
	}
	return email.Send()
}

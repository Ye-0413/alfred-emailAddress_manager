package internal

import (
	"os"
	"strings"
)

// EmailAddress represents a stored email address with a name
type EmailAddress struct {
	Name  string
	Email string
}

// GetEmailAddresses returns email addresses from environment configuration
func GetEmailAddresses() []EmailAddress {
	var emails []EmailAddress

	// Try to get email configuration from environment variable
	emailConfig := os.Getenv("EMAIL_CONFIG")

	// If environment variable is set, parse it
	if emailConfig != "" {
		emailPairs := strings.Split(emailConfig, ",")

		for _, pair := range emailPairs {
			// Split each pair into name and email
			parts := strings.SplitN(pair, ":", 2)
			if len(parts) == 2 {
				emails = append(emails, EmailAddress{
					Name:  parts[0],
					Email: parts[1],
				})
			}
		}
	} else {
		// Fallback to example email addresses if no configuration is provided
		emails = []EmailAddress{
			{Name: "Personal", Email: "your.name@example.com"},
			{Name: "Work", Email: "work@example.com"},
			{Name: "School", Email: "student@university.edu"},
			{Name: "Gmail Example", Email: "example@gmail.com"},
			{Name: "Alternative", Email: "alternative@example.com"},
		}
	}

	return emails
}

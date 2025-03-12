package helpers

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

func isEmailFormatValid(email string) bool {
	// Basic email regex pattern
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

// hasMXRecords checks if the domain has valid MX records
func hasMXRecords(domain string) bool {
	mx, err := net.LookupMX(domain)
	return err == nil && len(mx) > 0
}

func checkSMTP(email string) (bool, error) {
	parts := strings.Split(email, "@")
	domain := parts[1]

	mx, err := net.LookupMX(domain)
	if err != nil || len(mx) == 0 {
		return false, fmt.Errorf("no MX records found")
	}

	// Connect to the first MX server
	client, err := net.DialTimeout("tcp", fmt.Sprintf("%s:25", mx[0].Host), 10*time.Second)
	if err != nil {
		return false, fmt.Errorf("couldn't connect to mail server: %v", err)
	}
	defer client.Close()
	return true, nil
}

// validateEmail performs all checks on an email
func ValidateEmail(email string) map[string]interface{} {
	result := map[string]interface{}{
		"email":       email,
		"validFormat": false,
		"hasMX":       false,
		"smtpCheck":   nil,
		"reachable":   false,
	}

	// Check format
	result["validFormat"] = isEmailFormatValid(email)
	if !result["validFormat"].(bool) {
		return result
	}

	parts := strings.Split(email, "@")
	domain := parts[1]

	// Check MX records
	result["hasMX"] = hasMXRecords(domain)
	if !result["hasMX"].(bool) {
		return result
	}

	// Attempt SMTP check
	smtpReachable, smtpErr := checkSMTP(email)
	if smtpErr != nil {
		result["smtpCheck"] = smtpErr.Error()
	} else {
		result["smtpCheck"] = "Connection successful"
		result["reachable"] = smtpReachable
	}
	return result
}

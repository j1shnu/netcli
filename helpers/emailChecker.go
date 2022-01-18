package helpers

import (
	"encoding/json"
	"fmt"
)

type Email struct {
	ValidFormat bool `json:"validFormat"`
	Deliverable bool `json:"deliverable"`
}

func EmailChecker(email string) string {
	url := "https://api.trumail.io/v2/lookups/json?email=" + email
	validate := GetData(url)
	var emailData Email
	err := json.Unmarshal(validate, &emailData)
	ErrorHandler(err)
	return fmt.Sprintf("Deliverable :- %v\nIs in valid format :- %v", emailData.Deliverable, emailData.ValidFormat)
}

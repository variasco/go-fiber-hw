package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

func FormatErrors(errors *validate.Errors) []string {
	var messages []string

	for _, value := range errors.Errors {
		messages = append(messages, strings.Join(value, ", "))
	}

	return messages
}

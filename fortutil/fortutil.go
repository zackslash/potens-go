package fortutil

import (
	"errors"
	"regexp"
	"strings"
)

var (

	// ErrInvalidID Error for an invalid ID
	ErrInvalidID = errors.New("The ID specified is invalid")
)

// CreateID converts a string to a valid ID
func CreateID(input string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	output := reg.ReplaceAllString(input, "-")
	output = strings.ToLower(output)
	output = strings.Trim(output, "-")
	return output
}

// ValidateID validates an ID is a valid fortifi id (Not FID)
func ValidateID(input string) error {
	reg, _ := regexp.Compile("^[a-z0-9][a-z0-9\\-]+[a-z0-9]$")
	if !reg.MatchString(input) {
		return ErrInvalidID
	}
	return nil
}

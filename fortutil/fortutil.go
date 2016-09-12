package fortutil

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidID = errors.New("The ID specified is invalid")
)

func CreateID(input string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	output := reg.ReplaceAllString(input, "-")
	output = strings.ToLower(output)
	output = strings.Trim(output, "-")
	return output
}

func ValidateID(input string) error {
	reg, _ := regexp.Compile("^[a-z0-9][a-z0-9-]+[a-z0-9]$")
	if !reg.MatchString(input) {
		return ErrInvalidID
	}
	return nil
}

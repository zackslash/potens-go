package cubeutil

import (
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var (

	// ErrInvalidID Error for an invalid ID
	ErrInvalidID          = errors.New("The ID specified is invalid")
	ErrInvalidGlobalAppID = errors.New("The Global App ID specified is invalid")
)

//SplitGaID Split a Global App ID into Vendor and App ID
func SplitGaID(gaid string) (string, string, error) {
	parts := strings.SplitN(gaid, "/", 3)
	if len(parts) != 2 {
		return "", "", errors.New("Invalid GAID Provided")
	}
	return parts[0], parts[1], nil
}

// CreateID converts a string to a valid ID
func CreateID(input string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	output := reg.ReplaceAllString(input, "-")
	output = strings.ToLower(output)
	output = strings.Trim(output, "-")
	return output
}

// ValidateID validates an ID is a valid cubex id (Not FID)
func ValidateID(input string) error {
	reg, _ := regexp.Compile("^[a-z0-9][a-z0-9\\-]+[a-z0-9]$")
	if !reg.MatchString(input) {
		return ErrInvalidID
	}
	return nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandomAlphaNum generates a random alpha numeric string
func RandomAlphaNum(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// Acronym forms an acronym from an input string
func Acronym(source string, length int) string {
	nameParts := strings.Split(source, " ")
	if length == 0 {
		length = len(nameParts)
	}
	acronym := ""
	for i, part := range nameParts {
		for x := 0; x < length; x++ {
			acronym += string([]rune(part)[x])
			if i+1 < len(nameParts) || len(acronym) == length {
				break
			}
		}
		if len(acronym) == length {
			break
		}
	}
	return strings.ToUpper(acronym)
}

func ValidateGlobalAppID(gapid string) error {
	if strings.Trim(gapid, "/") != gapid {
		return ErrInvalidGlobalAppID
	}
	if !strings.Contains(gapid, "/") {
		return ErrInvalidGlobalAppID
	}
	if len(gapid) < 3 {
		return ErrInvalidGlobalAppID
	}
	return nil
}

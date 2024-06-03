package helpers

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func ValidatePattern(format string, str string) bool {
	// emailF, _ := regexp.Compile(os.Getenv("REGEX_MAIL"))
	var email = os.Getenv("REGEX_EMAIL")
	phoneF := os.Getenv("REGEX_PHONE")

	if strings.Contains(strings.ToLower(format), "email") {
		if match, _ := regexp.MatchString(email, str); !match {
			return false
		}
	}

	if slices.Contains([]string{"phone", "mobile", "telephone", "hp", "smarthone", "telp", "cell phone", "cell-phone", "kontak", "contact"}, strings.ToLower(format)) {
		if match, _ := regexp.MatchString(phoneF, str); !match {
			return false
		}
	}

	if slices.Contains([]string{"post", "post code", "kode pos", "pos"}, strings.ToLower(format)) {
		if match, _ := regexp.MatchString("abcdefghizklmnopqrstuvwxyz@#$;,.}|-_", str); match {
			fmt.Println(str)
			return false
		}
	}
	return true
}

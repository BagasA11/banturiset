package helpers

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

const RekeningPattern = "^\\d{10,16}$"
const EmailPattern = "@"

func ValidatePattern(format string, str string) bool {
	forbidden_char := "{}/|<>"
	// emailF, _ := regexp.Compile(os.Getenv("REGEX_MAIL"))

	phoneF := os.Getenv("REGEX_PHONE")
	postf := os.Getenv("REGEX_POST")
	if match, _ := regexp.MatchString(forbidden_char, str); match {
		fmt.Println(match)
		return false
	}

	if strings.Contains(strings.ToLower(format), "email") {
		if match, _ := regexp.MatchString(EmailPattern, str); !match {
			return false
		}
	}

	if slices.Contains([]string{"phone", "mobile", "telephone", "hp", "smarthone", "telp", "cell phone", "cell-phone", "kontak", "contact"}, strings.ToLower(format)) {
		if match, _ := regexp.MatchString(phoneF, str); !match {
			return false
		}
	}

	if slices.Contains([]string{"post", "post code", "kode pos", "pos"}, strings.ToLower(format)) {
		if match, _ := regexp.MatchString(postf, str); !match {
			fmt.Println(str)
			return false
		}
	}
	return true
}

func ValidateRekening(input string) bool {
	if match, _ := regexp.MatchString(RekeningPattern, input); !match {
		return false
	}
	return true
}

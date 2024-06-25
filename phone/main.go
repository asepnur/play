package main

import (
	"regexp"
)

var regex = "^(628)|^(\\+628)|^(08)|^(8)"

func main() {

}

func SanitizePhoneNumber(phoneNumber string) string {
	re := regexp.MustCompile(regex)
	finalPhoneNumber := re.ReplaceAllString(phoneNumber, "628")
	return finalPhoneNumber
}

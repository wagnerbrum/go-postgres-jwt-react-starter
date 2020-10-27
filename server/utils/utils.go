package utils

import (
	"fmt"
	"regexp"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

//ValidateEmail - validate an email
func ValidateEmail(email string) (err bool, message string) {
	emailCheck := regexp.MustCompile(emailRegex).MatchString(email)

	if emailCheck {
		return false, ""
	}

	return true, "Invalid email, please enter a email valid"
}

//ValidateStrMinLength - validate min length
func ValidateStrMinLength(value, properyName string, length int16) (err bool, message string) {
	strLen := int16(len(value))

	if length == 0 {
		return true, ""
	}

	if strLen < length {
		return true, fmt.Sprintf("the '%s' field must contain at least %d characters", properyName, length)
	}

	return false, ""
}

//ValidateStrMaxLength - validate max length
func ValidateStrMaxLength(value, properyName string, length int16) (err bool, message string) {
	strLen := int16(len(value))

	if length == 0 {
		return true, ""
	}

	if strLen > length {
		return true, fmt.Sprintf("the '%s' field must contain a maximum of %d characters", properyName, length)
	}

	return false, ""
}

//ValidateStrMinMaxLength - validate between min and max length
func ValidateStrMinMaxLength(value, properyName string, minLength int16, maxLength int16) (err bool, message string) {
	strLen := int16(len(value))

	if minLength == 0 || maxLength == 0 {
		return true, ""
	}

	if strLen > minLength || strLen < maxLength {
		return true, fmt.Sprintf("the '%s' field must contain between '%d' and '%d' characters", properyName, minLength, maxLength)
	}

	return false, ""
}

//ValidateStrExactLength - validate length is exactly
func ValidateStrExactLength(value, properyName string, length int16) (err bool, message string) {
	strLen := int16(len(value))

	if length == 0 {
		return true, ""
	}

	if strLen != length {
		return true, fmt.Sprintf("the '%s' field must contain exactly %d characters", properyName, length)
	}

	return false, ""
}

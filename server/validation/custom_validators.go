package validation

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

func CustomFutureValidator(fl validator.FieldLevel) bool {
	if fl.Field().Type() != reflect.TypeOf(time.Time{}) {
		return false
	}
	timeValue, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}
	if timeValue.After(time.Now()) {
		return true
	}
	return false
}

func CustomPasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	length := utf8.RuneCountInString(password)
	if length < 8 {
		return false
	}

	hasUpperCase := false
	hasLowerCase := false
	hasDigit := false
	hasSpecialSymbol := false

	for _, char := range password {
		if strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", char) {
			hasUpperCase = true
		} else if strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", char) {
			hasLowerCase = true
		} else if strings.ContainsRune("0123456789", char) {
			hasDigit = true
		} else if matched, _ := regexp.MatchString(`[^a-zA-Z0-9]`, string(char)); matched {
			hasSpecialSymbol = true
		}
	}

	return hasUpperCase && hasLowerCase && hasDigit && hasSpecialSymbol
}

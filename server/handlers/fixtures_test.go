package handlers_test

import (
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/dto"
)

const (
	UserPassword   = "testTEST123*"
	UserEmail      = "user@example.com"
	AdminUserEmail = "admin@example.com"
)

func getBaseRegisterInput() dto.Register {
	return dto.Register{
		FirstName:       "John",
		LastName:        "Dou",
		Email:           UserEmail,
		CurrentCountry:  "US",
		CountryOfBirth:  "UA",
		Gender:          constants.GenderMale,
		Timezone:        "Pacific/Midway",
		Birthday:        "01-01-1984",
		AgreeTerms:      true,
		Photo:           "example.jpg",
		Password:        UserPassword,
		ConfirmPassword: UserPassword,
	}
}

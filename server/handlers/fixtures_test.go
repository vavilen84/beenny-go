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
		Birthday:        "1984-01-23",
		Photo:           "/2024/01/23/s09d8fs09dfu.jpg",
		AgreeTerms:      true,
		Password:        UserPassword,
		ConfirmPassword: UserPassword,
	}
}

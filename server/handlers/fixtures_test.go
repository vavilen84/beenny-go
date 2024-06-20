package handlers_test

import (
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/dto"
)

const (
	testUserPassword   = "testTEST123*"
	testUserEmail      = "user@beenny.com"
	testAdminUserEmail = "admin@example.com"
)

func getBaseRegisterInput() dto.Register {
	return dto.Register{
		FirstName:       "John",
		LastName:        "Dou",
		Email:           testUserEmail,
		CurrentCountry:  "US",
		CountryOfBirth:  "UA",
		Gender:          constants.GenderMale,
		Timezone:        "Pacific/Midway",
		Birthday:        "1984-01-23",
		Photo:           "/2024/01/23/s09d8fs09dfu.jpg",
		AgreeTerms:      true,
		Password:        testUserPassword,
		ConfirmPassword: testUserPassword,
	}
}

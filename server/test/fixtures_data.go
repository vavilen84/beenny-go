package test

import (
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/dto"
)

const (
	TestUserEmail    = "user@example.com"
	TestUserPassword = "testTEST123*"
)

func GetValidRegisterUserDTO() dto.Register {
	return dto.Register{
		FirstName:       "John",
		LastName:        "Dou",
		CurrentCountry:  "UA",
		CountryOfBirth:  "UA",
		Gender:          constants.GenderMale,
		Timezone:        "US/Arizona",
		Birthday:        "1984-01-23",
		Photo:           "/2024/01/23/s09d8fs09dfu.jpg",
		Email:           TestUserEmail,
		Password:        TestUserPassword,
		ConfirmPassword: TestUserPassword,
		AgreeTerms:      true,
	}
}

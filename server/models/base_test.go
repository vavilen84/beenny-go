package models

import (
	"github.com/vavilen84/beenny-go/constants"
)

func GetTestValidUserModel() User {
	return User{
		FirstName:      "John",
		LastName:       "Dou",
		Email:          "email@example.com",
		CurrentCountry: "UA",
		CountryOfBirth: "UA",
		Gender:         constants.GenderMale,
		Timezone:       "US/Arizona",
		Birthday:       "1984-01-23",
		Password:       "12345678lT*",
		Photo:          "/2024/01/23/9edfcc70-3e89-447f-833b-954208f9463a.png",
		Role:           constants.RoleUser,
		EmailTwoFaCode: "123456",
	}
}

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
		Photo:          "/2024/01/23/s09d8fs09dfu.jpg",
		Role:           constants.RoleUser,
		EmailTwoFaCode: "123456",
	}
}

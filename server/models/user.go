package models

import (
	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id                          int        `json:"id" column:"id" gorm:"primaryKey;autoIncrement:true"`
	FirstName                   string     `json:"firstName"`
	LastName                    string     `json:"lastName"`
	Email                       string     `json:"email"`
	Password                    string     `json:"-"`
	PasswordSalt                string     `json:"-"`
	PasswordResetToken          string     `json:"-"`
	PasswordResetTokenExpiresAt *time.Time `json:"-"`
	Role                        int        `json:"role"`
	IsEmailVerified             bool       `json:"isEmailVerified"`
	CurrentCountry              string     `json:"currentCountry"`
	CountryOfBirth              string     `json:"countryOfBirth"`
	Gender                      string     `json:"gender"`
	Timezone                    string     `json:"timezone"`
	Birthday                    string     `json:"birthday"`
	Photo                       string     `json:"photo"`
	EmailTwoFaCode              string     `json:"-"`
	AuthToken                   string     `json:"authToken" gorm:"-"`
}

func (m *User) TableName() string {
	return "users"
}

func (User) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioCreate: validation.FieldRules{
			"FirstName":      "max=255,required",
			"LastName":       "max=255,required",
			"Email":          "max=255,email,required",
			"CurrentCountry": "max=2,required",
			"CountryOfBirth": "max=2,required",
			"Gender":         "max=10,required",
			"Timezone":       "max=255,required",
			"Birthday":       "required",
			"Photo":          "max=255,required",
			"Password":       "max=255,required,customPasswordValidator",
			"Role":           "required,gt=0,lt=2",
			"EmailTwoFaCode": "required",
		},
		constants.ScenarioHashPassword: validation.FieldRules{
			"Password":     "max=5000,required",
			"PasswordSalt": "max=5000,required",
		},
		constants.ScenarioForgotPassword: validation.FieldRules{
			"PasswordResetToken":          "max=5000,required",
			"PasswordResetTokenExpiresAt": "required,customFutureValidator",
		},
		constants.ScenarioChangePassword: validation.FieldRules{
			"Password": "max=5000,required,customPasswordValidator",
		},
		constants.ScenarioResetPassword: validation.FieldRules{
			"Password": "max=5000,required,customPasswordValidator",
		},
		constants.ScenarioVerifyEmail: validation.FieldRules{
			"IsEmailVerified": "eq=true",
			"EmailTwoFaCode":  "eq=",
		},
		constants.ScenarioLoginTwoFaStepOne: validation.FieldRules{
			"EmailTwoFaCode": "required,min=6,max=6",
		},
	}
}

func (User) GetValidator() interface{} {
	v := validator.New()
	err := v.RegisterValidation("customPasswordValidator", validation.CustomPasswordValidator)
	if err != nil {
		helpers.LogError(err)
		return nil
	}
	err = v.RegisterValidation("customFutureValidator", validation.CustomFutureValidator)
	if err != nil {
		helpers.LogError(err)
		return nil
	}
	return v
}

func InsertUser(db *gorm.DB, m *User) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioCreate, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	m.EncodePassword()
	err = validation.ValidateByScenario(constants.ScenarioHashPassword, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	err = db.Create(m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func ForgotPassword(db *gorm.DB, m *User) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioForgotPassword, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	sql := "UPDATE users SET password_reset_token = ?, password_reset_token_expire_at = ? WHERE id = ?"
	return db.Exec(sql, m.PasswordResetToken, m.PasswordResetTokenExpiresAt, m.Id).Error
}

func SetEmailTwoFaCode(db *gorm.DB, m *User) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioLoginTwoFaStepOne, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	sql := "UPDATE users SET email_two_fa_code = ? WHERE id = ?"
	return db.Exec(sql, m.EmailTwoFaCode, m.Id).Error
}

func ResetEmailTwoFaCode(db *gorm.DB, m *User) (err error) {
	sql := "UPDATE users SET email_two_fa_code = '' WHERE id = ?"
	return db.Exec(sql, m.Id).Error
}

func ResetResetPasswordToken(db *gorm.DB, m *User) (err error) {
	sql := "UPDATE users SET password_reset_token = '', password_reset_token_expire_at = NULL WHERE id = ?"
	return db.Exec(sql, m.Id).Error
}

func SetUserEmailVerified(db *gorm.DB, m *User) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioVerifyEmail, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	sql := "UPDATE users SET is_email_verified = ?, email_two_fa_code = ? WHERE id = ?"
	return db.Exec(sql, m.IsEmailVerified, m.EmailTwoFaCode, m.Id).Error
}

func UserResetPassword(db *gorm.DB, m *User) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioResetPassword, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	m.EncodePassword()
	err = validation.ValidateByScenario(constants.ScenarioHashPassword, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	sql := "UPDATE users SET password = ?, password_salt = ? WHERE id = ?"
	return db.Exec(sql, m.Password, m.PasswordSalt, m.Id).Error
}

func UserChangePassword(db *gorm.DB, m *User) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioChangePassword, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	m.EncodePassword()
	err = validation.ValidateByScenario(constants.ScenarioHashPassword, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	sql := "UPDATE users SET password = ?, password_salt = ? WHERE id = ?"
	return db.Exec(sql, m.Password, m.PasswordSalt, m.Id).Error
}

func (m *User) EncodePassword() {
	salt, encodedPwd := password.Encode(m.Password, nil)
	m.Password = encodedPwd
	m.PasswordSalt = salt
}

func FindUserById(db *gorm.DB, id int) (*User, error) {
	m := User{}
	err := db.First(&m, id).Error
	if err != nil {
		helpers.LogError(err)
	}
	return &m, err
}

func FindUserByTwoFAToken(db *gorm.DB, token string) (*User, error) {
	m := User{}
	err := db.Where("email_two_fa_code = ?", token).First(&m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return &m, err
}

func FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	m := User{}
	err := db.Where("email = ?", email).First(&m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return &m, err
}

func FindUserByResetPasswordToken(db *gorm.DB, token string) (*User, error) {
	m := User{}
	err := db.
		Where("password_reset_token = ?", token).
		Where("password_reset_token_expire_at > ?", time.Now().Unix()).
		First(&m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return &m, err
}

func (u *User) SetForgotPasswordData() string {
	token := helpers.GenerateRandomString(32)
	currentTime := time.Now()
	oneHourLater := currentTime.Add(time.Hour)

	u.PasswordResetToken = token
	u.PasswordResetTokenExpiresAt = &oneHourLater

	return token
}

package models

import (
	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id                         int        `json:"id" column:"id" gorm:"primaryKey;autoIncrement:true"`
	Email                      string     `json:"email" column:"email"`
	Nickname                   string     `json:"nickname" column:"nickname"`
	Password                   string     `json:"password" column:"password"`
	PasswordSalt               string     `column:"password_salt"`
	PasswordResetToken         string     `column:"password_reset_token"`
	PasswordResetTokenExpireAt *time.Time `column:"password_reset_token_expire_at"`
	Role                       int        `json:"role" column:"role"`
	IsEmailVerified            bool       `json:"is_email_verified" column:"is_email_verified"`
	EmailTwoFaCode             string     `json:"email_two_fa_code" column:"email_two_fa_code"`
}

func (m *User) TableName() string {
	return "user"
}

func (User) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioCreate: validation.FieldRules{
			"Email":          "min=3,max=255,email,required",
			"Nickname":       "min=3,max=255,required",
			"Password":       "min=8,max=5000,required,customPasswordValidator",
			"Role":           "required,gt=0,lt=2", // we can create only users, admin should be created separately
			"EmailTwoFaCode": "required",
		},
		constants.ScenarioHashPassword: validation.FieldRules{
			"Password":     "min=8,max=5000,required",
			"PasswordSalt": "min=3,max=5000,required",
		},
		constants.ScenarioForgotPassword: validation.FieldRules{
			"PasswordResetToken":         "min=3,max=5000,required",
			"PasswordResetTokenExpireAt": "required,customFutureValidator",
		},
		constants.ScenarioChangePassword: validation.FieldRules{
			"Password": "min=8,max=5000,required,customPasswordValidator",
		},
		constants.ScenarioResetPassword: validation.FieldRules{
			"Password": "min=8,max=5000,required,customPasswordValidator",
		},
		constants.ScenarioVerifyEmail: validation.FieldRules{
			"IsEmailVerified": "eq=true",
			"EmailTwoFaCode":  "eq=",
		},
		constants.ScenarioLoginTwoFaStepOne: validation.FieldRules{
			"EmailTwoFaCode": "min=6,max=6,required",
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
	m.encodePassword()
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
	sql := "UPDATE user SET password_reset_token = ?, password_reset_token_expire_at = ? WHERE id = ?"
	return db.Exec(sql, m.PasswordResetToken, m.PasswordResetTokenExpireAt, m.Id).Error
}

func SetEmailTwoFaCode(db *gorm.DB, m *User) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioLoginTwoFaStepOne, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	sql := "UPDATE user SET email_two_fa_code = ? WHERE id = ?"
	return db.Exec(sql, m.EmailTwoFaCode, m.Id).Error
}

func ResetEmailTwoFaCode(db *gorm.DB, m *User) (err error) {
	sql := "UPDATE user SET email_two_fa_code = '' WHERE id = ?"
	return db.Exec(sql, m.Id).Error
}

func ResetResetPasswordToken(db *gorm.DB, m *User) (err error) {
	sql := "UPDATE user SET password_reset_token = '', password_reset_token_expire_at = NULL WHERE id = ?"
	return db.Exec(sql, m.Id).Error
}

func SetUserEmailVerified(db *gorm.DB, m *User) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioVerifyEmail, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	sql := "UPDATE user SET is_email_verified = ?, email_two_fa_code = ? WHERE id = ?"
	return db.Exec(sql, m.IsEmailVerified, m.EmailTwoFaCode, m.Id).Error
}

func UserResetPassword(db *gorm.DB, m *User) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioResetPassword, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	m.encodePassword()
	err = validation.ValidateByScenario(constants.ScenarioHashPassword, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	sql := "UPDATE user SET password = ?, password_salt = ? WHERE id = ?"
	return db.Exec(sql, m.Password, m.PasswordSalt, m.Id).Error
}

func UserChangePassword(db *gorm.DB, m *User) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioChangePassword, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	m.encodePassword()
	err = validation.ValidateByScenario(constants.ScenarioHashPassword, *m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	sql := "UPDATE user SET password = ?, password_salt = ? WHERE id = ?"
	return db.Exec(sql, m.Password, m.PasswordSalt, m.Id).Error
}

func (m *User) encodePassword() {
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

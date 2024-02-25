package constants

import (
	"errors"
	"time"
)

var (
	ServerError       = errors.New("Server Error")
	BadRequestError   = errors.New("Bad Request")
	UnauthorizedError = errors.New("Unauthorized")
)

const (
	RoleUser  = 1
	RoleAdmin = 2

	GenderMale   = "male"
	GenderFemale = "female"
	GenderOther  = "other"

	// email
	TwoFaLoginSubject               = "beenny.com: 2FA login verification"
	EmailVerificationSubject        = "beenny.com: email verification"
	ResetPasswordSubject            = "beenny.com: reset forgotten password"
	ResetPasswordHtmlBodyFormat     = "In order to reset your password, please forward this link %s"
	EmailVerificationHtmlBodyFormat = "Please, verify your email by forwarding this link <a href='%s'>Verify Email</a>"
	LoginTwoFaCodeHtmlBodyFormat    = "Please, forward this link in order to login <a href='%s'>Login</a>"
	EmailCharSet                    = "UTF-8"

	// TODO: replace with real domen
	NoReplySenderEmail = "no-reply@beenny.com"

	//common
	SqlDsnFormat = `%s:%s@tcp(%s:%s)/%s`

	MigrationsFolder = "migrations"

	// validation tags
	RequiredTag                = "required"
	MinTag                     = "min"
	MaxTag                     = "max"
	EmailTag                   = "email"
	GreaterThanTag             = "gt"
	LowerThanTag               = "lt"
	EqTag                      = "eq"
	FutureTag                  = "customFutureValidator"
	CustomPasswordValidatorTag = "customPasswordValidator"

	// validation error messages
	FutureErrorMsg                     = "'%s' should be in the future"
	EqErrorMsg                         = "'%s' should be %s"
	RequiredErrorMsg                   = "'%s' is required"
	MinValueErrorMsg                   = "'%s' min value is %s"
	MaxValueErrorMsg                   = "'%s' max value is %s"
	EmailErrorMsg                      = "'Email' is not valid"
	GreaterThanTagErrorMsg             = "value should be greater than %s"
	LowerThanTagErrorMsg               = "value should be lower than %s"
	CustomPasswordValidatorTagErrorMsg = "'%s' must have: 1 small letter, 1 big letter, 1 special symbol, 1 digit and be at least 8 symbols long"

	// Server
	DefaultWriteTimout  = 60 * time.Second
	DefaultReadTimeout  = 60 * time.Second
	DefaultStoreTimeout = 60 * time.Second

	// scenarios
	ScenarioCreate            = "create"
	ScenarioUpdate            = "update"
	ScenarioDelete            = "delete"
	ScenarioSignUp            = "sign-up"
	ScenarioSignIn            = "sign-in"
	ScenarioHashPassword      = "hash-password"
	ScenarioForgotPassword    = "forgot-password"
	ScenarioRegister          = "register"
	ScenarioLoginTwoFaStepOne = "login-two-fa-step-one"
	ScenarioChangePassword    = "change-password"
	ScenarioResetPassword     = "reset-password"
	ScenarioVerifyEmail       = "verify-email"
	ScenarioSetUserPhoto      = "set-user-photo"
	ScenarioTwoFaLoginStepOne = "two-fa-login-step-one"
	ScenarioTwoFaLoginStepTwo = "two-fa-login-step-two"

	RegisterUserURL = "/api/v1/security/register"
)

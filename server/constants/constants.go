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

	FreeBillingPlan     = 1
	StandardBillingPlan = 2
	ProBillingPlan      = 3

	// email
	// TODO: replace with real domain
	TwoFaLoginSubject               = "Benny.com: 2FA login verification"
	EmailVerificationSubject        = "Benny.com: email verification"
	ResetPasswordSubject            = "Benny.com: reset forgotten password"
	ResetPasswordHtmlBodyFormat     = "In order to reset your password, please forward this link %s"
	EmailVerificationHtmlBodyFormat = "Please, verify your email by forwarding this link <a href='%s'>Verify Email</a>"
	LoginTwoFaCodeHtmlBodyFormat    = "Please, forward this link in order to login <a href='%s'>Login</a>"
	EmailCharSet                    = "UTF-8"

	// TODO: replace with real domen
	NoReplySenderEmail = "no-reply@beenny.com"

	//common
	SqlDsnFormat = `%s:%s@tcp(%s:%s)/%s`

	DevelopmentAppEnv = "development"
	ProductionAppEnv  = "production"
	TestingAppEnv     = "testing"

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
	FutureErrorMsg                     = "%s resource: '%s' should be in the future"
	EqErrorMsg                         = "%s resource: '%s' should be %s"
	RequiredErrorMsg                   = "%s resource: '%s' is required"
	MinValueErrorMsg                   = "%s resource: '%s' min value is %s"
	MaxValueErrorMsg                   = "%s resource: '%s' max value is %s"
	EmailErrorMsg                      = "%s resource: email is not valid"
	GreaterThanTagErrorMsg             = "%s resource: value should be greater than %s"
	LowerThanTagErrorMsg               = "%s resource: value should be lower than %s"
	CustomPasswordValidatorTagErrorMsg = "%s resource: password must have: 1 small letter, 1 big letter, 1 special symbol"

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
	ScenarioTwoFaLoginStepOne = "two-fa-login-step-one"
	ScenarioTwoFaLoginStepTwo = "two-fa-login-step-two"
)

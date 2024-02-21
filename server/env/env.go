package env

import "os"

func IsTestEnv() bool {
	return os.Getenv("APP_ENV") == "test"
}

func IsDevelopmentEnv() bool {
	return os.Getenv("APP_ENV") == "development"
}

package env

import "os"

func IsTestEnv() bool {
	return os.Getenv("APP_ENV") == "test"
}

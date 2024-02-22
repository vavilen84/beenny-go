package handlers_test

import (
	"testing"
)

func Test_Integration_Security_Register_OK(t *testing.T) {
	ts := initApp()
	defer ts.Close()
	registerUser(t, ts)
}

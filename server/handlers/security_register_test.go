package handlers_test

import (
	"testing"
)

func TestRegister_OK(t *testing.T) {
	ts := initApp()
	defer ts.Close()
	registerUser(t, ts)
}

package mocks

import (
	"errors"
	"strings"
)

type CustomMatcher struct{}

func (c CustomMatcher) Match(expectedSQL, actualSQL string) error {
	if !strings.Contains(actualSQL, expectedSQL) {
		return errors.New("SQL doesnt match")
	}
	return nil
}

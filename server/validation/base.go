package validation

import (
	"reflect"
)

type Scenario string

type Rules string

type ScenarioRules map[Scenario]FieldRules

type FieldRules map[string]Rules

func getType(s interface{}) string {
	if t := reflect.TypeOf(s); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

package validation

import (
	"fmt"
	"github.com/vavilen84/nft-project/constants"
)

type FieldError struct {
	Tag     string
	Field   string
	Message string
	Value   string
	Param   string
	Name    string
}

func (s *FieldError) setErrorMessage() {
	switch s.Tag {
	case constants.RequiredTag:
		s.Message = fmt.Sprintf(constants.RequiredErrorMsg, s.Name, s.Field)
	case constants.EqTag:
		s.Message = fmt.Sprintf(constants.EqErrorMsg, s.Name, s.Field, s.Param)
	case constants.FutureTag:
		s.Message = fmt.Sprintf(constants.FutureErrorMsg, s.Name, s.Field)
	case constants.MinTag:
		s.Message = fmt.Sprintf(constants.MinValueErrorMsg, s.Name, s.Field, s.Param)
	case constants.MaxTag:
		s.Message = fmt.Sprintf(constants.MaxValueErrorMsg, s.Name, s.Field, s.Param)
	case constants.EmailTag:
		s.Message = fmt.Sprintf(constants.EmailErrorMsg, s.Name)
	case constants.GreaterThanTag:
		s.Message = fmt.Sprintf(constants.GreaterThanTagErrorMsg, s.Name, s.Param)
	case constants.LowerThanTag:
		s.Message = fmt.Sprintf(constants.LowerThanTagErrorMsg, s.Name, s.Param)
	case constants.CustomPasswordValidatorTag:
		s.Message = fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, s.Name)
	default:
		s.Message = "Undefined validation error"
	}
}

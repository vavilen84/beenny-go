package interfaces

type Model interface {
	GetValidator() interface{}
	GetValidationRules() interface{}
}

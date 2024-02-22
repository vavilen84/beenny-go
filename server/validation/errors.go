package validation

type Errors []error

func (s Errors) Error() (result string) {
	for _, e := range s {
		result += e.Error() + ";\n"
	}
	return
}

package validation

type Errors []string

func (s Errors) Error() (result string) {
	for _, e := range s {
		result += e + ";\n"
	}
	return
}

package test

type SesError struct {
	error
	OrgigErrorData string
	MessageData    string
	CodeData       string
}

func (s SesError) Code() string {
	return s.CodeData
}

func (s SesError) Message() string {
	return s.MessageData
}

func (s SesError) OrigErr() error {
	return nil
}

func (s SesError) Error() string {
	return s.Message()
}

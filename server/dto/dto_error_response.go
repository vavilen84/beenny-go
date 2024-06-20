package dto

type ErrorResponse struct {
	Error  string   `json:"error"`
	Errors []string `json:"errors"`
}

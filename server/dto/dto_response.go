package dto

type Response struct {
	Data   []byte   `json:"data"`
	Errors []string `json:"errors"`
}

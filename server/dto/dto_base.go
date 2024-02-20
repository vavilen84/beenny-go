package dto

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
	Errors []string    `json:"errors"`

	FormErrors map[string][]string `json:"formErrors"`
}

type ResponseData map[string]interface{}

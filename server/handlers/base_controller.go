package handlers

import (
	"github.com/vavilen84/beenny-go/dto"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"net/http"
)

type BaseController struct{}

func (*BaseController) WriteSuccessResponse(w http.ResponseWriter, data []byte, status int) {
	resp := dto.Response{
		Data: data,
	}
	helpers.WriteResponse(w, helpers.MarshalGeneric(resp), status)
}

func (*BaseController) WriteErrorResponse(w http.ResponseWriter, err interface{}, status int) {
	resp := dto.Response{}
	e, ok := err.(error)
	if ok {
		helpers.LogError(e)
	}
	ok = false
	errorsSlice := make([]string, 0)
	errs, ok := err.(validation.Errors)
	if ok {
		resp = dto.Response{
			Errors: errs,
		}
	} else {
		errorsSlice = append(errorsSlice, e.Error())
		resp = dto.Response{
			Errors: errorsSlice,
		}
	}
	helpers.WriteResponse(w, helpers.MarshalGeneric(resp), status)
}

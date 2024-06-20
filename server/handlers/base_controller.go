package handlers

import (
	"github.com/vavilen84/beenny-go/dto"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/validation"
	"net/http"
)

type BaseController struct{}

func (*BaseController) WriteSuccessResponse(w http.ResponseWriter, data []byte, status int) {
	helpers.WriteResponse(w, data, status)
}

func (*BaseController) WriteErrorResponse(w http.ResponseWriter, i interface{}, status int) {
	resp := dto.ErrorResponse{}
	e, ok := i.(error)
	if ok {
		helpers.LogError(e)
		resp = dto.ErrorResponse{
			Error: e.Error(),
		}
		helpers.WriteResponse(w, helpers.MarshalGeneric(resp), status)
		return
	}
	errs, ok := i.(validation.Errors)
	if ok {
		resp = dto.ErrorResponse{
			Errors: errs,
		}
	}
	helpers.WriteResponse(w, helpers.MarshalGeneric(resp), status)
}

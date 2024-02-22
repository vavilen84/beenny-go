package handlers

import (
	"encoding/json"
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
	writeResponse(w, resp, status)
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
	writeResponse(w, resp, status)
}

func writeResponse(w http.ResponseWriter, resp dto.Response, status int) {
	b, e := json.Marshal(resp)
	if e != nil {
		helpers.LogError(e)
		return
	}
	setCacheHeaders(w)
	setContentTypeHeader(w)
	w.WriteHeader(status)
	_, err := w.Write(b)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

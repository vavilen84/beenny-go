package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vavilen84/beenny-go/aws"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/dto"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"github.com/vavilen84/beenny-go/validation"
	"gorm.io/gorm"
	"net/http"
)

func (c *SecurityController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dtoModel := dto.ForgotPassword{}
	err := dec.Decode(&dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	err = validation.ValidateByScenario(constants.ScenarioForgotPassword, dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	db := store.GetDB()
	u, err := models.FindUserByEmail(db, dtoModel.Email)
	if err != nil {
		helpers.LogError(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(fmt.Sprintf("user with email %s not found", dtoModel.Email))
			c.WriteErrorResponse(w, err, http.StatusNotFound)
		} else {
			c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		}
		return
	}
	token := u.SetForgotPasswordData()
	err = models.ForgotPassword(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}

	err = aws.SendResetPasswordEmail(u.Email, token)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}
	c.WriteSuccessResponse(w, bytes, http.StatusOK)
}

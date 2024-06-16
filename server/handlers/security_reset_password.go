package handlers

import (
	"encoding/json"
	"errors"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/dto"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"github.com/vavilen84/beenny-go/validation"
	"gorm.io/gorm"
	"net/http"
)

func (c *SecurityController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dtoModel := dto.ResetPassword{}
	err := dec.Decode(&dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	err = validation.ValidateByScenario(constants.ScenarioResetPassword, dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	db := store.GetDB()
	u, err := models.FindUserByResetPasswordToken(db, dtoModel.Token)
	if err != nil {
		helpers.LogError(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("user not found")
			c.WriteErrorResponse(w, err, http.StatusNotFound)
		} else {
			c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		}
		return
	}
	u.Password = dtoModel.NewPassword
	err = models.UserResetPassword(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}
	err = models.ResetResetPasswordToken(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
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

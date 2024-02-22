package handlers

import (
	"encoding/json"
	"errors"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/vavilen84/beenny-go/constants"
	"github.com/vavilen84/beenny-go/dto"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"github.com/vavilen84/beenny-go/validation"
	"net/http"
)

func (c *SecurityController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	db := store.GetDB()

	dec := json.NewDecoder(r.Body)
	dtoModel := dto.ChangePassword{}
	err := dec.Decode(&dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	err = validation.ValidateByScenario(constants.ScenarioChangePassword, dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}

	u, ok := r.Context().Value("user").(*models.User)
	if !ok {
		err = errors.New("No logged in user")
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.UnauthorizedError, http.StatusUnauthorized)
		return
	}

	passwordIsValid := password.Verify(dtoModel.OldPassword, u.PasswordSalt, u.Password, nil)
	if !passwordIsValid {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.UnauthorizedError, http.StatusUnauthorized)
		return
	}
	u.Password = dtoModel.NewPassword
	err = models.UserChangePassword(db, u)
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

package handlers

import (
	"encoding/json"
	"errors"
	"github.com/vavilen84/nft-project/auth"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/gorm"
	"net/http"
)

func (c *SecurityController) TwoFaLoginStepTwo(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dtoModel := dto.TwoFaLoginStepTwo{}
	err := dec.Decode(&dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	err = validation.ValidateByScenario(constants.ScenarioTwoFaLoginStepTwo, dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	db := store.GetDB()
	u, err := models.FindUserByTwoFAToken(db, dtoModel.EmailTwoFaCode)
	if err != nil {
		helpers.LogError(err)
		if err == gorm.ErrRecordNotFound {
			err = errors.New("user not found")
			c.WriteErrorResponse(w, err, http.StatusNotFound)
		} else {
			c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		}
		return
	}

	u.EmailTwoFaCode = ""
	err = models.ResetEmailTwoFaCode(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}

	tok, err := auth.CreateJWT(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	resp := make(dto.ResponseData)
	resp["token"] = string(tok)

	c.WriteSuccessResponse(w, resp, http.StatusOK)
}

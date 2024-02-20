package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vavilen84/nft-project/auth"
	"github.com/vavilen84/nft-project/aws"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/gorm"
	"net/http"
)

func (c *SecurityController) Register(w http.ResponseWriter, r *http.Request) {
	db := store.GetDB()
	dec := json.NewDecoder(r.Body)
	dtoModel := dto.Register{}
	err := dec.Decode(&dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	err = validation.ValidateByScenario(constants.ScenarioRegister, dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	u, err := models.FindUserByEmail(db, dtoModel.Email)
	if err != nil {
		helpers.LogError(err)
		if err != gorm.ErrRecordNotFound {
			c.WriteErrorResponse(w, err, http.StatusInternalServerError)
			return
		}
	} else {
		err := errors.New(fmt.Sprintf("user with email %s already exists", dtoModel.Email))
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	emailVerificationToken := helpers.GenerateRandomString(6)
	u = &models.User{
		Email:           dtoModel.Email,
		Password:        dtoModel.Password,
		Nickname:        dtoModel.Nickname,
		Role:            constants.RoleUser,
		IsEmailVerified: false,
		EmailTwoFaCode:  emailVerificationToken,
	}

	err = models.InsertUser(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = aws.SendEmailVerificationMail(u.Email, emailVerificationToken)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	_, err = auth.CreateJWT(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	resp := make(dto.ResponseData)
	c.WriteSuccessResponse(w, resp, http.StatusOK)
}

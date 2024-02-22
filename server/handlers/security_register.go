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
		c.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	if dtoModel.Password != dtoModel.ConfirmPassword {
		err = errors.New("Passwords don't match")
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	u, err := models.FindUserByEmail(db, dtoModel.Email)
	if err != nil {
		helpers.LogError(err)
		if err != gorm.ErrRecordNotFound {
			c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
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
		FirstName:       dtoModel.FirstName,
		LastName:        dtoModel.LastName,
		Email:           dtoModel.Email,
		CurrentCountry:  dtoModel.CurrentCountry,
		CountryOfBirth:  dtoModel.CountryOfBirth,
		Gender:          dtoModel.Gender,
		Timezone:        dtoModel.Timezone,
		Birthday:        dtoModel.Birthday,
		Password:        dtoModel.Password,
		Role:            constants.RoleUser,
		Photo:           dtoModel.Photo,
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
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}

	resp := make(dto.ResponseData)
	bytes, err := json.Marshal(u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}
	resp["user"] = bytes
	c.WriteSuccessResponse(w, resp, http.StatusOK)
}

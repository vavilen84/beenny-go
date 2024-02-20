package handlers

import (
	"errors"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"gorm.io/gorm"
	"net/http"
)

func (c *SecurityController) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	token := params.Get("token")

	db := store.GetDB()
	u, err := models.FindUserByTwoFAToken(db, token)
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
	u.IsEmailVerified = true
	u.EmailTwoFaCode = ""
	err = models.SetUserEmailVerified(db, u)
	if err != nil {
		helpers.LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	resp := make(dto.ResponseData)
	c.WriteSuccessResponse(w, resp, http.StatusOK)
}

package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/vavilen84/nft-project/auth"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"gorm.io/gorm"
	"net/http"
)

func UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := store.GetDB()
		token := r.Header.Get("Authorization")
		isValid, err := auth.VerifyJWT(db, []byte(token))
		if err != nil || token == "" || !isValid {
			helpers.LogError(err)
			http.Error(w, http.StatusText(401), 401)
			return
		}

		p, err := auth.ParseJWTPayload([]byte(token))
		if err != nil {
			helpers.LogError(err)
			http.Error(w, http.StatusText(401), 401)
			return
		}

		jwtInfo, err := models.FindJWTInfoById(db, p.JWTInfoId)
		if err != nil {
			helpers.LogError(err)
			http.Error(w, http.StatusText(401), 401)
			return
		}

		u, err := models.FindUserById(db, jwtInfo.UserId)
		if err != nil {
			helpers.LogError(err)
			if err == gorm.ErrRecordNotFound {
				err = errors.New(fmt.Sprintf("user with email %s not found", jwtInfo.User.Email))
				helpers.LogError(err)
				http.Error(w, http.StatusText(404), 404)
				return
			} else {
				http.Error(w, http.StatusText(500), 500)
				return
			}
		}
		ctx := context.WithValue(r.Context(), "user", u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

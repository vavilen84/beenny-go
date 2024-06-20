package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/vavilen84/beenny-go/auth"
	"github.com/vavilen84/beenny-go/dto"
	"github.com/vavilen84/beenny-go/helpers"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"gorm.io/gorm"
	"net/http"
)

func writeResponse(w http.ResponseWriter, status int, payload string) {
	resp := dto.ErrorResponse{
		Error: payload,
	}
	helpers.WriteResponse(w, helpers.MarshalGeneric(resp), status)
}

func UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := store.GetDB()
		token := r.Header.Get("Authorization")
		isValid, err := auth.VerifyJWT(db, []byte(token))
		if err != nil || token == "" || !isValid {
			helpers.LogError(err)
			writeResponse(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		p, err := auth.ParseJWTPayload([]byte(token))
		if err != nil {
			helpers.LogError(err)
			writeResponse(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		jwtInfo, err := models.FindJWTInfoById(db, p.JWTInfoId)
		if err != nil {
			helpers.LogError(err)
			writeResponse(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		u, err := models.FindUserById(db, jwtInfo.UserId)
		if err != nil {
			helpers.LogError(err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helpers.LogError(err)
				errMsg := fmt.Sprintf("user with email %s not found", jwtInfo.User.Email)
				writeResponse(w, http.StatusNotFound, errMsg)
				return
			} else {
				writeResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}
		ctx := context.WithValue(r.Context(), "user", u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

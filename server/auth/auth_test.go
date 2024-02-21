package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/models"
	"log"
	"testing"
)

func TestParseJWTToken(t *testing.T) {
	tok := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE4MjMzMjgsImlhdCI6MTY5MDI4NzMyOCwiand0X2luZm9faWQiOjE1fQ.beGtWScxnFaBut5LJ2HIX61dtog_y6gdxpOskeHGAoU"
	jwtPayload, err := ParseJWTPayload([]byte(tok))
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, jwtPayload.JWTInfoId, 15)
	assert.NotEmpty(t, jwtPayload.Payload.ExpirationTime)
	assert.NotEmpty(t, jwtPayload.Payload.IssuedAt)
}

func TestGenerateJWTAndParse(t *testing.T) {
	u := models.User{
		Id: 15,
	}
	jwtInfo := getJWTInfo(&u)
	jwtInfo.Id = 15 // like we have inserted it
	jwtInfo.GenerateSecret()
	token, err := generateJWT(jwtInfo)
	assert.Nil(t, err)
	jwtPayload, err := ParseJWTPayload(token)
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, jwtPayload.JWTInfoId, 15)
	assert.NotEmpty(t, jwtPayload.Payload.ExpirationTime)
	assert.NotEmpty(t, jwtPayload.Payload.IssuedAt)
}

package auth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestParseJWTToken(t *testing.T) {
	tok := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE4MjMzMjgsImlhdCI6MTY5MDI4NzMyOCwiand0X2luZm9faWQiOjE1fQ.beGtWScxnFaBut5LJ2HIX61dtog_y6gdxpOskeHGAoU"
	jwtPayload, err := ParseJWTPayload([]byte(tok))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(jwtPayload)
	assert.Equal(t, jwtPayload.JWTInfoId, 15)
	assert.NotEmpty(t, jwtPayload.Payload.ExpirationTime)
	assert.NotEmpty(t, jwtPayload.Payload.IssuedAt)
}

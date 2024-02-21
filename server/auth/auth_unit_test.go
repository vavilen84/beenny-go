package auth

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/mocks"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
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

func Test_getJWTInfo(t *testing.T) {
	m := models.User{
		Id: 15,
	}
	jwtInfo := getJWTInfo(&m)
	assert.Equal(t, 15, jwtInfo.UserId)
	assert.NotNil(t, jwtInfo.ExpiresAt)
}

func Test_insertJWTInfo_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "INSERT INTO `jwt_info`"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(2, 1))

	m := models.User{
		Id: 15,
	}
	jwtInfo, err := insertJWTInfo(gormDB, &m)
	assert.Nil(t, err)
	assert.Equal(t, 2, jwtInfo.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_CreateJWT_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	sql := "INSERT INTO `jwt_info`"
	sqlMock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(2, 1))

	m := models.User{
		Id: 15,
	}
	token, err := CreateJWT(gormDB, &m)
	assert.Nil(t, err)
	jwtPayload, err := ParseJWTPayload(token)
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, 2, jwtPayload.JWTInfoId)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

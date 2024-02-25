package auth

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/beenny-go/mocks"
	"github.com/vavilen84/beenny-go/models"
	"github.com/vavilen84/beenny-go/store"
	"log"
	"testing"
	"time"
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

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	expectedSQL := "SELECT * FROM `users`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

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

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	expectedSQL := "SELECT * FROM `users`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

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

func Test_VerifyJWT_ok(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	m := models.User{
		Id: 15,
	}
	jwtInfo := getJWTInfo(&m)
	jwtInfo.GenerateSecret()
	token, err := generateJWT(jwtInfo)
	assert.Nil(t, err)

	rows := sqlmock.NewRows([]string{"id", "secret"}).AddRow(2, jwtInfo.Secret)
	expectedSQL := "SELECT * FROM `jwt_info`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	isValid, err := VerifyJWT(gormDB, token)
	assert.Nil(t, err)
	assert.True(t, isValid)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_VerifyJWT_notOk_couldNotParseJWTPayload(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	token := []byte("123sd-f09sdf0-9sd-0f9s-d0f9s-d0f9-sd0f9s-d0f9sd-0f9") // fake token

	isValid, err := VerifyJWT(gormDB, token)
	assert.NotNil(t, err)
	assert.False(t, isValid)
}

func Test_VerifyJWT_notOk_couldNotFindJWTInfoById(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	m := models.User{
		Id: 15,
	}
	jwtInfo := getJWTInfo(&m)
	jwtInfo.GenerateSecret()
	token, err := generateJWT(jwtInfo)
	assert.Nil(t, err)

	rows := sqlmock.NewRows([]string{"id", "secret"}) // no rows returned
	expectedSQL := "SELECT * FROM `jwt_info`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	isValid, err := VerifyJWT(gormDB, token)
	assert.NotNil(t, err)
	assert.False(t, isValid)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_VerifyJWT_notOk_wrongSecret(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	m := models.User{
		Id: 15,
	}
	jwtInfo := getJWTInfo(&m)
	jwtInfo.GenerateSecret()
	token, err := generateJWT(jwtInfo)
	assert.Nil(t, err)

	rows := sqlmock.NewRows([]string{"id", "secret"}).AddRow(2, "123456") // wrong secret
	expectedSQL := "SELECT * FROM `jwt_info`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	isValid, err := VerifyJWT(gormDB, token)
	assert.NotNil(t, err)
	assert.False(t, isValid)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_VerifyJWT_notOk_expired(t *testing.T) {
	customMatcher := mocks.CustomMatcher{}
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	gormDB := store.GetMockDB(db)

	m := models.User{
		Id: 15,
	}
	jwtInfo := getJWTInfo(&m)
	jwtInfo.ExpiresAt = time.Now().Add(-1 * time.Hour) // already expired
	jwtInfo.GenerateSecret()
	token, err := generateJWT(jwtInfo)
	assert.Nil(t, err)

	rows := sqlmock.NewRows([]string{"id", "secret"}).AddRow(2, jwtInfo.Secret)
	expectedSQL := "SELECT * FROM `jwt_info`"
	sqlMock.ExpectQuery(expectedSQL).WillReturnRows(rows)

	isValid, err := VerifyJWT(gormDB, token)
	assert.NotNil(t, err)
	assert.False(t, isValid)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

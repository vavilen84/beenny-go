package test

import (
	"fmt"
	"github.com/vavilen84/beenny-go/handlers"
	"github.com/vavilen84/beenny-go/store"
	"net/http/httptest"
)

// for integration tests
func InitTestApp() *httptest.Server {
	LoadEnv()
	store.InitTestDB()
	db := store.GetDB()
	if err := db.Exec("delete from `jwt_info`").Error; err != nil {
		fmt.Println("Error deleting entities:", err)
	}
	if err := db.Exec("delete from `users`").Error; err != nil {
		fmt.Println("Error deleting entities:", err)
	}
	return MakeTestServer()
}

func MakeTestServer() *httptest.Server {
	handler := handlers.MakeHandler()
	ts := httptest.NewServer(handler)
	return ts
}

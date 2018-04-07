package auth

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/gopok/gopok-backend/pkg/core"
)

func marshalJSONNoError(data interface{}) []byte {
	out, _ := json.Marshal(data)
	return out
}

func createTestUser(app *core.Application) string {
	userUsername := "test_user_" + strconv.Itoa(rand.Intn(10000))

	createUserReq, _ := http.NewRequest("POST", "/api/auth/users", bytes.NewBuffer(marshalJSONNoError(map[string]string{
		"username": userUsername,
		"password": "test",
		"email":    "test@test.pl",
	})))
	createUserReq.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, createUserReq)
	if rr.Code != 200 {
		panic("Failed to create test user" + rr.Body.String())
	}

	return userUsername
}

func loginWithTestUser(app *core.Application, username string) string {
	loginReq, _ := http.NewRequest("POST", "/api/auth/sessions/login", bytes.NewBuffer(marshalJSONNoError(map[string]string{
		"username": username,
		"password": "test",
		"email":    "test@test.pl",
	})))
	loginReq.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, loginReq)
	if rr.Code != 200 {
		panic("Failed to login with test user" + rr.Body.String())
	}
	var retData map[string]interface{}
	dec := json.NewDecoder(rr.Body)
	err := dec.Decode(&retData)
	if err != nil {
		panic(err)
	}
	return retData["token"].(string)
}

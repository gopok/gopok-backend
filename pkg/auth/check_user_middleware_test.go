package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopok/gopok-backend/pkg/core"
)

func TestExtractsTokenFromRequest(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test", nil)
	r.Header.Set("Authorization", "Bearer fuhfdchu76e365bhvdfhgdsfg26re32fyFT")
	token, err := extractSessionToken(r)
	if err != nil || token != "fuhfdchu76e365bhvdfhgdsfg26re32fyFT" {
		t.Errorf("Token wasn't extracted correctly")
		t.Fail()
		return
	}

	r2, _ := http.NewRequest("GET", "/test", nil)
	_, shouldErr := extractSessionToken(r2)
	if shouldErr == nil {
		t.Errorf("Expected an error when no header is set")
		t.Fail()
		return
	}
}

func TestCheckUserMiddleware(t *testing.T) {

	app := &core.Application{}
	app.Config = &core.CoreConfiguration{
		HTTPPort: 3002,
		MongoURL: "mongodb://localhost/go_test_TestCheckUserMiddleware",
	}
	app.Prepare()
	uc := UsersController{}
	uc.Register(app)
	sc := SessionsController{}
	sc.Register(app)

	username := createTestUser(app)
	sessionToken := loginWithTestUser(app, username)
	rr := httptest.NewRecorder()

	authedRequest, err := http.NewRequest("GET", "/abvfyhdhfd", nil)
	if err != nil {
		panic(err)
	}
	authedRequest.Header.Set("Authorization", "Bearer "+sessionToken)

	mid := CheckUserMiddleware(app)
	var called bool
	mid(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		if r.Context().Value(UserContextKey).(*User).Username != username {
			t.Error("Users mismatch")
			t.Fail()
			return
		}
	})).ServeHTTP(rr, authedRequest)
	if !called {
		t.Error("Handler wasn't called by CheckUserMiddleware")
	}

}

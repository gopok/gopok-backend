package auth

import (
	"encoding/json"
	"net/http"
	"testing"
)

func marshalJSONNoError(data interface{}) []byte {
	out, _ := json.Marshal(data)
	return out
}

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
	/*userUsername := "test_user_" + string(time.Now().UnixNano())

	app := &core.Application{}
	app.DbOverride = "go_test_db_TestCheckUserMiddleware"
	app.Prepare()

	createUserReq, _ := http.NewRequest("POST", "/api/auth/users", bytes.NewBuffer(marshalJSONNoError(map[string]string{
		"username": userUsername,
		"password": "test",
		"email":    "test@test.pl",
	})))
	uc := UsersController{}
	uc.Register(app)
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, createUserReq)
	if rr.Code != 200 {
		t.Error("Failed to create user ", rr.Body.String())
	}*/
}

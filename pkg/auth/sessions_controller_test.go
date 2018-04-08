package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gopok/gopok-backend/pkg/core"
	"gopkg.in/mgo.v2/bson"
)

func TestCheckLogout(t *testing.T) {

	app := &core.Application{}
	app.Config = &core.CoreConfiguration{
		HTTPPort: 3002,
		MongoURL: "mongodb://localhost/go_test_TestCheckLogout",
	}
	app.Prepare()
	defer app.Db.DropDatabase() // cleanup
	uc := UsersController{}
	uc.Register(app)
	sc := SessionsController{}
	sc.Register(app)

	username := createTestUser(app)
	sessionToken := loginWithTestUser(app, username)
	rr := httptest.NewRecorder()

	logoutRequest, err := http.NewRequest("POST", "/api/auth/sessions/logout", nil)
	if err != nil {
		panic(err)
	}
	logoutRequest.Header.Set("Authorization", "Bearer "+sessionToken)

	app.Router.ServeHTTP(rr, logoutRequest)

	if rr.Code != 200 {
		panic(rr.Body.String())
	}

	sess := &Session{}
	app.Db.C("sessions").Find(bson.M{
		"token": sessionToken,
	}).One(sess)
	if sess.Active {
		t.Error("Session is still active after logout")
	}
}

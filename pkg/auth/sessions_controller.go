package auth

import (
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"

	"github.com/gopok/gopok-backend/pkg/core"
)

/*
SessionsController handles actions with session creation (login, logout, etc.).
*/
type SessionsController struct {
	app            *core.Application
	sessionsRouter *mux.Router
}

type loginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/*
Register registers the controller
*/
func (sc *SessionsController) Register(app *core.Application) {
	sc.app = app
	sc.sessionsRouter = app.Router.PathPrefix("/api/auth/sessions").Subrouter()
	sc.sessionsRouter.HandleFunc("/login", core.WrapRest(sc.login)).Methods("POST")

}

func (sc *SessionsController) login(r *core.RestRequest) interface{} {
	ld := &loginData{}
	r.DecodeJSON(ld)
	user := &User{}
	findErr := sc.app.Db.C("users").Find(bson.M{
		"username": ld.Username,
	}).One(user)
	if findErr != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ld.Password)) != nil {
		r.SetCode(400)
		return core.ErrorResponse{Code: 400,
			Message: "Invalid username or password.",
		}
	}
	sess := &session{
		ID:        bson.NewObjectId(),
		UserID:    user.ID,
		IPAddress: r.OriginalRequest.RemoteAddr,
		Active:    true,
		CreatedOn: time.Now(),
		ExpiresOn: time.Now().Add(time.Hour * 24 * 7),
	}
	sess.AssignToken()
	insertErr := sc.app.Db.C("sessions").Insert(sess)
	if insertErr != nil {
		r.SetCode(500)
		return core.ErrorResponse{
			Code:    500,
			Message: insertErr.Error(),
		}
	}
	return sess
}

func init() {
	core.ControllersToRegister.PushBack(&SessionsController{})
}

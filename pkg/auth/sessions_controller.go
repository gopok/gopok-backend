package auth

import (
	"net/http"
	"time"

	"github.com/fatih/structs"

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
	sc.sessionsRouter.Handle("/current-user", CheckUserMiddleware(app)(http.HandlerFunc(core.WrapRest(sc.getCurrentUser)))).Methods("GET")
	sc.sessionsRouter.Handle("/logout", CheckUserMiddleware(app)(http.HandlerFunc(core.WrapRest(sc.logout)))).Methods("POST")

}

func (sc *SessionsController) login(r *core.RestRequest) interface{} {
	ld := &loginData{}
	jsonErr := r.DecodeJSON(ld)
	if jsonErr != nil {
		return core.NewErrorResponse("invalid JSON request: "+jsonErr.Error(), 400)
	}
	user := &User{}
	findErr := sc.app.Db.C("users").Find(bson.M{
		"username": ld.Username,
	}).One(user)
	if findErr != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ld.Password)) != nil {
		return core.NewErrorResponse("Invalid username or password.", 401)
	}
	sess := &Session{
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
		return core.NewErrorResponse(insertErr.Error(), 500)
	}
	sessMap := structs.Map(&sess)
	sessMap["user"] = user
	return sessMap
}

func (sc *SessionsController) logout(r *core.RestRequest) interface{} {
	sess := r.OriginalRequest.Context().Value(SessionContextKey).(*Session)
	sess.Active = false
	sc.app.Db.C("sessions").Update(bson.M{
		"_id": sess.ID,
	}, bson.M{
		"$set": bson.M{
			"active": false,
		},
	})

	return bson.M{}
}

func (sc *SessionsController) getCurrentUser(r *core.RestRequest) interface{} {
	user := r.OriginalRequest.Context().Value(UserContextKey).(*User)
	return user
}

func init() {
	core.ControllersToRegister.PushBack(&SessionsController{})
}

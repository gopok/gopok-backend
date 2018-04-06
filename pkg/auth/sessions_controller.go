package auth

import (
	"github.com/gorilla/mux"

	"github.com/gopok/gopok-backend/pkg/core"
)

/*
SessionsController handles actions with session creation (login, logout, etc.).
*/
type SessionsController struct {
	app            *core.Application
	sessionsRouter *mux.Router
}

/*
Register registers the controller
*/
func (uc *SessionsController) Register(app *core.Application) {
	uc.app = app
	uc.sessionsRouter = app.Router.PathPrefix("/api/auth/sessions").Subrouter()
	uc.sessionsRouter.HandleFunc("/login", core.WrapRest(uc.login)).Methods("POST")

}

func (uc *SessionsController) login(r *core.RestRequest) interface{} {
	return nil
}

func init() {
	core.ControllersToRegister.PushBack(&SessionsController{})
}

package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gopok/gopok-backend/pkg/core"
)

/*
UsersController handles actions with user administration.
*/
type UsersController struct {
	app         *core.Application
	usersRouter *mux.Router
}

/*
Register registers the controller
*/
func (uc *UsersController) Register(app *core.Application) {
	uc.app = app
	uc.usersRouter = app.Router.PathPrefix("/api/auth/users").Subrouter()
	uc.usersRouter.HandleFunc("", uc.postUser).Methods("POST")
	
}

func (uc *UsersController) postUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TEST"))
}

func init() {
	core.ControllersToRegister.PushBack(&UsersController{})
}

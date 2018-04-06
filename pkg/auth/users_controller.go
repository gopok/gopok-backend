package auth

import (
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
	uc.usersRouter.HandleFunc("", core.WrapRest(uc.postUser)).Methods("POST")

}

func (uc *UsersController) postUser(r *core.RestRequest) interface{} {
	var u user
	r.DecodeJSON(&u)
	var allData map[string]string
	r.DecodeJSON(&allData)
	u.hashPassword(allData["password"])
	err := uc.app.Db.C("users").Insert(u)
	if err != nil {
		r.SetCode(500)
		return err.Error()
	}
	return u
}

func init() {
	core.ControllersToRegister.PushBack(&UsersController{})
}

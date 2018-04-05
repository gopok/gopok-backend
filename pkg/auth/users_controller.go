package auth

import "github.com/gopok/gopok-backend/pkg/core"

/*
UsersController handles actions with user administration.
*/
type UsersController struct{}

/*
Register registers the controller
*/
func (*UsersController) Register(app *core.Application) {
}

func init() {
	core.ControllersToRegister.PushBack(&UsersController{})
}

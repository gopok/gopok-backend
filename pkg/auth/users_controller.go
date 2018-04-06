package auth

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

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
	uc.usersRouter.HandleFunc("/index/{id}", core.WrapRest(uc.getUserByID)).Methods("GET")

}

func (uc *UsersController) postUser(r *core.RestRequest) interface{} {

	var allData map[string]string
	r.DecodeJSON(&allData)
	u := &User{
		Username: allData["username"],
		Email:    allData["email"],
	}
	u.HashPassword(allData["password"])

	validationError := u.Validate()
	if validationError != nil {
		return validationError
	}
	u.ID = bson.NewObjectId()
	err := uc.app.Db.C("users").Insert(&u)
	if err != nil {
		return core.NewErrorResponse(err.Error(), 500)
	}
	return u
}

func (uc *UsersController) getUserByID(r *core.RestRequest) interface{} {
	userID := mux.Vars(r.OriginalRequest)["id"]
	var u User
	if bson.IsObjectIdHex(userID) {
		err := uc.app.Db.C("users").FindId(bson.ObjectIdHex(userID)).One(&u)
		if err != nil {
			return core.NewErrorResponse(err.Error(), 500)
		}
	} else {
		return core.NewErrorResponse("User not found", 400)
	}
	return u
}

func init() {
	core.ControllersToRegister.PushBack(&UsersController{})
}

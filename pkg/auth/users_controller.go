package auth

import (
	"github.com/gopok/gopok-backend/pkg/core"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
UsersController handles actions with user administration.
*/
type UsersController struct {
	app         *core.Application
	usersRouter *mux.Router
}

/*
Register registers the controller (do not confuse with user sign up).
*/
func (uc *UsersController) Register(app *core.Application) {
	uc.app = app
	uc.usersRouter = app.Router.PathPrefix("/api/auth/users").Subrouter()
	uc.usersRouter.HandleFunc("", core.WrapRest(uc.postUser)).Methods("POST")
	uc.usersRouter.HandleFunc("/{id}", core.WrapRest(uc.getUserByID)).Methods("GET")

	indexErr := app.Db.C("users").EnsureIndex(mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	})

	if indexErr != nil {
		log.Panicf("Failed to create unique username index in users controller: %v", indexErr)
	}

}

func (uc *UsersController) postUser(r *core.RestRequest) interface{} {

	var allData map[string]string
	jsonErr := r.DecodeJSON(&allData)
	if jsonErr != nil {
		return core.NewErrorResponse("invalid JSON request: "+jsonErr.Error(), 400)
	}
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
		return core.NewErrorResponse("not found", 404)
	}
	return u
}

func init() {
	core.ControllersToRegister.PushBack(&UsersController{})
}

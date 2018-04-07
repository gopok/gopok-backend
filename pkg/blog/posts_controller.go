package blog

import (
"github.com/gopok/gopok-backend/pkg/core"
"github.com/gopok/gopok-backend/pkg/auth"
"github.com/gorilla/mux"
"gopkg.in/mgo.v2/bson"
"net/http"
)

/*
PostsController handles actions with posts.
*/
type PostsController struct {
	app             *core.Application
	postsController *mux.Router
}

/*
Register registers the controller
*/
func (pc *PostsController) Register(app *core.Application) {
	pc.app = app
	pc.postsController = app.Router.PathPrefix("/api/blog/posts").Subrouter()
	pc.postsController.Handle("", auth.CheckUserMiddleware(app)(http.HandlerFunc(core.WrapRest(pc.addPost)))).Methods("POST")
}

func (uc *PostsController) addPost(r *core.RestRequest) interface{} {
	user := r.OriginalRequest.Context().Value(auth.UserContextKey).(*auth.User)
	var allData map[string]string
	r.DecodeJSON(&allData)
	p := &Post{
		Content:    allData["content"],
		AuthorID:   user.ID,
	}
	validationError := p.Validate()
	if validationError != nil {
		return validationError
	}
	p.ID = bson.NewObjectId()
	err := uc.app.Db.C("posts").Insert(&p)
	if err != nil {
		return core.NewErrorResponse(err.Error(), 500)
	}
	return p
}

func init() {
	core.ControllersToRegister.PushBack(&PostsController{})
}

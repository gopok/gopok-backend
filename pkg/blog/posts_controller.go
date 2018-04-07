package blog

import (
	"net/http"
	"time"

	"github.com/gopok/gopok-backend/pkg/auth"
	"github.com/gopok/gopok-backend/pkg/core"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
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

func (pc *PostsController) addPost(r *core.RestRequest) interface{} {
	user := r.OriginalRequest.Context().Value(auth.UserContextKey).(*auth.User)
	var allData map[string]string
	jsonErr := r.DecodeJSON(&allData)
	if jsonErr != nil {
		return core.NewErrorResponse("invalid JSON request: "+jsonErr.Error(), 400)
	}
	p := &Post{
		Content:  allData["content"],
		AuthorID: user.ID,
	}
	validationError := p.Validate()
	if validationError != nil {
		return validationError
	}
	p.ID = bson.NewObjectId()
	p.CreatedOn = time.Now()
	err := pc.app.Db.C("posts").Insert(&p)
	if err != nil {
		return core.NewErrorResponse(err.Error(), 500)
	}
	return p
}

func init() {
	core.ControllersToRegister.PushBack(&PostsController{})
}

package blog

import (
	"time"

	"github.com/fatih/structs"
	"github.com/gopok/gopok-backend/pkg/auth"
	"github.com/gopok/gopok-backend/pkg/core"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func (pc *PostsController) addComment(r *core.RestRequest) interface{} {
	postID := mux.Vars(r.OriginalRequest)["id"]
	user := r.OriginalRequest.Context().Value(auth.UserContextKey).(*auth.User)
	var allData map[string]string
	jsonErr := r.DecodeJSON(&allData)
	if jsonErr != nil {
		return core.NewErrorResponse("invalid JSON request: "+jsonErr.Error(), 400)
	}

	c := &Comment{
		AuthorID: user.ID,
		Content:  allData["content"],
	}
	validationError := c.Validate()
	if validationError != nil {
		return validationError
	}
	c.ID = bson.NewObjectId()
	c.CreatedOn = time.Now()

	if bson.IsObjectIdHex(postID) {

		err := pc.app.Db.C("posts").Update(bson.M{
			"_id": bson.ObjectIdHex(postID),
		}, bson.M{
			"$push": bson.M{
				"comments": c,
			},
		})
		if err != nil {
			return core.NewErrorResponse(err.Error(), 500)
		}

	} else {
		return core.NewErrorResponse("not found", 404)
	}

	return c
}

func (pc *PostsController) attachAuthorToComment(c *Comment) map[string]interface{} {
	cm := structs.Map(c)
	author := &auth.User{}
	pc.app.Db.C("users").FindId(c.AuthorID).One(author)
	cm["author"] = author
	return cm
}

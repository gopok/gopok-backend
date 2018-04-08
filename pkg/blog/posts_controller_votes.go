package blog

import (
	"github.com/gopok/gopok-backend/pkg/auth"
	"github.com/gopok/gopok-backend/pkg/core"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (pc *PostsController) upvotePost(r *core.RestRequest) interface{} {
	postID := mux.Vars(r.OriginalRequest)["id"]
	user := r.OriginalRequest.Context().Value(auth.UserContextKey).(*auth.User)
	post := &Post{}
	if bson.IsObjectIdHex(postID) {

		_, findErr := pc.app.Db.C("posts").Find(bson.M{
			"_id": bson.ObjectIdHex(postID),

			"upvoters":   bson.M{"$ne": user.ID},
			"downvoters": bson.M{"$ne": user.ID},
		}).Apply(mgo.Change{
			Update: bson.M{
				"$addToSet": bson.M{
					"upvoters": user.ID,
				},
			},
			ReturnNew: true,
		}, post)
		if findErr != nil {
			return core.NewErrorResponse(findErr.Error(), 500)
		}
	} else {
		return core.NewErrorResponse("not found", 404)
	}

	return map[string]interface{}{}
}

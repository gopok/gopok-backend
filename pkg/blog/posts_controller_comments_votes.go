package blog

import (
	"github.com/gopok/gopok-backend/pkg/auth"
	"github.com/gopok/gopok-backend/pkg/core"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (pc *PostsController) upvoteComment(r *core.RestRequest) interface{} {
	postID := mux.Vars(r.OriginalRequest)["id"]
	commentID := mux.Vars(r.OriginalRequest)["commentID"]
	user := r.OriginalRequest.Context().Value(auth.UserContextKey).(*auth.User)
	post := &Post{}
	if bson.IsObjectIdHex(postID) && bson.IsObjectIdHex(commentID) {
		_, findErr := pc.app.Db.C("posts").Find(bson.M{
			"_id": bson.ObjectIdHex(postID),
			"comments": bson.M{
				"$elemMatch": bson.M{
					"_id":        bson.ObjectIdHex(commentID),
					"upvoters":   bson.M{"$ne": user.ID},
					"downvoters": bson.M{"$ne": user.ID},
				},
			},
		}).Apply(mgo.Change{
			Update: bson.M{
				"$addToSet": bson.M{
					"comments.$.upvoters": user.ID,
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

func (pc *PostsController) downvoteComment(r *core.RestRequest) interface{} {
	postID := mux.Vars(r.OriginalRequest)["id"]
	commentID := mux.Vars(r.OriginalRequest)["commentID"]
	user := r.OriginalRequest.Context().Value(auth.UserContextKey).(*auth.User)
	post := &Post{}
	if bson.IsObjectIdHex(postID) && bson.IsObjectIdHex(commentID) {
		_, findErr := pc.app.Db.C("posts").Find(bson.M{
			"_id": bson.ObjectIdHex(postID),
			"comments": bson.M{
				"$elemMatch": bson.M{
					"_id":        bson.ObjectIdHex(commentID),
					"upvoters":   bson.M{"$ne": user.ID},
					"downvoters": bson.M{"$ne": user.ID},
				},
			},
		}).Apply(mgo.Change{
			Update: bson.M{
				"$addToSet": bson.M{
					"comments.$.downvoters": user.ID,
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

package blog

import (
	"time"

	"github.com/gopok/gopok-backend/pkg/core"
	"gopkg.in/mgo.v2/bson"
)

/*
Post is a model of... a post.
*/
type Post struct {
	ID         bson.ObjectId   `json:"id" structs:"id" bson:"_id,omitempty"`
	AuthorID   bson.ObjectId   `json:"authorID" structs:"authorID" bson:"authorID,omitempty"`
	Content    string          `json:"content" structs:"content" bson:"content"`
	Comments   []Comment       `json:"comments" structs:"comments" bson:"comments"`
	CreatedOn  time.Time       `json:"createdOn" structs:"createdOn" bson:"createdOn"`
	Upvoters   []bson.ObjectId `json:"upvoters" structs:"upvoters" bson:"upvoters"`
	Downvoters []bson.ObjectId `json:"downvoters" structs:"downvoters" bson:"downvoters"`
}

/*
Validate checks if all fields of the user conform to the rules.
Currently checks if content is between 3 and 1000 characters.
*/
func (u *Post) Validate() core.ValidationError {
	if len(u.Content) < 3 {
		return core.NewValidationError("Post content is too short", "content", "blog.post")
	}

	if len(u.Content) > 3000 {
		return core.NewValidationError("Post content is too long", "content", "blog.post")
	}
	return nil
}

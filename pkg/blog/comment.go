package blog

import (
	"time"

	"github.com/gopok/gopok-backend/pkg/core"
	"gopkg.in/mgo.v2/bson"
)

/*
Comment is a model of... a comment.
*/
type Comment struct {
	ID        bson.ObjectId `json:"id" structs:"id" bson:"_id,omitempty"`
	AuthorID  bson.ObjectId `json:"authorID" structs:"authorID" bson:"authorID,omitempty"`
	Content   string        `json:"content" structs:"content" bson:"content"`
	CreatedOn time.Time     `json:"createdOn" structs:"createdOn" bson:"createdOn"`
}

/*
Validate checks if all fields of the comment conform to the rules.
Currently checks if content is between 3 and 1000 characters.
*/
func (u *Comment) Validate() core.ValidationError {
	if len(u.Content) < 3 {
		return core.NewValidationError("Post content is too short", "content", "blog.comment")
	}

	if len(u.Content) > 1000 {
		return core.NewValidationError("Post content is too long", "content", "blog.comment")
	}
	return nil
}

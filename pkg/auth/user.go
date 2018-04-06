package auth

import (
	"regexp"

	"github.com/gopok/gopok-backend/pkg/core"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

/*
User is a model of... a user.
*/
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"-" bson:"password"`
	Email    string        `json:"email" bson:"email"`
}

func (u *User) HashPassword(password string) error {

	rawHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return errors.Wrap(err, "failed to hash user password:")
	}
	u.Password = string(rawHash)
	return nil
}

var usernameRegexp = regexp.MustCompile("^([A-Za-z0-9_]){3,20}$")
var emailRegexp = regexp.MustCompile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)")

func (u *User) Validate() core.ValidationError {
	if !usernameRegexp.Match([]byte(u.Username)) {
		return core.NewValidationError("Invalid username", "username", "auth.user")
	}
	if !emailRegexp.Match([]byte(u.Email)) {
		return core.NewValidationError("Invalid email", "email", "auth.user")
	}
	return nil
}
